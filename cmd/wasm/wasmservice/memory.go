package wasmservice

import(
	"github.com/go-interpreter/wagon/exec"
)

// C API:int32 memcpy(void *dest, const void *src, uint32 n)
func (ws *WasmService) memcpy(proc *exec.Process, dest, src, len int32) int32{
	if dest < src && dest + len > src{
		return -1
	}
	msg := make([]byte, len)
	err := proc.ReadAt(msg, int(src), int(len))
	if err != nil {
		return -1
	}
	err = proc.WriteAt(msg[:], int(dest), int(len))
	if err != nil{
		return -1
	}
	return 0
}

// C API:int32 memset(void *s, char ch, uint32 n)
func (ws *WasmService) memset(proc *exec.Process, addr, ch, len int32) int32{
	tmp := make([]byte, len)
	for i:= 0; i < int(len); i++{
		tmp[i] = byte(ch)
	}
	err := proc.WriteAt(tmp[:], int(addr), int(len))
	if err != nil{
		return -1
	}
	return 0
}














