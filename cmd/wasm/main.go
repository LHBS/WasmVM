package main

import (
	"io/ioutil"
	"github.com/go-interpreter/wagon/cmd/wasm/wasmservice"
)

func main() {
	data, _ := ioutil.ReadFile("testdata/test.wasm")

	paras := make([]wasmservice.Param,2)
	paras[0] = wasmservice.Param{Ptype:"int32",Pval:"007"}
	paras[1] = wasmservice.Param{Ptype:"string",Pval:"IronMan"}

	args := wasmservice.Params{
		Arg:  paras,
		Addrs: make([]int64,2),
	}

	ws := wasmservice.WasmService{
		Code:      data,
		Args:      args,
		Method:    "init",
	}
	ws.RegisterApi()
	ws.Execute()
}



