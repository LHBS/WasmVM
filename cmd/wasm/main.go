package main

import (
	"io/ioutil"
	"github.com/go-interpreter/wagon/cmd/wasm/wasmservice"
)

func main() {
	data, _ := ioutil.ReadFile("testdata/test.wasm")

	paras := make([]wasmservice.Param,2)
	paras[0] = wasmservice.Param{Type:"int32",Val:"007"}
	paras[1] = wasmservice.Param{Type:"string",Val:"IronMan"}

	args := wasmservice.Params{
		Data:  paras,
		Addr:  make([]int64,2),
	}

	ws := wasmservice.WasmService{
		Code:      data,
		Args:      args,
		Method:    "init",
	}
	ws.RegisterApi()
	ws.Execute()
}



