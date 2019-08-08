package exists

import (
	"reflect"
)

//func main() {
//	rules := map[string][]string{
//		"name":    {"required", "min-len:20"},
//		"surname": {"required", "min-len:15"},
//	}
//	//k := []string{"required", "min-len:20"}
//	fmt.Print(in_array("required", rules))
//}

// find value in string array of map
//func in_array(search interface{}, array interface{}) bool { // recursive call made.
//	exists := false
//
//	val := reflect.ValueOf(array)
//	val = val.Convert(val.Type())
//
//	typ := reflect.TypeOf(array).Kind()
//
//	switch typ {
//	case reflect.Map:
//		s := val.MapRange()
//		for s.Next() {
//			s.Value().Convert(s.Value().Type())
//
//			switch reflect.TypeOf(s.Value().Interface()).Kind() {
//			case reflect.Slice, reflect.Array:
//				in_array(search, s.Value().Interface())
//			default:
//				fmt.Println(reflect.DeepEqual(search, s.Value().Interface()))
//				if reflect.DeepEqual(search, s.Value().Interface()) {
//					exists = true
//				}
//			}
//		}
//	case reflect.Slice, reflect.Array:
//		for i := 0; i < val.Len(); i++ {
//			if reflect.DeepEqual(search, val.Index(i).Interface()) {
//				exists = true
//			}
//		}
//	}
//
//	return exists
//}

func In_array(search interface{}, array interface{}) bool {

	val := reflect.ValueOf(array)
	val = val.Convert(val.Type())

	typ := reflect.TypeOf(array).Kind()

	switch typ {
	case reflect.Map:
		s := val.MapRange()
		for s.Next() {
			s.Value().Convert(s.Value().Type())

			for i := 0; i < val.Len(); i++ {
				if reflect.DeepEqual(search, val.Index(i).Interface()) {
					return true
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
