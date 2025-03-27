package anonymity

import (
	"fmt"
	"github.com/meta-quick/mask/utils"
	"github.com/rkritchat/jsonmask"
	"github.com/stretchr/testify/assert"
	"testing"
)

var j = []byte(`{"foo":1,"bar":2,"baz":[3,4],"phoneNo":123456789, "newField":"test", "userInfo":{"firstname":"Kritchat", "lastname": "Rojanaphruk"}}`)

func Test_mask(tt *testing.T) {
	m := jsonmask.Init([]string{"newField", "foo", "bar", "baz"}) //optional
	t, err := m.Json(j)
	if err != nil {
		panic(err)
	}
	fmt.Println(*t)
}

func Test_mask1(tt *testing.T) {
	assert.Equal(tt, "张*", utils.MaskString("张三", 1, 1, "*"))
	assert.Equal(tt, "张*三", utils.MaskString("张三三", 1, 1, "*"))
	assert.Equal(tt, "张三三三三三", utils.MaskString("张三三三三三", 6, 1, "*"))
	assert.Equal(tt, "张*三三三三", utils.MaskString("张三三三三三", 1, 6, "*"))
	assert.Equal(tt, "张*****", utils.MaskString("张三三三三三", 1, -1, "*"))
	assert.Equal(tt, "******", utils.MaskString("张三三三三三", 0, 0, "*"))
	assert.Equal(tt, "张三三三三三", utils.MaskString("张三三三三三", 0, 6, "*"))
}
