package whmcs

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// SerializeParams mengonversi struct atau map Go menjadi url.Values yang sesuai dengan
// konvensi array bertingkat PHP pada API WHMCS (application/x-www-form-urlencoded).
func SerializeParams(input interface{}) (url.Values, error) {
	values := url.Values{}
	if input == nil {
		return values, nil
	}

	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return values, nil
		}
		val = val.Elem()
	}

	if err := flattenValue(val, "", values); err != nil {
		return nil, err
	}

	return values, nil
}

func flattenValue(val reflect.Value, prefix string, out url.Values) error {
	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			kStr := fmt.Sprintf("%v", key.Interface())
			childPrefix := kStr
			if prefix != "" {
				childPrefix = fmt.Sprintf("%s[%s]", prefix, kStr)
			}
			mapVal := val.MapIndex(key)
			if err := flattenValue(mapVal, childPrefix, out); err != nil {
				return err
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			childPrefix := fmt.Sprintf("%s[%d]", prefix, i)
			if err := flattenValue(elem, childPrefix, out); err != nil {
				return err
			}
		}
	case reflect.Struct:
		t := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := t.Field(i)
			if !field.IsExported() {
				continue
			}

			tag := field.Tag.Get("url")
			if tag == "-" {
				continue
			}

			parts := strings.Split(tag, ",")
			tagKey := parts[0]
			omitempty := false
			if len(parts) > 1 && parts[1] == "omitempty" {
				omitempty = true
			}

			if tagKey == "" {
				tagKey = strings.ToLower(field.Name)
			}

			fieldVal := val.Field(i)
			if omitempty && isEmptyValue(fieldVal) {
				continue
			}

			childPrefix := tagKey
			if prefix != "" {
				childPrefix = fmt.Sprintf("%s[%s]", prefix, tagKey)
			}

			if err := flattenValue(fieldVal, childPrefix, out); err != nil {
				return err
			}
		}
	case reflect.Interface:
		if !val.IsNil() {
			return flattenValue(val.Elem(), prefix, out)
		}
	case reflect.Ptr:
		if !val.IsNil() {
			return flattenValue(val.Elem(), prefix, out)
		}
	default:
		if prefix != "" {
			out.Set(prefix, formatScalar(val))
		}
	}
	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func formatScalar(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.String:
		return v.String()
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}
