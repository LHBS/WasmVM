package wasmservice
import (
	"github.com/go-interpreter/wagon/exec"
)

//for c api: int32 read_param(int32 index)
func(ws *WasmService)read_param(proc *exec.Process, index int32) int32{
	if int(index) > len(ws.Args.Addrs) || index < 1{
		return -1
	}
	data := ws.Args.Addrs[index-1]
	return int32(data)
}
