package exec

import (
	"encoding/binary"
	"reflect"
	"github.com/go-interpreter/wagon/util"
	"errors"
)
type DataType uint32

const (
	DInt8 DataType = iota
	DInt16
	DInt32
	DInt64
	DFloat32
	DFloat64
	DString
	DStruct
	DUnkown
)

const (
	NIL_ADDR = 0
)

type Block struct {
	ptype   DataType
	size    int
}

type Memmanage struct {
	memory          []byte
	blocks          map[int]*Block
	allocatedBytes  int
}

func (mm *Memmanage) Malloc(size int, ptype DataType) (int, error) {
	if mm.memory == nil || len(mm.memory) == 0{
		return 0, errors.New("memory is not initialized")
	}
	if mm.allocatedBytes+size > len(mm.memory) || size < 0{
		return 0, errors.New("parameter error")
	}

	offset := mm.allocatedBytes + 1
	mm.allocatedBytes += size
	mm.blocks[offset] = &Block{ptype: ptype, size:size}
	return offset, nil
}

func (mm *Memmanage) GetBlockSize(addr int) (int, error) {

	if addr == NIL_ADDR{
		return 0, errors.New("addr wrong")
	}
	v, ok := mm.blocks[addr]
	if ok {
		return v.size,nil
	} else {
		return 0, errors.New("addr wrong")
	}
}

func (mm *Memmanage) GetBlockData(addr int) ([]byte, error) {

	if addr == NIL_ADDR {
		return nil, errors.New("addr wrong")
	}

	length, err := mm.GetBlockSize(addr)
	if err != nil {
		return nil, err
	}

	if addr + length > len(mm.memory) {
		return nil, errors.New("memory out of bound")
	} else {
		return mm.memory[addr:addr+length], nil
	}
}

func (mm *Memmanage) SetBlockData(addr int, val []byte) (int, error) {

	if addr == NIL_ADDR {
		return 0,errors.New("addr wrong")
	}

	length, err := mm.GetBlockSize(addr)
	if err != nil {
		return 0,err
	}
	if length > len(val) {
		length = len(val)
	}

	copy(mm.memory[addr:addr+length], val)
	return length,nil
}
func (mm *Memmanage) Load2Mem(b []byte, ptype DataType) (int, error) {
	index, err := mm.Malloc(len(b), ptype)
	if err != nil {
		return 0, err
	}
	copy(mm.memory[index:index+len(b)], b)

	return index, nil
}

func (mm *Memmanage) SetBlock(val interface{}) (int, error) {

	if val == nil {
		return NIL_ADDR, nil
	}

	switch reflect.TypeOf(val).Kind() {
	case reflect.String:
		b := []byte(val.(string))
		b = append(b,0)
		return mm.Load2Mem(b, DString)
	case reflect.Slice:
		switch val.(type) {
		case []byte:
			return mm.Load2Mem(val.([]byte), DString)

		case []int32:
			intBytes := make([]byte, len(val.([]int32))*4)
			for i, v := range val.([]int32) {
				tmp := make([]byte, 4)
				binary.LittleEndian.PutUint32(tmp, uint32(v))
				copy(intBytes[i*4:(i+1)*4], tmp)
			}
			return mm.Load2Mem(intBytes, DInt32)
		case []int64:
			intBytes := make([]byte, len(val.([]int64))*8)
			for i, v := range val.([]int64) {
				tmp := make([]byte, 8)
				binary.LittleEndian.PutUint64(tmp, uint64(v))
				copy(intBytes[i*8:(i+1)*8], tmp)
			}
			return mm.Load2Mem(intBytes, DInt64)

		case []float32:
			floatBytes := make([]byte, len(val.([]float32))*4)
			for i, v := range val.([]float32) {
				tmp := util.Float32ToBytes(v)
				copy(floatBytes[i*4:(i+1)*4], tmp)
			}
			return mm.Load2Mem(floatBytes, DFloat32)

		case []float64:
			floatBytes := make([]byte, len(val.([]float64))*8)
			for i, v := range val.([]float64) {
				tmp := util.Float64ToBytes(v)
				copy(floatBytes[i*8:(i+1)*8], tmp)
			}
			return mm.Load2Mem(floatBytes, DFloat64)

		case []string:
			addrs := make([]byte, len(val.([]string))*4)
			for i, s := range val.([]string) {
				addr, err := mm.SetBlock(s)
				if err != nil {
					return 0, err
				}
				tmp := make([]byte, 4)
				binary.LittleEndian.PutUint32(tmp, uint32(addr))
				copy(addrs[i*4:(i+1)*4], tmp)
			}
			return mm.Load2Mem(addrs, DInt32)

		default:
			return 0, errors.New("unsupported slice type")
		}
	default:
		return 0, errors.New("unsupported type")
	}
	return 0,nil
}