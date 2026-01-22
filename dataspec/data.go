package dataspec

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"

	"github.com/shopspring/decimal"
)

// DataDescription 数据描述，代表某个变量的数据元数据
type DataDescription struct {
	// Type 数据类型，目前支持 string|number|boolean|integer|array|struct
	Type DataType `json:"type"`

	// Specs 对应Type的数据类型，供外部使用
	Specs DataSpec `json:"specs"`
}

// UnmarshalJSON 自定义 JSON 反序列化
func (d *DataDescription) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Type  DataType        `json:"type"`
		Specs json.RawMessage `json:"specs"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	d.Type = tmp.Type

	switch d.Type {
	case NumberType:
		d.Specs = &NumericDataSpec{
			Min:       -math.MaxFloat64,
			Max:       math.MaxFloat64,
			Precision: 1e-12,
		}
	case StringType:
		d.Specs = &StringDataSpec{}
	case IntegerType:
		d.Specs = &IntegerDataSpec{
			Min: math.MinInt64,
			Max: math.MaxInt64,
		}
	case BooleanType:
		d.Specs = &BooleanDataSpec{}
	case VoidType:
		d.Specs = &VoidDataSpec{}
	}

	return d.parseSpecs(tmp.Specs)
}

// MarshalJSON 自定义 JSON 序列化
func (d *DataDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type  DataType `json:"type"`
		Specs DataSpec `json:"specs"`
	}{
		Type:  d.Type,
		Specs: d.Specs,
	})
}

func (d *DataDescription) parseSpecs(specs json.RawMessage) error {
	if d.Type == ArrayType {
		return d.parseArray(specs)
	}
	if d.Type == StructType {
		return d.parseStruct(specs)
	}

	if err := json.Unmarshal(specs, d.Specs); err != nil {
		return err
	}

	if d.Specs == nil {
		switch d.Type {
		case NumberType:
			d.Specs = &NumericDataSpec{
				Min:       -math.MaxFloat64,
				Max:       math.MaxFloat64,
				Precision: 1e-12,
			}
		case StringType:
			d.Specs = &StringDataSpec{}
		case IntegerType:
			d.Specs = &IntegerDataSpec{
				Min: math.MinInt64,
				Max: math.MaxInt64,
			}
		case BooleanType:
			d.Specs = &BooleanDataSpec{}
		case VoidType:
			d.Specs = &VoidDataSpec{}
		}
	}
	return nil
}

func (d *DataDescription) parseStruct(specs json.RawMessage) error {
	ds := StructDataSpec{}
	if err := json.Unmarshal(specs, &ds); err != nil {
		return err
	}
	d.Specs = ds
	return nil
}

func (d *DataDescription) parseArray(specs json.RawMessage) error {
	arr := &ArrayDataSpec{}
	if err := json.Unmarshal(specs, arr); err != nil {
		return err
	}
	d.Specs = arr

	if arr.Length == 0 {
		return fmt.Errorf("ArrayDataSpecs: array max length could not be zero")
	}

	if arr.Data == nil {
		return fmt.Errorf("ArrayDataSpecs: data field could not be empty")
	}
	return nil
}

func (d *DataDescription) Validate(v interface{}) (bool, error) {
	return validateData(d, v)
}

func validateReflectData(ds *DataDescription, v reflect.Value) (bool, error) {
	kind := v.Kind()
	if kind == reflect.Interface || kind == reflect.Pointer {
		v = v.Elem()
	}

	switch specs := ds.Specs.(type) {
	case *StringDataSpec:
		if v.Kind() == reflect.String {
			return specs.ValidateString(v.String())
		}
	case *IntegerDataSpec:
		if v.CanInt() {
			return specs.ValidateInteger(v.Int())
		} else if v.CanUint() {
			return specs.ValidateInteger(int64(v.Uint()))
		} else if v.CanFloat() {
			// 这里为了解决json数据的整数情况，因为json是不存在整数的，所以当浮点没有小数点后的数，则为整数时，认为是整数
			val := v.Float()
			num := decimal.NewFromFloat(val)
			rval := math.Round(val)
			if !num.Equal(decimal.NewFromFloat(rval)) {
				return false, fmt.Errorf("DataSpecs: type [%s] or value [%s] is not supported", ds.Type, v.Kind().String())
			}
			// 这里因为转换，为了保证能转换为整数，那么加上0.5，保证超过整数值
			result := int64(rval + 0.5)
			return specs.ValidateInteger(result)
		}
	case *NumericDataSpec:
		if v.CanFloat() {
			return specs.ValidateNumber(v.Float())
		} else if v.CanInt() {
			return specs.ValidateNumber(float64(v.Int()))
		} else if v.CanUint() {
			return specs.ValidateNumber(float64(v.Uint()))
		}
	case *BooleanDataSpec:
		if v.Kind() == reflect.Bool {
			return true, nil
		}
	case *ArrayDataSpec:
		return specs.Validate(v.Interface())
	case StructDataSpec:
		return specs.Validate(v.Interface())
	case *VoidDataSpec:
		return true, nil
	}
	return false, fmt.Errorf("DataSpecs: type [%s] or value [%s] is not supported", ds.Type, v.Kind().String())
}

func validateData(ds *DataDescription, v interface{}) (bool, error) {
	i := reflect.ValueOf(v)
	return validateReflectData(ds, i)
}
