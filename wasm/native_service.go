package wasm

import (
	"reflect"
)

//TODO add in VM
var  Funs  *NativeFuns

type NativeFuns struct{
	funmap map[string]reflect.Value
}

func InitNativeFuns() *NativeFuns{
	fun := NativeFuns{make(map[string]reflect.Value)}
	Funs = &fun
	return Funs
}

func GetFuns() *NativeFuns{
	return Funs
}
func (n *NativeFuns) Register(name string, i interface{}) bool{
	if _, ok := n.funmap[name]; ok {
		return false
	}
	value := reflect.ValueOf(i)
	n.funmap[name] = value
	return true
}

func (n *NativeFuns) GetValue (name string) reflect.Value{
	if value, ok := n.funmap[name]; ok {
		return value
	}
	return reflect.Value{}
}

