package reflect

import (
	"reflect"
	"testing"
)

type Goods struct {
	Id   int    `json:"prod_id"`
	Sku  string `json:"sku"`
	Name string `json:"prod_name"`
}

func TestTag(t *testing.T) {
	g := Goods{
		Id:   100,
		Sku:  "C31001",
		Name: "IphoneXs",
	}
	gt := reflect.TypeOf(g)
	gv := reflect.ValueOf(g)
	// reflect.Type类型的FieldByName(name)方法，返回field, ok两个值
	// reflect.Value类型的FieldByName(name)获取其值

	if id, ok := gt.FieldByName("Id"); !ok {
		t.Log("field not exist")
	} else {
		t.Logf("Id=%d, Tag=%s", gv.FieldByName("Id").Int(), id.Tag)
	}
	if name, ok := gt.FieldByName("Name"); ok {
		t.Log("Name=", gv.FieldByName("Name"))
		t.Log("Tag=", name.Tag.Get("json"))
	}
}
