package converters

import (
	"fmt"
	"reflect"
	"strings"
)

// drillKey gets deep down key value
func drillKey(dto map[string]any, path string) (any, bool) {
	if dto == nil {
		return nil, false
	}

	keys := strings.Split(path, ".0.")
	keysLen := len(keys) - 1
	for i := 0; ; i++ {
		v, ok := dto[keys[i]]
		if !ok {
			return nil, false
		}

		isLast := i == keysLen
		if isLast {
			return v, true
		}

		next, ok := v.(map[string]any)
		if ok {
			dto = next
			continue
		}

		// Gets the first element of an array
		list, ok := v.([]any)
		if !ok || len(list) == 0 {
			return nil, false
		}

		next, ok = list[0].(map[string]any)
		if !ok {
			return nil, false
		}
		dto = next
	}
}

// castType returns an error on invalid type
func castType[T any](v any) (T, error) {
	t, ok := v.(T)
	if !ok {
		var empty T
		return empty, fmt.Errorf("invalid type. Expected %T, got %T", empty, v)
	}
	return t, nil
}

func renameAliases(dto map[string]any) {
	for k, v := range aliasFieldsMap() {
		renameAlias(dto, k, v)
	}
}

// renameAlias renames ip_filter_string to ip_filter
func renameAlias(dto map[string]any, path, orig string) {
	keys := strings.Split(path, ".0.")
	for i, k := range keys {
		v, ok := dto[k]
		if !ok {
			// the key does not exist
			return
		}

		// If reached the key
		isLast := len(keys)-1 == i
		if isLast {
			delete(dto, k)
			if !isZero(dto[orig]) {
				// When there multiple choices, like ip_filter_string and ip_filter_object,
				// keeps the one that has a non-empty value.
				// This might happen when a user migrates from ip_filter to ip_filter_string
				v = dto[orig]
			}
			dto[orig] = v
			return
		}

		if a, ok := v.([]any); ok {
			for _, j := range a {
				renameAlias(j.(map[string]any), strings.Join(keys[i+1:], ".0."), orig)
			}
			return
		}
		dto = v.(map[string]any)
	}
}

// isZero returns true for zero-values
func isZero(v any) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Map, reflect.Slice:
		return value.Len() == 0
	}
	return value.IsZero()
}
