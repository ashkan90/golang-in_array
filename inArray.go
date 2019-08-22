package exists

import (
	"reflect"
	"strings"
)

func In_array(search interface{}, array interface{}, deep bool) bool {

	val := reflect.ValueOf(array)
	val = val.Convert(val.Type())

	typ := reflect.TypeOf(array).Kind()

	switch typ {
	case reflect.Map:
		s := val.MapRange()

		for s.Next() {
			s.Value().Convert(s.Value().Type())
			for i := 0; i < s.Value().Len(); i++ {
				if deep {
					if reflect.DeepEqual(search, s.Value().Index(i).Interface()) {
						return true
					}
				} else {
					str := s.Value().Index(i).String()
					if strings.Contains(str, search.(string)) {
						return true
					}
				}
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(search, val.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}
