package copier

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type Option struct {
	ignoreEmpty bool
	overwrite   bool // slice not support overwrite, will panic
}

type copyOption struct {
	needCheckLevel *bool
}

func NewOption() *Option {
	return &Option{
		ignoreEmpty: false,
		overwrite:   true,
	}
}

func (o *Option) SetOverwrite(overwrite bool) *Option {
	o.overwrite = overwrite
	return o
}

func (o *Option) SetIgnoreEmpty(ignoreEmpty bool) *Option {
	o.ignoreEmpty = ignoreEmpty
	return o
}

func Instance(option *Option) Mapper {
	if option == nil {
		option = &Option{
			ignoreEmpty: false,
			overwrite:   true,
		}
	} else {
		// empty
	}
	return &mapper{
		converterRepository: newConverterRepository(option),
	}
}

type mapper struct {
	converterRepository *converterRepository
}

type copyCommand struct {
	*mapper
	fromValue interface{}
}

func (c *copyCommand) CopyTo(toValue interface{}) (err error) {
	return c.mapper.copy(toValue, c.fromValue)
}

func (m *mapper) From(fromValue interface{}) CopyCommand {
	return &copyCommand{mapper: m, fromValue: fromValue}
}

func (m *mapper) copy(toValue, fromValue interface{}) error {
	return m.copyValue(reflect.ValueOf(toValue), reflect.ValueOf(fromValue))
}

func (m *mapper) copyValue(to, from reflect.Value) error {
	m.converterRepository.level++
	if !from.IsValid() {
		return nil
	}
	if m.shouldCopy(to, from) {
		if from.Kind() == reflect.Ptr && to.Kind() == reflect.Ptr && from.IsNil() {
			to.Set(reflect.Zero(to.Type()))
			return nil
		}
		v, err := m.convert(indirect(from), indirect(to), indirectType(to.Type()))
		if err != nil {
			indirectAsNonNil(to).Set(reflect.New(indirectType(to.Type())).Elem())
			fmt.Printf("can't convert data %+v -> %+v\n", indirect(from), indirectType(to.Type()))
			return nil
		}
		indirectAsNonNil(to).Set(v)
	}
	m.converterRepository.level--
	return nil
}

func (m *mapper) convertSlice(from reflect.Value, toType reflect.Type) (reflect.Value, error) {
	amount := from.Len()
	destType := toType.Elem()
	to := reflect.MakeSlice(toType, 0, amount)

	for i := 0; i < amount; i++ {
		source := from.Index(i)

		dest, err := m.convert(source, reflect.ValueOf(nil), indirectType(destType))
		if err != nil {
			return to, err
		}

		if destType.Kind() == reflect.Ptr {
			to = reflect.Append(to, forceAddr(dest))
		} else {
			to = reflect.Append(to, dest)
		}
	}

	return to, nil
}

func (m *mapper) namesFromDiffFields(field reflect.StructField) []string {
	if name, ok := m.converterRepository.diffFieldsMapper[field.Name]; ok {
		return name
	} else {
		return []string{field.Name}
	}
}

func (m *mapper) handleMultiLevelFields(from, to reflect.Value) {
	needCheckLevel := false
	option := copyOption{needCheckLevel: &needCheckLevel}
	for originKey, targetKeys := range m.converterRepository.diffFieldsMapper {
		for _, targetKey := range targetKeys {
			if strings.Contains(originKey, ".") || strings.Contains(targetKey, ".") {
				target, origin := getValueByFiledName(to, targetKey), getValueByFiledName(from, originKey)
				if m.shouldCopy(target, origin, option) {
					var value reflect.Value
					var err error
					if transformerMethod, ok := m.converterRepository.transformer[targetKey]; ok {
						f := reflect.ValueOf(transformerMethod)
						result := f.Call(
							[]reflect.Value{origin},
						)
						value, err = m.convert(result[0], target, target.Type(), option)
					} else {
						value, err = m.convert(origin, target, target.Type(), option)
					}
					if err != nil {
						panic(err)
					}
					target.Set(value)
				}
			}
		}
	}
}

func (m *mapper) convertStruct(from, to reflect.Value, toType reflect.Type) (reflect.Value, error) {
	if m.needOverwrite(from) || !to.IsValid() {
		to = reflect.New(toType).Elem()
	}
	toFields := asNamesToFieldMap(deepFields(to.Type()))

	copied := make(map[string]struct{})

	for _, fromField := range deepFields(from.Type()) {
		if fromValue := from.FieldByName(fromField.Name); fromValue.IsValid() {
			names := m.namesFromDiffFields(fromField)
			for _, name := range names {
				if toField, found := toFields[name]; found {
					if _, ok := copied[toField.Name]; !ok {
						if toValue := to.FieldByName(toField.Name); toValue.IsValid() && toValue.CanSet() {
							if transformerMethod, ok := m.converterRepository.transformer[toField.Name]; ok {
								f := reflect.ValueOf(transformerMethod)
								result := f.Call(
									[]reflect.Value{fromValue},
								)
								if err := m.copyValue(toValue, result[0]); err != nil {
									return to, err
								}
							} else if err := m.copyValue(toValue, fromValue); err != nil {
								return to, err
							}
						}
						copied[toField.Name] = struct{}{}
					}
				}
			}
		}
	}

	// 处理多级
	if m.converterRepository.level == 0 {
		m.handleMultiLevelFields(from, to)
	}

	return to, nil
}

func (m *mapper) shouldCopy(toValue, fromValue reflect.Value, options ...copyOption) bool {
	if m.checkLevel(options...) && (m.converterRepository.ignoreEmpty && fromValue.IsZero() || !m.converterRepository.overwrite && !toValue.IsZero()) {
		return false
	}
	return true
}

func (m *mapper) needOverwrite(fromValue reflect.Value) bool {
	if m.converterRepository.level > 0 && !fromValue.IsZero() && m.converterRepository.overwrite {
		return true
	}
	return false
}

func (m *mapper) checkLevel(options ...copyOption) bool {
	var option *copyOption
	if len(options) > 0 {
		option = &options[0]
	}
	if option != nil && !*option.needCheckLevel {
		return true
	}

	return m.converterRepository.level > 0
}

func (m *mapper) convert(from, to reflect.Value, toType reflect.Type, options ...copyOption) (reflect.Value, error) {
	if !from.IsValid() {
		return reflect.Zero(toType), nil
	}
	if converter := m.converterRepository.Get(Target{To: toType, From: from.Type()}); converter != nil {
		return converter(from, toType)

	} else if from.Type().ConvertibleTo(toType) && m.checkLevel(options...) {
		return from.Convert(toType), nil

	} else if m.canScan(toType) {
		return m.scan(from, toType)

	} else if from.Kind() == reflect.Ptr || to.Kind() == reflect.Ptr {
		if from.Kind() == reflect.Ptr && to.Kind() == reflect.Ptr {
			return m.convert(from.Elem(), to.Elem(), toType)
		} else if from.Kind() == reflect.Ptr {
			return m.convert(from.Elem(), to, toType)
		} else {
			return m.convert(from, to.Elem(), toType)
		}

	} else if from.Kind() == reflect.Struct && toType.Kind() == reflect.Struct {
		return m.convertStruct(from, to, toType)

	} else if from.Kind() == reflect.Slice && toType.Kind() == reflect.Slice {
		return m.convertSlice(from, toType)

	} else {
		return reflect.Zero(toType), errors.Errorf("can't convert data %+v -> %+v", from, toType)

	}
}

func (m *mapper) canScan(t reflect.Type) bool {
	return reflect.PtrTo(t).Implements(scannerType)
}

func (m *mapper) scan(from reflect.Value, toType reflect.Type) (reflect.Value, error) {
	v := reflect.New(toType)
	scanner := v.Interface().(sql.Scanner)
	err := scanner.Scan(from.Interface())
	if err != nil {
		return reflect.Zero(toType), err
	}
	return v.Elem(), nil
}

var scannerType = reflect.TypeOf((*sql.Scanner)(nil)).Elem()

func forceAddr(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return v
	} else if v.CanAddr() {
		return v.Addr()
	}

	ptr := reflect.New(v.Type())
	ptr.Elem().Set(v)
	return ptr
}

func asNamesToFieldMap(fields []reflect.StructField) map[string]reflect.StructField {
	m := make(map[string]reflect.StructField)
	for _, field := range fields {
		m[field.Name] = field
	}
	return m
}

func (m *mapper) RegisterConverter(matcher TypeMatcher, converter Converter) Mapper {
	m.converterRepository.Put(matcher, converter)
	return m
}

func (m *mapper) RegisterResetDiffField(diffFields []DiffFieldPair) Mapper {
	for _, diffField := range diffFields {
		m.converterRepository.diffFieldsMapper[diffField.Origin] = diffField.Targets
	}
	return m
}

func (m *mapper) RegisterTransformer(transformer Transformer) Mapper {
	for toField, transformerMethod := range transformer {
		if reflect.TypeOf(transformerMethod).Kind() != reflect.Func {
			panic("transfer need to be a function")
		}
		m.converterRepository.transformer[toField] = transformerMethod
	}
	return m
}

func (m *mapper) RegisterConverterFunc(matcherFunc TypeMatcherFunc, converter Converter) Mapper {
	return m.RegisterConverter(matcherFunc, converter)
}

func (m *mapper) Install(module Module) Mapper {
	module(m)
	return m
}

type converterPair struct {
	Matcher   TypeMatcher
	converter Converter
}

type converterRepository struct {
	converters       []converterPair
	diffFieldsMapper map[string][]string
	transformer      map[string]interface{}
	ignoreEmpty      bool
	overwrite        bool
	level            int
	lastLevel        int
}

func newConverterRepository(option *Option) *converterRepository {
	return &converterRepository{
		converters:       nil,
		diffFieldsMapper: make(map[string][]string),
		transformer:      make(map[string]interface{}),
		ignoreEmpty:      option.ignoreEmpty,
		overwrite:        option.overwrite,
		level:            -1,
		lastLevel:        -1,
	}
}

func (r *converterRepository) Put(matcher TypeMatcher, converter Converter) {
	r.converters = append(r.converters, converterPair{matcher, converter})
}

func (r *converterRepository) Get(target Target) Converter {
	for _, pair := range r.converters {
		matches := pair.Matcher.Matches(target)
		if matches {
			return pair.converter
		}
	}
	return nil
}

func getValueByFiledName(value reflect.Value, name string) reflect.Value {
	if strings.ContainsAny(name, ".") {
		firstKey := strings.Split(name, ".")[0]
		name = strings.Join(append(strings.Split(name, ".")[1:]), ".")
		return getValueByFiledName(getValueByFiledName(value, firstKey), name)
	} else {
		switch value.Kind() {
		case reflect.Ptr:
			if value.IsNil() {
				return reflect.Zero(value.Type())
			}
			return value.Elem().FieldByName(name)
		case reflect.Struct:
			return value.FieldByName(name)
		default:
			panic(fmt.Sprintf("not support get value from type[%s] by key[%s]", value.Kind(), name))
		}
	}
}
