package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func MarshalJson(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encodeJson(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encodeJson(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encodeJson(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encodeJson(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(buf, "%q:", v.Type().Field(i).Name)
			if err := encodeJson(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encodeJson(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encodeJson(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "t")
		} else {
			fmt.Fprintf(buf, "nil")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%f %f)", real(v.Complex()), imag(v.Complex()))

	case reflect.Interface:
		if err := encodeJson(buf, v.Elem()); err != nil {
			return err
		}

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
