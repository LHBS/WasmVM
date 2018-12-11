package wasmservice

import(
	"fmt"
	"github.com/go-interpreter/wagon/exec"
	"bytes"
)

//C API :void prints(char *p)
func (ws *WasmService) prints(proc *exec.Process, p uint32) int32{
	data := proc.LoadAt(int(p))
	length := bytes.IndexByte(data,0)
	fmt.Printf("%s\n",data[:length])
	return 0
}

//C API :void prints_l(char *p, uint32 len)
func (ws *WasmService) prints_l(proc *exec.Process, p uint32, length uint32) int32{
	msg := make([]byte, length)
	proc.ReadAt(msg,int(p), int(length))
	fmt.Printf("%s\n",msg)
	return 0
}

//C API :void printi(int64 v)
func (ws *WasmService) printi(proc *exec.Process, v int64) uint32{
	fmt.Printf("%d\n",v)
	return 0
}

//C API :void printui(uint64 v)
func (ws *WasmService) printui(proc *exec.Process, v uint64) uint32{
	fmt.Printf("%d\n",v)
	return 0
}

//C API :void printsf(float v)
func (ws *WasmService) printsf(proc *exec.Process, v float32) uint32{
	fmt.Printf("%f\n",v)
	return 0
}

//C API :void printdf(double v)
func (ws *WasmService) printdf(proc *exec.Process, v float64) uint32{
	fmt.Printf("%f\n",v)
	return 0
}

