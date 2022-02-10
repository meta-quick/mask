package jsonmask

import (
	"fmt"
	"github.com/meta-quick/mask/types"
	xxmask "github.com/rkritchat/jsonmask"
	"testing"
)

func TestMask(t *testing.T) {
	ctx := &types.BuiltinContext{
		Current: "10002122102",
		Fn: types.PFE_MASK_NUM.Name,
		Args: []string{
			"6",
		},
	}

	types.Eval(ctx)
	fmt.Println(ctx.Result)
}

var j = []byte(`{"foo":1,"bar":2,"baz":[3,4],"phoneNo":123456789, "newField":"test", "userInfo":{"firstname":"Kritchat", "lastname": "Rojanaphruk"}}`)


func TestJsonMask(tt *testing.T	) {
	m := NewMasker()

	handles := map[string]ProcessHandle{
		"newField": ProcessHandle{
			Fn: types.PFE_MASK_STR.Name,
			Args: []string{
				"3",
			},
		},
	}

	m.Init(handles)

	t, err := m.Process(j)
	if err != nil {
		panic(err)
	}
	fmt.Println(*t)
}

func Test3rdMask(tt *testing.T	) {
	m := xxmask.Init([]string{"newField"}) //optional
	t, err := m.Json(j)
	if err != nil {
		panic(err)
	}
	fmt.Println(*t)
}
