package behavioral

import "reflect"

// find value in string array of map
func in_array(search interface{}, array interface{}) bool {
	val := reflect.ValueOf(array)
	val.Convert(val.Type())

	switch array.(type) {
	case map[string][]string:
		s := val.MapRange()
		for s.Next() {
			s.Value().Convert(s.Value().Type())
			for i := 0; i < s.Value().Len(); i++ {
				if reflect.DeepEqual(search, s.Value().Index(i).Interface()) {
					return true
				}
			}
		}
	}
	return false
}