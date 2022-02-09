package anonymity

import (
	"fmt"
	"github.com/rkritchat/jsonmask"
	"testing"
)


var j = []byte(`{"foo":1,"bar":2,"baz":[3,4],"phoneNo":123456789, "newField":"test", "userInfo":{"firstname":"Kritchat", "lastname": "Rojanaphruk"}}`)

func Test_mask(tt *testing.T) {
	m := jsonmask.Init([]string{"newField","foo","bar","baz"}) //optional
	t, err := m.Json(j)
	if err != nil {
		panic(err)
	}
	fmt.Println(*t)
}
