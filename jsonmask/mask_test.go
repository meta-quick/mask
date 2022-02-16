package jsonmask

import (
	"fmt"
	"github.com/meta-quick/mask/types"
	xxmask "github.com/rkritchat/jsonmask"
	"testing"
	"time"
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

var j = []byte(`{
  "f1": "xsccxfdadf",
  "f2": 12345,
  "f3": "sexy",
  "f4":  false,
  "f5":  9796.435,
  "f6":  65535,
  "f7":  65535,
  "f8":  "2021-01-02 15:04:05",
  "f9":  1645004457408,
  "f10":  "1234567890987654321",
  "f11":  77090.2345,
  "f12":  1645004457408,
  "f13":  "2021-01-02 15:04:05",
  "fxx": "xxyymmxx"
}`)


func TestJsonMask(tt *testing.T	) {
	m := NewMasker()

	fmt.Println(time.Now().UnixMilli())

	handles := map[string]ProcessHandle{
		"f1": ProcessHandle{
			Fn: types.PFE_MASK_STR.Name, //mx.pfe.mask_string
			Args: []string{
				"3",  //保留位数
			},
		},
		"f2": ProcessHandle{
			Fn: types.PFE_MASK_NUM.Name, //mx.pfe.mask_number
			Args: []string{
				"3", //保留位数
			},
		},
		"f3": ProcessHandle{
			Fn: types.HIDE_MASK_STR.Name, //mx.hide.mask_string
			Args: []string{
				"#####", //用于遮盖的字符串
			},
		},
		"f4": ProcessHandle{
			Fn: types.HIDE_MASK_BOOLEAN.Name, //mx.hide.mask_boolean
			Args: []string{
				//无参数
			},
		},
		"f5": ProcessHandle{
			Fn: types.HIDE_MASK_FLOAT64.Name, //mx.hide.mask_float64
			Args: []string{
				//无参数
			},
		},
		"f6": ProcessHandle{
			Fn: types.HIDE_MASK_INT32.Name, //mx.hide.mask_int32
			Args: []string{
				//无参数
			},
		},
		"f7": ProcessHandle{
			Fn: types.HIDE_MASK_INT64.Name, //mx.hide.mask_int64
			Args: []string{
				//无参数
			},
		},
		"f8": ProcessHandle{
			Fn: types.HIDE_MASK_DATESTRING.Name, //mx.hide.mask_timestring
			Args: []string{
				//无参数
			},
		},
		"f9": ProcessHandle{
			Fn: types.HIDE_MASK_DATE_MSEC.Name, //mx.hide.mask_timemesc
			Args: []string{
			},
		},
		"f10": ProcessHandle{
			Fn: types.HIDE_MASK_STRX.Name, //mx.hide.mask_strx
			Args: []string{
				"#####", //用于遮盖的字符串
				"4",     //起始位置
				"10",    //结束位置
			},
		},
		"f11": ProcessHandle{
			Fn: types.FLOOR_MASK_FLOAT64.Name, //mx.floor.mask_float64
			Args: []string{
				//无参数
			},
		},
		"f12": ProcessHandle{
			Fn: types.FLOOR_MASK_TIMEINMSEC.Name, //mx.floor.mask_time_msec
			Args: []string{
               "YMD", //时间格式
			},
		},
		"f13": ProcessHandle{
			Fn: types.FLOOR_MASK_TIMESTRING.Name, //mx.floor.mask_timestring
			Args: []string{
		      "YMDHms",//时间格式
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
