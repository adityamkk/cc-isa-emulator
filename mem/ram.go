package mem

import (
	"encoding/binary"
	"io"
	"sync"
)

type RandomAccessMemory struct {
	Memmap  sync.Map
	binFile io.ReaderAt
}

func NewRandomAccessMemory(binFile io.ReaderAt) *RandomAccessMemory {
	return &RandomAccessMemory{binFile: binFile}
}

func (ram *RandomAccessMemory) readByte(address uint16) uint8 {
	val, ok := ram.Memmap.Load(address)
	if ok {
		return val.(uint8)
	}
	buf := make([]byte, 1)
	if _, err := ram.binFile.ReadAt(buf, int64(address)); err != nil {
		return 0
	}
	return buf[0]
}

func (ram *RandomAccessMemory) readWord(address uint16) uint16 {
	_valFirst, okFirst := ram.Memmap.Load(address)
	_valSecond, okSecond := ram.Memmap.Load(address + 1)
	valFirst := _valFirst.(uint8)
	valSecond := _valSecond.(uint8)
	buf := make([]byte, 2)

	if okFirst && okSecond {
		buf[0] = valFirst
		buf[1] = valSecond
		return binary.LittleEndian.Uint16(buf)
	}

	if !okFirst && !okSecond {
		if _, err := ram.binFile.ReadAt(buf, int64(address)); err != nil {
			return 0
		}
	} else if okFirst {
		buf[0] = valFirst
		secondByte := make([]byte, 1)
		if _, err := ram.binFile.ReadAt(secondByte, int64(address)); err != nil {
			return 0
		}
		buf[1] = secondByte[0]
	} else {
		buf[1] = valSecond
		firstByte := make([]byte, 1)
		if _, err := ram.binFile.ReadAt(firstByte, int64(address)); err != nil {
			return 0
		}
		buf[0] = firstByte[0]
	}
	return binary.LittleEndian.Uint16(buf)
}

func (ram *RandomAccessMemory) writeByte(address uint16, data uint8) {
	ram.Memmap.Store(address, data)
}

func (ram *RandomAccessMemory) writeWord(address uint16, data uint16) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, data)
	ram.Memmap.Store(address, buf[0])
	ram.Memmap.Store(address+1, buf[1])
}
