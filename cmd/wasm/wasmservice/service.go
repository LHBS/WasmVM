package wasmservice
import (
	"bytes"
	"errors"
	"os"
	"strconv"
	"fmt"
	"github.com/go-interpreter/wagon/exec"
	"github.com/go-interpreter/wagon/util"
	"github.com/go-interpreter/wagon/validate"
	"github.com/go-interpreter/wagon/wasm"
)

type Params struct{
	Data    []Param
	Addr    []int64
}

type Param struct {
	Type string
	Val  string
}

type WasmService struct {
	Code      []byte
	Args      Params
	Method    string
}

func (ws *WasmService) ParseParam(vm *exec.VM)([]uint64, error){
	method, err := vm.Memmanage.SetBlock(ws.Method)
	if err != nil{
		return nil,err
	}
	params := make([]uint64,1)
	params[0] = uint64(method)

	pCount := len(ws.Args.Data)
	for index := 0; index < pCount; index++ {
		switch ws.Args.Data[index].Type {
		case "string":
			vString := ws.Args.Data[index].Val
			addr, err := vm.Memmanage.SetBlock(vString)
			if err != nil{
				return nil, errors.New("store string error")
			}
			ws.Args.Addr[index] = int64(addr)

		case "int8","int16","int32","int64":
			vString := ws.Args.Data[index].Val
			vInt64, _ := strconv.ParseInt(vString, 10, 64)
			ws.Args.Addr[index] = vInt64

		case "uint8","uint16","uint32":
			vString := ws.Args.Data[index].Val
			vUint64, _ := strconv.ParseUint(vString, 10, 64)
			ws.Args.Addr[index] = int64(vUint64)

		default:
			return nil, errors.New("unSupport type")

		}
	}
	return params,nil
}

func (ws *WasmService) Execute() ([]byte, error) {
	bf := bytes.NewBuffer(ws.Code)
	method := "apply"

	m, err := wasm.ReadModule(bf, importer)
	if err != nil {
		fmt.Printf("could not read module: %v", err)
		return nil, err
	}

	if m.Export == nil {
		fmt.Printf("module has no export section")
	}

	vm, err := exec.NewVM(m)
	if err != nil {
		fmt.Printf("could not create VM: %v", err)
		return nil, err
	}

	entry, ok := m.Export.Entries[method]
	if ok == false {
		fmt.Printf("method does not exist!")
		return nil, err
	}

	args, err:= ws.ParseParam(vm)

	if err != nil{
		fmt.Printf("parse parameter error!")
		return nil, err
	}
	index := int64(entry.Index)
	fIdx := m.Function.Types[int(index)]
	fType := m.Types.Entries[int(fIdx)]

	res, err := vm.ExecCode(index, args...)
	if err != nil {
		fmt.Printf("%v\n",err)
		return nil, err
	}
	switch fType.ReturnTypes[0] {
	case wasm.ValueTypeI32:
		return util.Int32ToBytes(res.(uint32)), nil
	case wasm.ValueTypeI64:
		return util.Int64ToBytes(res.(uint64)), nil
	case wasm.ValueTypeF32:
		return util.Float32ToBytes(res.(float32)), nil
	case wasm.ValueTypeF64:
		return util.Float64ToBytes(res.(float64)), nil
	default:
		return nil, errors.New("unknown return type")
	}
}

func importer(name string) (*wasm.Module, error) {
	f, err := os.Open(name + ".wasm")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, err := wasm.ReadModule(f, nil)
	if err != nil {
		return nil, err
	}
	err = validate.VerifyModule(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (ws *WasmService) RegisterApi() {
	functions := wasm.InitNativeFuns()
	//memory
	functions.Register("memset", ws.memset)
	functions.Register("memcpy", ws.memcpy)
	//console
	functions.Register("prints", ws.prints)
	functions.Register("prints_l", ws.prints_l)
	functions.Register("printi", ws.printi)
	functions.Register("printui", ws.printui)
	functions.Register("printsf", ws.printsf)
	functions.Register("printdf", ws.printdf)
	//system
	functions.Register("assert", ws.assert)
	functions.Register("exit", ws.exit)
	//runtime
	functions.Register("read_param", ws.read_param)
}