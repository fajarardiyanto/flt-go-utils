package flags

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/urfave/cli/v2"
)

func ParseStruct(data interface{}) {
	if reflect.ValueOf(data).Type().Kind() == reflect.Struct {
		fmt.Println("expected input not be pointer")
		return
	}

	v := reflect.ValueOf(data)
	t := reflect.Indirect(v).Type()

	for i := 0; i < t.NumField(); i++ {
		val := t.Field(i).Tag.Get("default")
		elem := v.Elem().Field(i)
		if elem.CanSet() {
			switch elem.Type().Kind() {
			case reflect.String:
				elem.SetString(val)
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				if sp, err := strconv.ParseInt(val, 10, 64); err == nil {
					elem.SetInt(sp)
				}
			case reflect.Float32, reflect.Float64:
				if sp, err := strconv.ParseFloat(val, 64); err == nil {
					elem.SetFloat(sp)
				}
			case reflect.Bool:
				if sp, err := strconv.ParseBool(val); err == nil {
					elem.SetBool(sp)
				}
			case reflect.Slice:
				if vas, ok := v.Interface().([]string); ok {
					sclie := cli.NewStringSlice(vas...)
					var values []string
					for _, sv := range sclie.Value() {
						if len(sv) != 0 {
							values = append(values, sv)
						}
					}

					if len(values) != 0 {
						elem.Set(reflect.ValueOf(values))
					}
				}

				if vas, ok := isIntSlice(v); ok {
					sclie := cli.NewInt64Slice(vas...)
					if _, ok := v.Interface().([]int); ok {
						var values []int
						for _, sv := range sclie.Value() {
							values = append(values, int(sv))
						}
						v.Set(reflect.ValueOf(values))
					}

					if _, ok := v.Interface().([]int16); ok {
						var values []int16
						for _, sv := range sclie.Value() {
							values = append(values, int16(sv))
						}
						v.Set(reflect.ValueOf(values))
					}

					if _, ok := v.Interface().([]int32); ok {
						var values []int32
						for _, sv := range sclie.Value() {
							values = append(values, int32(sv))
						}
						v.Set(reflect.ValueOf(values))
					}

					if _, ok := v.Interface().([]int64); ok {
						v.Set(reflect.ValueOf(sclie.Value()))
					}
				}

				if vas, ok := isFloatSlice(v); ok {
					sclie := cli.NewFloat64Slice(vas...)
					if _, ok := v.Interface().([]float32); ok {
						var values []float32
						for _, sv := range sclie.Value() {
							values = append(values, float32(sv))
						}
						v.Set(reflect.ValueOf(values))
					}

					if _, ok := v.Interface().([]float64); ok {
						v.Set(reflect.ValueOf(sclie.Value()))
					}
				}
			}
		}
	}
}

func isIntSlice(val reflect.Value) ([]int64, bool) {
	if v, ok := val.Interface().([]int); ok {
		var vss []int64
		for _, s := range v {
			vss = append(vss, int64(s))
		}
		return vss, true
	}
	if v, ok := val.Interface().([]int16); ok {
		var vss []int64
		for _, s := range v {
			vss = append(vss, int64(s))
		}
		return vss, true
	}
	if v, ok := val.Interface().([]int32); ok {
		var vss []int64
		for _, s := range v {
			vss = append(vss, int64(s))
		}
		return vss, true
	}
	if v, ok := val.Interface().([]int64); ok {
		return v, true
	}

	return nil, false
}

func isFloatSlice(val reflect.Value) ([]float64, bool) {
	if v, ok := val.Interface().([]float32); ok {
		var vss []float64
		for _, s := range v {
			vss = append(vss, float64(s))
		}
		return vss, true
	}
	if v, ok := val.Interface().([]float64); ok {
		return v, true
	}

	return nil, false
}
