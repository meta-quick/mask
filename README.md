---
modified: 2022-02-16T11:07:39.837Z
title: null
---

脱敏方法定义

1. 前缀保留

- Name: mx.pfe.mask_string
   - Args:  
      1. 保留位数

   - 说明: 字符串处理函数，保留指定位数前缀
  

- Name: mx.pfe.mask_number
   - Args:  
      1. 保留位数

   - 说明: 字符串处理数值，保留指定位数前缀
  

2. 遮盖算法

- Name: mx.hide.mask_string
   - Args:  
      1. 用于遮盖的字符串

   - 说明: 字符串处理函数，字符串遮盖

- Name: mx.hide.mask_boolean
   - Args:  //无参数

   - 说明: 字符串处理Boolean，返回true

- Name: mx.hide.mx.hide.mask_float64
   - Args:  //无参数

   - 说明: 浮点数遮盖

- Name: mx.hide.mask_int32
   - Args:  //无参数

   - 说明: int32 遮盖

- Name: mx.hide.mask_int64
   - Args:  //无参数

   - 说明: int64 遮盖

- Name: mx.hide.mask_timestring
   - Args:  //无参数

   - 说明: 时间字符串遮盖


- Name: mx.hide.mask_timemesc
   - Args:  //无参数

   - 说明: 毫秒时间遮盖


- Name: mx.hide.mask_strx
   - Args: 
      1. //用于遮盖的字符串
      2. //起始位置
      3. //结束位置

   - 说明: 字符串指定位置遮盖


3. 取整算法

- Name: mx.floor.mask_float64
   - Args: //无参数

   - 说明: float64

- Name: mx.floor.mask_time_msec
   - Args: 
     1. 取整格式 //YMDHms, 分别表示年、月、日、时、分、秒

   - 说明: 毫秒时间取整

- Name: mx.floor.mask_timestring
   - Args: 
     1. 取整格式 //YMDHms, 分别表示年、月、日、时、分、秒

   - 说明: 字符串时间取整


使用说明
```go
//输入结果
input := `
{
  "a": {
	"b": {
	  "c": "100"
	}
  },
  "d": {
	"e": {
	  "f": "1000xx",
      "g": [
          {"x":100},
          {"x":200}
       ]
	}
  }
}
`
//定义模型
model := `
{
   "filters" : {
      "denied": [
          "a/b/c" 
       ]
   },
   "shuffle" : {
      "d/e/f" : {
         "mx.pfe.mask_string": [ 
            "2"
          ]
      },
	 "d/e/g/:/x" : {
         "mx.pfe.mask_string": [ 
            "1"
          ]
      }
   }
}
`
//加载模型
topdown.ShuffleModelAddString("/api",model)

//定义规则
rule := `
		package test
		p = output { 
          output := json.shuffle(input,"/api",[{"op": "add", "path": "/a/bar", "value": 2}])
}
`

//调用规则
var v interface{}
sonic.Unmarshal([]byte(input), &v)

r := New(
    Query("data.test.p"),
    Module("", module),
)

pq, err := r.PrepareForEval(ctx)
if err != nil {
    t.Fatalf("Unexpected error: %s", err.Error())
}

if err != nil {
    t.Fatal(err)
}

output, err := pq.Eval(ctx,EvalInput(v))
ret,_ :=sonic.Marshal(output)
fmt.Println(string(ret[:]))
```
