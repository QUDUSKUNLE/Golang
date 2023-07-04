package interfaces

import (
	"fmt"
	"reflect"
)

func DisPlayType(interf interface{}) interface{} {
	fmt.Println(reflect.TypeOf(interf))
	switch theType := interf.(type) {
	case int: 
		return "This is a int " + fmt.Sprintf("%v", theType)
	case float64:
		return "This is a float64 " + fmt.Sprintf("%v", theType)
	case float32:
		return "This is a float32 " + fmt.Sprintf("%v", theType)
	case string:
		return "This is a string " + theType
	default:
		return "unknown"
	}
}
