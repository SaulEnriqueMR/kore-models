package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func DecodeInterface[T any](in *T, data []byte, concreteType ...reflect.Type) (err error) {
	if reflect.TypeOf(in).Elem().Kind() != reflect.Interface {
		return fmt.Errorf("in must be a pointer to an interface")
	}

	if len(concreteType) == 0 {
		err = json.Unmarshal(data, &in)
		return
	}

	errs := make([]error, 0)
	for _, concreteType := range concreteType {
		valuePtr := reflect.New(concreteType).Interface()
		if err := json.Unmarshal(data, &valuePtr); err != nil {
			errs = append(errs, err)
			continue
		}

		valueVal := reflect.ValueOf(valuePtr).Elem().Interface()
		if conv, ok := valueVal.(T); ok {
			*in = conv
			return
		}
		if conv, ok := valuePtr.(T); ok {
			*in = conv
			return
		}
	}

	return fmt.Errorf("type %T is not assignable to Value (expects %T)\n%e", data, in, errors.Join(errs...))
}
