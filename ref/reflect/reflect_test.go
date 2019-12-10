package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndVal(t *testing.T) {
	var f float64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
	t.Log(reflect.ValueOf(f).Kind())
	//t.Log(reflect.ValueOf(f).Int())	// panic
	//t.Log(reflect.ValueOf(f).String())
	//t.Log(reflect.ValueOf(f).Bytes())
}
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.String:
		fmt.Println("String")
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Int")
	default:
		fmt.Println("Unknown Type")
	}
}

func TestCheckType(t *testing.T) {
	CheckType(12)
	CheckType(12.5)
	CheckType("12")
}
