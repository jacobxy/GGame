package stream

import (
	"errors"
	"math"
)

type Stream struct {
	pos  int
	data []byte
}

const (
	STREAM_POOL = 1000
)

var (
	_pool = make(chan *Stream, STREAM_POOL)
)

func init() {
	go func() {
		for {
			_pool <- &Stream{data: make([]byte, 0, 512)}
		}
	}()
}

func GetStream() *Stream {
	return <-_pool
}

func Reader(data []byte) *Stream {
	return &Stream{data: data}
}

func (st *Stream) Data() []byte {
	return st.data
}

func (st *Stream) Length() int {
	return len(st.data)
}

func (st *Stream) ReadBool() (ret bool, err error) {
	b, _err := st.ReadByte()
	if b != byte(1) {
		return false, _err
	}
	return true, _err
}

func (st *Stream) ReadByte() (ret byte, err error) {
	if st.pos >= len(st.data) {
		err = errors.New("read byte failed")
		return
	}

	ret = st.data[st.pos]
	st.pos++
	return
}

func (st *Stream) ReadString() (ret string, err error) {
	if st.pos+2 > len(st.data) {
		err = errors.New("read string error")
		return
	}

	size, _ := st.Read16()
	if st.pos+int(size) > len(st.data) {
		err = errors.New("read string data error")
		return
	}

	bytes := st.data[st.pos : st.pos+int(size)]
	st.pos += int(size)
	ret = string(bytes)
	return
}

func (st *Stream) ReadU16() (ret uint16, err error) {
	if st.pos+2 >= len(st.data) {
		err = errors.New("read u16 failed")
		return
	}

	buf := st.data[st.pos : st.pos+2]
	ret = uint16(buf[0])<<8 | uint16(buf[1])
	st.pos += 2
	return
}

func (st *Stream) Read16() (ret int16, err error) {
	_ret, _err := st.ReadU16()
	ret = int16(_ret)
	err = _err
	return
}

func (st *Stream) ReadU24() (ret uint32, err error) {
	if st.pos+4 >= len(st.data) {
		err = errors.New("read u16 failed")
		return
	}

	buf := st.data[st.pos : st.pos+4]
	ret = uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])
	st.pos += 2
	return
}

func (st *Stream) Read24() (ret int32, err error) {
	_ret, _err := st.ReadU24()
	ret = int32(_ret)
	err = _err
	return
}

func (st *Stream) ReadU32() (ret uint32, err error) {
	if st.pos+4 >= len(st.data) {
		err = errors.New("read u16 failed")
		return
	}

	buf := st.data[st.pos : st.pos+4]
	ret = uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])
	st.pos += 2
	return
}

func (st *Stream) Read32() (ret int32, err error) {
	_ret, _err := st.ReadU32()
	ret = int32(_ret)
	err = _err
	return
}

func (st *Stream) ReadU64() (ret uint64, err error) {
	if st.pos+8 >= len(st.data) {
		err = errors.New("read u64 failed")
		return
	}

	buf := st.data[st.pos : st.pos+8]
	ret = uint64(buf[0])<<24 | uint64(buf[1])<<16 | uint64(buf[2])<<8 | uint64(buf[3])
	st.pos += 2
	return
}

func (st *Stream) Read64() (ret int64, err error) {
	_ret, _err := st.ReadU64()
	ret = int64(_ret)
	err = _err
	return
}

func (st *Stream) ReadFloat32() (ret float32, err error) {
	bits, _err := st.ReadU32()
	if _err != nil {
		return float32(0), _err
	}

	ret = math.Float32frombits(bits)

	if math.IsNaN(float64(ret)) || math.IsInf(float64(ret), 0) {
		return 0, nil
	}
	return ret, nil
}

func (st *Stream) ReadFloat64() (ret float64, err error) {
	bits, _err := st.ReadU64()
	if _err != nil {
		return float64(0), _err
	}

	ret = math.Float64frombits(bits)

	if math.IsNaN(float64(ret)) || math.IsInf(float64(ret), 0) {
		return 0, nil
	}

	return ret, nil
}

func (st *Stream) ReadBytes() (ret []byte, err error) {
	size, _err := st.ReadU16()
	if _err != nil {
		err = _err
		return
	}

	if st.pos+int(size) > len(st.data) {
		err = errors.New("read bytes failed")
		return
	}

	ret = st.data[st.pos : st.pos+int(size)]
	st.pos += int(size)
	return
}

//=====================================writes

func (st *Stream) WriteZeros(cnt int) {
	for i := 0; i < cnt; i++ {
		st.data = append(st.data, byte(0))
	}
}

func (st *Stream) WriteBool(v bool) {
	if v {
		st.data = append(st.data, byte(1))
	} else {
		st.data = append(st.data, byte(0))
	}
}

func (st *Stream) WriteByte(v byte) {
	st.data = append(st.data, v)
}

func (st *Stream) WriteBytes(v []byte) {
	st.WriteU16(uint16(len(v)))
	st.data = append(st.data, v...)
}

func (st *Stream) WriteRawBytes(v []byte) {
	st.data = append(st.data, v...)
}

func (st *Stream) WriteString(str string) {
	bytes := []byte(str)
	st.WriteBytes(bytes)
}

func (st *Stream) WriteU16(v uint16) {
	st.data = append(st.data, byte(v>>8), byte(v))
}

func (st *Stream) Write16(v int16) {
	st.WriteU16(uint16(v))
}

func (st *Stream) WriteU24(v uint32) {
	st.data = append(st.data, byte(v>>16), byte(v>>8), byte(v))
}

func (st *Stream) Write24(v int32) {
	st.WriteU24(uint32(v))
}

func (st *Stream) WriteU32(v uint32) {
	st.data = append(st.data, byte(v>>16), byte(v>>8), byte(v))
}

func (st *Stream) Write32(v int32) {
	st.WriteU24(uint32(v))
}

func (st *Stream) WriteU64(v uint64) {
	st.data = append(st.data, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

func (st *Stream) Write64(v int64) {
	st.WriteU64(uint64(v))
}

func (st *Stream) WriteFloat32(v float32) {
	f := math.Float32bits(v)
	st.WriteU32(f)
}

func (st *Stream) WriteFloat64(v float64) {
	f := math.Float64bits(v)
	st.WriteU64(f)
}
