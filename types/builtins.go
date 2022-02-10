package types

import (
	"context"
	"github.com/meta-quick/mask/anonymity"
)

type (
	Builtin struct {
		Name       string          `json:"name"`
		Decl       *Function `json:"decl"`
		Infix      string          `json:"infix,omitempty"`
		Relation   bool            `json:"relation,omitempty"`
		deprecated bool
	}

	BuiltinContext struct {
		Context                context.Context
		Current                string  //current content to be handled
		Fn                     string       //current mask process function
		Args                   []string//current mask process function
		Result				   interface{}  //return result
		Err                    error        //error if any
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

func RegisterFunction(key string,fn BuiltinFunc) {
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


var DefaultBuiltins = []*Builtin {
	PFE_MASK_NUM,
	PFE_MASK_STR,
}

var DefaultHandlerBuiltins = map[string]BuiltinFunc {
	PFE_MASK_NUM.Name: PFE_MASK_NUM_HANDLE,
	PFE_MASK_STR.Name: PFE_MASK_STR_HANDLE,
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

func PFE_MASK_NUM_HANDLE(bctx *BuiltinContext,args []interface{}) interface{} {
	pfe := anonymity.NewPrefixPreserveMasker()
	output,_ :=	pfe.MaskInteger(args[0].(int64),args[1].(int))
	return output
}

func PFE_MASK_STR_HANDLE(bctx *BuiltinContext,args []interface{}) interface{} {
	pfe := anonymity.NewPrefixPreserveMasker()
	output,_ :=	pfe.MaskString(args[0].(string),args[1].(int))
	return output
}