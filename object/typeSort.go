package object

import (
	"fmt"
	"reflect"
)

type sortThype interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type SliceCount struct {
	slice     []interface{}
	limitType reflect.Type
}

func (sliceCount *SliceCount) Less(i, j int) bool {
	//switch sliceCount.limitType.Kind() {
	//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//	return reflect.ValueOf(sliceCount.slice[i]).Int() < reflect.ValueOf(sliceCount.slice[j]).Int()
	//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	//	return reflect.ValueOf(sliceCount.slice[i]).Uint() < reflect.ValueOf(sliceCount.slice[j]).Uint()
	//}
	return sliceCount.compare(sliceCount.slice[i], sliceCount.slice[j]) == 1
	return true
}

func (sliceCount *SliceCount) compare(e interface{}, f interface{}) int {
	//if reflect.TypeOf(e) != sliceCount.limitType {
	//	return 0
	//}
	switch sliceCount.limitType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v1 := reflect.ValueOf(e).Int()
		v2 := reflect.ValueOf(f).Int()
		if v1 == v2 {
			return 0
		} else if v1 < v2 {
			return 1
		} else {
			return 2
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v1 := reflect.ValueOf(e).Uint()
		v2 := reflect.ValueOf(f).Uint()
		if v1 == v2 {
			return 0
		} else if v1 < v2 {
			return 1
		} else {
			return 2
		}
	}
	return 1
}
func NewSliceCount(e interface{}) *SliceCount {
	_type := reflect.TypeOf(e)
	sliceCount := &SliceCount{}
	switch e.(type) {
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		sliceCount.limitType = _type
	default:
		return nil
	}
	return sliceCount
}

func (sliceCount *SliceCount) Len() int {
	return len(sliceCount.slice)
}

func (sliceCount *SliceCount) Swap(i, j int) {
	sliceCount.slice[i], sliceCount.slice[j] = sliceCount.slice[j], sliceCount.slice[i]
}

func (sliceCount *SliceCount) Add(e interface{}) bool {
	if reflect.TypeOf(e) != sliceCount.limitType {
		return false
	}
	sliceCount.slice = append(sliceCount.slice, e)
	return true
}

func (sliceCount *SliceCount) find(e interface{}) int {
	if reflect.TypeOf(e) != sliceCount.limitType {
		return 0
	}
	end := sliceCount.Len()
	//eValue := reflect.ValueOf(e)
	begin := 0
	index := end / 2
	for {
		if sliceCount.compare(e, sliceCount.slice[index]) == 1 {
			end = index
		} else if sliceCount.compare(e, sliceCount.slice[index]) == 2 {
			begin = index
		} else {
			return index + 1
		}
		if begin >= end {
			break
		}
		index = (begin + end) / 2
	}
	return 0
}

func (sliceCount *SliceCount) Remove(e interface{}) bool {
	pos := sliceCount.find(e)
	if pos == 0 {
		return false
	}
	if pos == 1 {
		sliceCount.slice = sliceCount.slice[1:len(sliceCount.slice)]
	} else if pos == len(sliceCount.slice)-1 {
		sliceCount.slice = sliceCount.slice[:len(sliceCount.slice)-1]
	} else {
		sliceCount.slice = append(sliceCount.slice[:pos-1], sliceCount.slice[(pos):sliceCount.Len()])
	}
	return true
}

func (sliceCount *SliceCount) String() {
	fmt.Println(sliceCount.slice)
}
