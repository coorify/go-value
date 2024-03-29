package value

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Get(v interface{}, name string) (interface{}, error) {
	ns := strings.Split(name, ".")
	nv := reflect.ValueOf(v)

	for idx, n := range ns {
		knd := nv.Kind()

		if knd == reflect.Ptr {
			nv = reflect.Indirect(nv)
			knd = nv.Kind()
		}

		if !nv.IsValid() {
			return nil, fmt.Errorf("invalid field: %s", n)
		}

		switch knd {
		case reflect.Struct:
			nv = nv.FieldByName(n)
			if nv.Kind() == reflect.Invalid {
				return nil, fmt.Errorf("invalid field: %s", n)
			}
		case reflect.Array, reflect.Slice:
			i, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			nv = nv.Index(i)
		case reflect.Map:
			fd := false

			for _, k := range nv.MapKeys() {
				if fmt.Sprint(k.Interface()) == n {
					nv = nv.MapIndex(k)
					fd = true
					break
				}
			}

			if !fd {
				return nil, fmt.Errorf("invalid field: %s", n)
			}
		default:
			return nil, fmt.Errorf("invalid field: %s, should be struct,map,array", ns[idx-1])
		}
	}

	if !nv.IsValid() {
		return nil, fmt.Errorf("invalid field: %s", name)
	} else if nv.CanInterface() {
		return nv.Interface(), nil
	}

	return nil, fmt.Errorf("unexported field: %s", name)
}

func MustGet(v interface{}, name string) interface{} {
	val, err := Get(v, name)
	if err != nil {
		panic(err)
	}
	return val
}

func GetWithDefault(v interface{}, name string, defaultValue interface{}) interface{} {
	val, err := Get(v, name)
	if err != nil {
		return defaultValue
	}
	return val
}
