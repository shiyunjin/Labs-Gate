package util

import "reflect"

func In_array(val interface{}, array interface{}) (exists bool) {
	exists = false
	//index := -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				//index = i
				exists = true
				return
			}
		}
	}

	return
}