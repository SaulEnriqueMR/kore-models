package helpers

// SafeUnwrap safely dereferences a pointer to a value of any type T.
//
//   - If the input pointer is not nil, it returns the dereferenced value.
//
//   - If the input pointer is nil and a default value is provided, it returns the first default value.
//
//   - If the input pointer is nil and no default value is provided, it returns the zero value of type T.
//
// example:
//
//	mayBeInt := new(int) // pointer to int with nil value
//	result = SafeUnwrap(mayBeInt) // result will be 0 (zero value of int)
//	result = SafeUnwrap(mayBeInt, 100) // result will be 100 (default value provided)
//
//	*mayBeInt = 42 // pointer to int with value 42
//	result := SafeUnwrap(mayBeInt) // result will be 42
//	result := SafeUnwrap(mayBeInt, 100) // result will be 42 (default value ignored)
func SafeUnwrap[T any](input *T, defaultValue ...T) T {
	var result T

	if input != nil {
		result = *input
	} else if len(defaultValue) > 0 {
		result = defaultValue[0]
	}

	return result
}
