// Copyright 2018 The go-ecoball Authors
// This file is part of the go-ecoball.
//
// The go-ecoball is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ecoball is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ecoball. If not, see <http://www.gnu.org/licenses/>.
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

