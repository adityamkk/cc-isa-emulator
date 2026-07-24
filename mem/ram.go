package mem

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync"
)

const LINE_SIZE = 64

type RandomAccessMemory struct {
	memmap  sync.Map
	binFile io.ReaderAt
	linemap sync.Map
}

func NewRandomAccessMemory(binFile io.ReaderAt) *RandomAccessMemory {
	return &RandomAccessMemory{binFile: binFile}
}

func (ram *RandomAccessMemory) Read1Byte(address uint16) uint8 {
	val, ok := ram.memmap.Load(address)
	if ok {
		return val.(uint8)
	}
	buf := make([]byte, 1)
	if _, err := ram.binFile.ReadAt(buf, int64(address)); err != nil {
		return 0
	}
	return buf[0]
}

func (ram *RandomAccessMemory) Read1Word(address uint16) uint16 {
	valFirst, okFirst := ram.memmap.Load(address)
	valSecond, okSecond := ram.memmap.Load(address + 1)
	buf := make([]byte, 2)

	if okFirst && okSecond {
		buf[0] = valFirst.(uint8)
		buf[1] = valSecond.(uint8)
		return binary.LittleEndian.Uint16(buf)
	}

	if !okFirst && !okSecond {
		if _, err := ram.binFile.ReadAt(buf, int64(address)); err != nil {
			return 0
		}
	} else if okFirst {
		buf[0] = valFirst.(uint8)
		secondByte := make([]byte, 1)
		if _, err := ram.binFile.ReadAt(secondByte, int64(address)); err != nil {
			return 0
		}
		buf[1] = secondByte[0]
	} else {
		buf[1] = valSecond.(uint8)
		firstByte := make([]byte, 1)
		if _, err := ram.binFile.ReadAt(firstByte, int64(address)); err != nil {
			return 0
		}
		buf[0] = firstByte[0]
	}
	return binary.LittleEndian.Uint16(buf)
}

func (ram *RandomAccessMemory) Write1Byte(address uint16, data uint8) {
	ram.memmap.Store(address, data)
}

func (ram *RandomAccessMemory) Write1Word(address uint16, data uint16) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, data)
	ram.memmap.Store(address, buf[0])
	ram.memmap.Store(address+1, buf[1])
}

// AcquireLine acquires a lock on the line that the address belongs to
func (ram *RandomAccessMemory) AcquireLine(address uint16) {
	lockIdx := address / LINE_SIZE
	_lock, _ := ram.linemap.LoadOrStore(lockIdx, &sync.Mutex{})
	lock := _lock.(*sync.Mutex)
	lock.Lock()
}

// ReleaseLine releases a lock on a line that the address belongs to
func (ram *RandomAccessMemory) ReleaseLine(address uint16) {
	lockIdx := address / LINE_SIZE
	_lock, ok := ram.linemap.Load(lockIdx)
	if !ok {
		fmt.Println("Error: Attempted to release an unacquired memory lock")
		os.Exit(1)
	}
	lock := _lock.(*sync.Mutex)
	lock.Unlock()
}
