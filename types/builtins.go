package types

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/meta-quick/mask/anonymity"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/tjfoc/gmsm/x509"
	"strings"
	"time"
)

type (
	Builtin struct {
		Name       string    `json:"name"`
		Decl       *Function `json:"decl"`
		Infix      string    `json:"infix,omitempty"`
		Relation   bool      `json:"relation,omitempty"`
		deprecated bool
	}

	BuiltinContext struct {
		Context context.Context
		Current string      //current content to be handled
		Fn      string      //current mask process function
		Args    []string    //current mask process function
		Result  interface{} //return result
		Err     error       //error if any
	}
	BuiltinFunc func(bctx *BuiltinContext, operands []interface{}) interface{}
)

var Builtins []*Builtin
var BuiltinMap map[string]*Builtin

func RegisterBuiltin(b *Builtin) {
	Builtins = append(Builtins, b)
	BuiltinMap[b.Name] = b
	if len(b.Infix) > 0 {
		BuiltinMap[b.Infix] = b
	}
}

var builtinFunctions = map[string]BuiltinFunc{}

func RegisterFunction(key string, fn BuiltinFunc) {
	builtinFunctions[key] = fn
}

var PFE_MASK_NUM = &Builtin{
	Name: "mx.pfe.mask_number",
	Decl: NewFunction(
		Args(
			N,
			N32,
		),
		N,
	),
}

var PFE_MASK_STR = &Builtin{
	Name: "mx.pfe.mask_string",
	Decl: NewFunction(
		Args(
			S,
			N32,
		),
		S,
	),
}

var DefaultBuiltins = []*Builtin{
	PFE_MASK_NUM,
	PFE_MASK_STR,
	HIDE_MASK_STR,
	HIDE_MASK_BOOLEAN,
	HIDE_MASK_FLOAT32,
	HIDE_MASK_FLOAT64,
	HIDE_MASK_INT32,
	HIDE_MASK_INT64,
	HIDE_MASK_DATESTRING,
	HIDE_MASK_DATE_MSEC,
	HIDE_MASK_STRX,
	FLOOR_MASK_FLOAT64,
	FLOOR_MASK_TIMEINMSEC,
	FLOOR_MASK_TIMESTRING,
	SM2_MASK_STR,
	SM4_MASK_STR,
	PHONE_MASK,
}

var DefaultHandlerBuiltins = map[string]BuiltinFunc{
	PFE_MASK_NUM.Name:               PFE_MASK_NUM_HANDLE,
	PFE_MASK_STR.Name:               PFE_MASK_STR_HANDLE,
	HIDE_MASK_STR.Name:              HIDING_MASK_STR_HANDLE,
	HIDE_MASK_BOOLEAN.Name:          HIDING_MASK_BOOLEAN_HANDLE,
	HIDE_MASK_FLOAT32.Name:          HIDING_MASK_FLOAT32_HANDLE,
	HIDE_MASK_FLOAT64.Name:          HIDING_MASK_FLOAT64_HANDLE,
	HIDE_MASK_INT32.Name:            HIDING_MASK_INT32_HANDLE,
	HIDE_MASK_INT64.Name:            HIDING_MASK_INT64_HANDLE,
	HIDE_MASK_DATESTRING.Name:       HIDING_MASK_TIME_HANDLE,
	HIDE_MASK_DATE_MSEC.Name:        HIDING_MASK_TIMEMSEC_HANDLE,
	HIDE_MASK_STRX.Name:             HIDING_MASK_STR0_HANDLE,
	FLOOR_MASK_FLOAT64.Name:         FLOOR_MASK_FLOAT64_HANDLE,
	FLOOR_MASK_TIMEINMSEC.Name:      FLOOR_MASK_TIMEMSEC_HANDLE,
	FLOOR_MASK_TIMESTRING.Name:      FLOOR_MASK_TIME_HANDLE,
	SM2_MASK_STR.Name:               SM2_MASK_STR_HANDLE,
	SM4_MASK_STR.Name:               SM4_MASK_STR_HANDLE,
	PHONE_MASK.Name:                 PHONE_MASK_HANDLE,
	CUSTOMER_MASK_MD_ID.Name:        CUSTOMER_MASK_MD_ID_HANDLE,
	CUSTOMER_MASK_PHONE_NUMBER.Name: CUSTOMER_MASK_PHONE_NUMBER_HANDLE,
}

func init() {
	BuiltinMap = map[string]*Builtin{}

	for _, b := range DefaultBuiltins {
		RegisterBuiltin(b)
	}

	for k, v := range DefaultHandlerBuiltins {
		RegisterFunction(k, v)
	}
}

func PFE_MASK_NUM_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	pfe := anonymity.NewPrefixPreserveMasker()
	output, _ := pfe.MaskInteger(args[0].(int64), args[1].(int))
	return output
}

func PFE_MASK_STR_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	pfe := anonymity.NewPrefixPreserveMasker()
	output, _ := pfe.MaskString(args[0].(string), args[1].(int))
	return output
}

var HIDE_MASK_STR = &Builtin{
	Name: "mx.hide.mask_string",
	Decl: NewFunction(
		Args(
			S,
			S,
		),
		S,
	),
}

func HIDING_MASK_STR_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskString(args[0].(string), args[1].(string))
	return output
}

var HIDE_MASK_BOOLEAN = &Builtin{
	Name: "mx.hide.mask_boolean",
	Decl: NewFunction(
		Args(
			B,
		),
		B,
	),
}

func HIDING_MASK_BOOLEAN_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskBool(args[0].(bool))
	return output
}

var HIDE_MASK_FLOAT32 = &Builtin{
	Name: "mx.hide.mask_float32",
	Decl: NewFunction(
		Args(
			F,
		),
		F,
	),
}

func HIDING_MASK_FLOAT32_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskFloat32(args[0].(float32))
	return output
}

var HIDE_MASK_FLOAT64 = &Builtin{
	Name: "mx.hide.mask_float64",
	Decl: NewFunction(
		Args(
			F,
		),
		F,
	),
}

func HIDING_MASK_FLOAT64_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskFloat64(args[0].(float64))
	return output
}

var HIDE_MASK_INT32 = &Builtin{
	Name: "mx.hide.mask_int32",
	Decl: NewFunction(
		Args(
			N32,
		),
		N32,
	),
}

func HIDING_MASK_INT32_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskInt(args[0].(int))
	return output
}

var HIDE_MASK_INT64 = &Builtin{
	Name: "mx.hide.mask_int64",
	Decl: NewFunction(
		Args(
			N,
		),
		N,
	),
}

func HIDING_MASK_INT64_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskInt64(args[0].(int64))
	return output
}

func HIDING_MASK_UINT32_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskUint(args[0].(uint))
	return output
}

func HIDING_MASK_UINT64_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskUint64(args[0].(uint64))
	return output
}

var HIDE_MASK_DATESTRING = &Builtin{
	Name: "mx.hide.mask_timestring",
	Decl: NewFunction(
		Args(
			S,
		),
		S,
	),
}

func HIDING_MASK_TIME_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	tt, _ := time.Parse("2006-01-02 15:04:05", args[0].(string))
	output, _ := hiding.MaskTime(&tt)
	return output.Format("2006-01-02 15:04:05")
}

var HIDE_MASK_DATE_MSEC = &Builtin{
	Name: "mx.hide.mask_timemesc",
	Decl: NewFunction(
		Args(
			N,
		),
		N,
	),
}

func HIDING_MASK_TIMEMSEC_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	tt := time.UnixMilli(args[0].(int64))
	output, _ := hiding.MaskTime(&tt)
	return output.UnixMilli()
}

var HIDE_MASK_STRX = &Builtin{
	Name: "mx.hide.mask_strx",
	Decl: NewFunction(
		Args(
			S,
			S,
			N32,
			N32,
		),
		S,
	),
}

func HIDING_MASK_STR0_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	hiding := anonymity.NewHidingMasker()
	output, _ := hiding.MaskString0(args[0].(string), args[1].(string), args[2].(int), args[3].(int))
	return output
}

var FLOOR_MASK_FLOAT32 = &Builtin{
	Name: "mx.floor.mask_float32",
	Decl: NewFunction(
		Args(
			S,
			F,
		),
		F,
	),
}

func FLOOR_MASK_FLOAT32_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	floor := anonymity.NewFloorMasker()
	output, _ := floor.MaskFloat32(args[0].(float32))
	return output
}

var FLOOR_MASK_FLOAT64 = &Builtin{
	Name: "mx.floor.mask_float64",
	Decl: NewFunction(
		Args(
			F,
		),
		N,
	),
}

func FLOOR_MASK_FLOAT64_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	floor := anonymity.NewFloorMasker()
	output, _ := floor.MaskFloat64(args[0].(float64))
	return output
}

var FLOOR_MASK_TIMESTRING = &Builtin{
	Name: "mx.floor.mask_timestring",
	Decl: NewFunction(
		Args(
			S,
			S,
		),
		S,
	),
}

func FLOOR_MASK_TIME_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	floor := anonymity.NewFloorMasker()
	tt, _ := time.Parse("2006-01-02 15:04:05", args[0].(string))
	output, _ := floor.MaskTime(tt, args[1].(string))
	return output.Format("2006-01-02 15:04:05")
}

var FLOOR_MASK_TIMEINMSEC = &Builtin{
	Name: "mx.floor.mask_time_msec",
	Decl: NewFunction(
		Args(
			N,
			S,
		),
		N,
	),
}

func FLOOR_MASK_TIMEMSEC_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	floor := anonymity.NewFloorMasker()
	tt := time.UnixMilli(args[0].(int64))
	output, _ := floor.MaskTime(tt, args[1].(string))
	return output.UnixMilli()
}

var SM2_MASK_STR = &Builtin{
	Name: "mx.sm2.mask_string",
	Decl: NewFunction(
		Args(
			S,
			S,
		),
		S,
	),
}

func SM2_MASK_STR_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	data := args[0].(string)
	publickKey := args[1].(string)

	if !strings.Contains(publickKey, "PUBLIC KEY") {
		publickKey = "-----BEGIN PUBLIC KEY-----\r\n" +
			publickKey + "\r\n" +
			"-----END PUBLIC KEY-----"
	}

	d2 := []byte(publickKey)
	pubMen, err := x509.ReadPublicKeyFromPem(d2)
	if err != nil {
		return data
	}

	msg := []byte(data)
	ciphertxt, err := sm2.Encrypt(pubMen, msg, nil, sm2.C1C3C2)
	if err != nil {
		return data
	}

	return base64.StdEncoding.EncodeToString(ciphertxt)
}

var SM4_MASK_STR = &Builtin{
	Name: "mx.sm4.mask_string",
	Decl: NewFunction(
		Args(
			S,
			S,
		),
		S,
	),
}

func SM4_MASK_STR_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	data := args[0].(string)
	key := args[1].(string)

	r, err := sm4.Sm4Ecb([]byte(key), []byte(data), true) //sm4Ecb模式pksc7填充加密
	if err != nil {
		return data
	}
	return base64.StdEncoding.EncodeToString(r)
}

var PHONE_MASK = &Builtin{
	Name: "mx.phone.mask_string",
	Decl: NewFunction(
		Args(
			S,
			N32,
			N32,
		),
		S,
	),
}

var CUSTOMER_MASK_MD_ID = &Builtin{
	Name: "mx.customer.mask_md_id",
	Decl: NewFunction(
		Args(),
		S,
	),
}

func CUSTOMER_MASK_MD_ID_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	inputS := bctx.Current
	inputMap := make(map[string]interface{})
	err := sonic.Unmarshal([]byte(inputS), &inputMap)
	if err != nil {
		return nil
	}
	userId := inputMap["sub"]

	bytes, err := convertor.ToBytes(userId)
	if err != nil {
		return nil
	}
	sum := md5.Sum(bytes)
	return convertor.ToString(sum)
}

var CUSTOMER_MASK_PHONE_NUMBER = &Builtin{
	Name: "mx.customer.mask_phone_number",
	Decl: NewFunction(
		Args(),
		S,
	),
}

func CUSTOMER_MASK_PHONE_NUMBER_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	inputS := bctx.Current
	inputMap := make(map[string]interface{})
	err := sonic.Unmarshal([]byte(inputS), &inputMap)
	if err != nil {
		return nil
	}
	userId := inputMap["sub"].(string)

	return generatePhoneNumber(userId)
}

var keyArr = []int{1, 7, 7, 9, 1, 8, 5, 0, 6, 0, 4, 1, 2, 9, 8, 4, 1, 5, 7, 6}
var shuffleArr = []int{7, 5, 4, 6, 2, 1, 3, 0, 9, 8}

func PHONE_MASK_HANDLE(bctx *BuiltinContext, args []interface{}) interface{} {
	phone := args[0].(string)
	start := args[1].(int)
	end := args[2].(int)

	if end <= start {
		return phone
	}
	if end > len(keyArr) {
		return phone
	}
	// 移除非数字字符
	digits := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, phone)
	if len(digits) < len(phone) {
		return phone
	}
	if len(digits) < end {
		return phone
	}
	runes := []rune(digits)
	// 旋转加数
	for i := range runes {
		if i >= start && i < end {
			temp := int(runes[i]) - '0'
			runes[i] = rune((temp+keyArr[i])%10 + '0')
		}
	}
	// 换位
	shuffleString(runes, start, end)
	return string(runes)
}

func shuffleString(runes []rune, start int, end int) {
	for i := range runes {
		if i >= start && i < end {
			temp := int(runes[i]) - '0'
			runes[i] = rune(shuffleArr[temp] + '0')
		}
	}
}

// 简化的哈希函数，将用户名转换为一个整数
func simpleHash(username string) int {
	hash := 0
	for _, char := range username {
		hash = (hash*31 + int(char)) % 100000000
	}
	return hash
}

func generatePhoneNumber(username string) string {
	// 定义支持的手机号前缀
	phonePrefixes := []string{}
	for i := 0; i < 10; i++ {
		phonePrefixes = append(phonePrefixes, fmt.Sprintf("13%d", i))
		phonePrefixes = append(phonePrefixes, fmt.Sprintf("15%d", i))
		phonePrefixes = append(phonePrefixes, fmt.Sprintf("18%d", i))
	}

	// 使用简化的哈希函数将用户名转换为整数
	hashCode := simpleHash(username)

	// 使用哈希值的前2位来选择前缀
	prefixIndex := hashCode % len(phonePrefixes)
	phonePrefix := phonePrefixes[prefixIndex]

	// 取哈希值的后8位作为手机号的后8位
	phoneSuffix := fmt.Sprintf("%08d", hashCode)

	// 将前缀和后缀组合成一个完整的手机号
	phoneNumber := phonePrefix + phoneSuffix

	return phoneNumber
}
