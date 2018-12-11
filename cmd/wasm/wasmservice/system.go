package wasmservice
import(
	"github.com/go-interpreter/wagon/exec"
)

//C API :void assert(uint32 expression, const char *msg)
func (ws *WasmService) assert(proc *exec.Process, exp, msg uint32) int32{
	if exp != 0{
		ws.prints(proc,msg)
		proc.Terminate()
	}
	return 0
}

//C API :void exit(void)
func (ws *WasmService) exit(proc *exec.Process) int32{
	proc.Terminate()
	return 0
}