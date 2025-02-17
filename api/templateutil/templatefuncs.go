package templateutil

import (
	"errors"
	"fmt"
	"strings"
)

func GetFromMap(m map[string]any, key string) (any, error) {
	val, ok := m[key]
	if !ok {
		return nil, errors.New("missing key " + key)
	}
	return val, nil
}

func TmplMap(pairs ...any) (map[string]any, error) {
	if len(pairs)%2 == 1 {
		return nil, errors.New("Map: want key value pairs")
	}

	m := make(map[string]any, len(pairs)/2)
	for i := 0; i < len(pairs); i += 2 {
		key := fmt.Sprintf("%v", pairs[i])
		value := pairs[i+1]

		m[key] = value
	}
	return m, nil
}

func TmplSlice(elements ...any) []any {
	return elements
}

func WithComData(pairs ...any) (map[string]any, error) {
	if len(pairs)%2 == 1 {
		return nil, errors.New("WithComData: want key value pairs")
	}

	comData := make(map[string]any, len(pairs)/2)
	for i := 0; i < len(pairs); i += 2 {
		key := fmt.Sprintf("%v", pairs[i])
		value := pairs[i+1]

		if key == "ComData" {
			comData, ok := value.(map[string]any)
			if ok {
				for key, value := range comData {
					comData[key] = value
				}

				continue
			}
		}

		comData[key] = value
	}
	return comData, nil
}

// Splice string
func TmplSubstr(s string) string {
	if len(s) < 30 {
		return s
	}
	s = strings.Join(strings.Fields(s), " ")
	return fmt.Sprintf("%s...", s[:30])
}
