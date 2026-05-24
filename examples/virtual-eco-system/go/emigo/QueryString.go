// A copy from original work: https://github.com/derekstavis/go-qs

package emigo

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func MarshalQs(hash map[string]interface{}) (string, error) {
	return buildNestedQuery(hash, "")
}

func buildNestedQuery(value interface{}, prefix string) (string, error) {
	components := ""

	switch vv := value.(type) {
	case []interface{}:
		for i, v := range vv {
			component, err := buildNestedQuery(v, prefix+"[]")

			if err != nil {
				return "", err
			}

			components += component

			if i < len(vv)-1 {
				components += "&"
			}
		}

	case map[string]interface{}:
		length := len(vv)

		for k, v := range vv {
			childPrefix := ""

			if prefix != "" {
				childPrefix = prefix + "[" + url.QueryEscape(k) + "]"
			} else {
				childPrefix = url.QueryEscape(k)
			}

			component, err := buildNestedQuery(v, childPrefix)

			if err != nil {
				return "", err
			}

			components += component
			length -= 1

			if length > 0 {
				components += "&"
			}
		}

	case string:
		if prefix == "" {
			return "", fmt.Errorf("value must be a map[string]interface{}")
		}

		components += prefix + "=" + url.QueryEscape(vv)

	default:
		components += prefix
	}

	return components, nil
}

var nameRegex = regexp.MustCompile(`\A[\[\]]*([^\[\]]+)\]*`)
var objectRegex1 = regexp.MustCompile(`^\[\]\[([^\[\]]+)\]$`)
var objectRegex2 = regexp.MustCompile(`^\[\](.+)$`)

func UnmarshalQs(qs string) (map[string]interface{}, error) {
	components := strings.Split(qs, "&")
	params := map[string]interface{}{}

	for _, c := range components {

		tuple := strings.SplitN(c, "=", 2)
		for i, item := range tuple {
			if unesc, err := url.QueryUnescape(item); err == nil {
				tuple[i] = unesc
			}
		}

		key := ""

		if len(tuple) > 0 {
			key = tuple[0]
		}

		value := interface{}(nil)

		if len(tuple) > 1 {
			value = tuple[1]
		}

		if err := normalizeParams(params, key, value); err != nil {
			return nil, err
		}
	}

	return params, nil
}

func normalizeParams(params map[string]interface{}, key string, value interface{}) error {
	after := ""

	if pos := nameRegex.FindIndex([]byte(key)); len(pos) == 2 {
		after = key[pos[1]:]
	}

	matches := nameRegex.FindStringSubmatch(key)
	if len(matches) < 2 {
		return nil
	}

	k := matches[1]
	if after == "" {
		params[k] = value
		return nil
	}

	if after == "[]" {
		ival, ok := params[k]

		if !ok {
			params[k] = []interface{}{value}
			return nil
		}

		array, ok := ival.([]interface{})

		if !ok {
			return fmt.Errorf("Expected type '[]interface{}' for key '%s', but got '%T'", k, ival)
		}

		params[k] = append(array, value)
		return nil
	}

	object1Matches := objectRegex1.FindStringSubmatch(after)
	object2Matches := objectRegex2.FindStringSubmatch(after)

	if len(object1Matches) > 1 || len(object2Matches) > 1 {
		childKey := ""

		if len(object1Matches) > 1 {
			childKey = object1Matches[1]
		} else if len(object2Matches) > 1 {
			childKey = object2Matches[1]
		}

		if childKey != "" {
			ival, ok := params[k]

			if !ok {
				params[k] = []interface{}{}
				ival = params[k]
			}

			array, ok := ival.([]interface{})

			if !ok {
				return fmt.Errorf("Expected type '[]interface{}' for key '%s', but got '%T'", k, ival)
			}

			if length := len(array); length > 0 {
				if hash, ok := array[length-1].(map[string]interface{}); ok {
					if _, ok := hash[childKey]; !ok {
						normalizeParams(hash, childKey, value)
						return nil
					}
				}
			}

			newHash := map[string]interface{}{}
			normalizeParams(newHash, childKey, value)
			params[k] = append(array, newHash)

			return nil
		}
	}

	ival, ok := params[k]

	if !ok {
		params[k] = map[string]interface{}{}
		ival = params[k]
	}

	hash, ok := ival.(map[string]interface{})

	if !ok {
		return fmt.Errorf("Expected type 'map[string]interface{}' for key '%s', but got '%T'", k, ival)
	}

	if err := normalizeParams(hash, after, value); err != nil {
		return err
	}

	return nil
}
