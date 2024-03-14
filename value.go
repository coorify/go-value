package value

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Get(v interface{}, name string) (interface{}, error) {
	ns := strings.Split(name, ".")
	nv := reflect.ValueOf(v)

	switch nv.Kind() {
	case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		for _, n := range ns {
			switch nv.Kind() {
			case reflect.Struct:
				nv = nv.FieldByName(n)
				if nv.Kind() == reflect.Invalid {
					return nil, fmt.Errorf("invalid field: %s", name)
				}
			case reflect.Array, reflect.Slice:
				i, err := strconv.Atoi(n)
				if err != nil {
					return nil, err
				}
				nv = nv.Index(i)
			case reflect.Map:
				mk := reflect.ValueOf(n)
				ks := nv.MapKeys()
				fd := false

				for _, k := range ks {
					if k.CanInt() || k.CanUint() {
						i, err := strconv.Atoi(n)
						if err != nil {
							return nil, err
						}
						mk = reflect.ValueOf(i)
					}

					if k.Convert(mk.Type()).Equal(mk) {
						nv = nv.MapIndex(k)
						fd = true
						break
					}
				}

				if !fd {
					return nil, fmt.Errorf("invalid field: %s", n)
				}
			}
		}
		if nv.CanInterface() {
			return nv.Interface(), nil
		}

		return nil, fmt.Errorf("unexported field: %s", name)
	default:
		return nil, errors.New("invalid src, should be struct,map,array")
	}
}

func MustGet(v interface{}, name string) interface{} {
	val, err := Get(v, name)
	if err != nil {
		panic(err)
	}
	return val
}