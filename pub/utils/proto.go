package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/golang/protobuf/proto"

	goProto "github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	ErrTypeNotMatch = errors.New("protoc type not match")
)

type FieldMeta struct {
	Index    int
	Name     string
	Offset   uintptr
	Type     reflect.Type
	ElemType reflect.Type
	Setter   func(ptr unsafe.Pointer, val string) error
	Getter   func(ptr unsafe.Pointer) interface{}
	IsNil    func(ptr unsafe.Pointer) bool
}

type ModelMeta struct {
	Name     string
	Type     reflect.Type
	ElemType reflect.Type
	Ptr      proto.Message
	Fields   map[string]*FieldMeta
	Reflect  protoreflect.Message // https://blog.golang.org/protobuf-apiv2
}

// nolint
func BuildModelMeta(name string) *ModelMeta {
	meta := &ModelMeta{
		Name:   name,
		Fields: map[string]*FieldMeta{},
	}

	meta.Type = MessageType(name)
	if meta.Type == nil {
		panic("proto: BuildModelMeta model is nil")
	}
	meta.ElemType = meta.Type.Elem()
	ptr, ok := reflect.New(meta.ElemType).Interface().(proto.Message)
	if !ok {
		panic("proto: BuildModelMeta model need PB Message")
	}
	meta.Ptr = ptr

	meta.Reflect = proto.MessageReflect(meta.Ptr)
	fd := meta.Reflect.Descriptor().Fields()
	for i := 0; i < fd.Len(); i++ {
		fName := string(fd.Get(i).Name())
		// Note: field index MUST be the same order as field no. (order by field no. asc)
		f := meta.ElemType.Field(i)
		fMeta := &FieldMeta{
			Index:    i,
			Name:     fName,
			Offset:   f.Offset,
			Type:     f.Type,
			ElemType: f.Type.Elem(),
		}

		// bytes
		if f.Type.Kind() == reflect.Slice && f.Type.Elem().Kind() == reflect.Uint8 {
			fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
				fPtr := (*[]byte)(ptr)
				if val == "" {
					*fPtr = nil
					return nil
				}
				*fPtr = []byte(val)
				return nil
			}
			fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
				fPtr := (*[]byte)(ptr)
				return *fPtr
			}
			fMeta.IsNil = func(ptr unsafe.Pointer) bool {
				fPtr := (*[]byte)(ptr)
				return *fPtr == nil
			}
		}

		// scalar types
		if f.Type.Kind() == reflect.Ptr {
			// scalar data types
			t := f.Type.Elem()
			switch t.Kind() {
			case reflect.Uint64:
				fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
					fPtr := (**uint64)(ptr)
					if val == "" {
						*fPtr = proto.Uint64(0)
						return nil
					}
					v1, err := strconv.ParseUint(val, 10, 64)
					if err != nil {
						return err
					}
					*fPtr = proto.Uint64(v1)
					return nil
				}
				fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
					fPtr := (**uint64)(ptr)
					if *fPtr != nil {
						return **fPtr
					}
					return uint64(0)
				}
				fMeta.IsNil = func(ptr unsafe.Pointer) bool {
					fPtr := (**uint64)(ptr)
					return *fPtr == nil
				}
			case reflect.Int64:
				fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
					fPtr := (**int64)(ptr)
					if val == "" {
						*fPtr = proto.Int64(0)
						return nil
					}
					v1, err := strconv.ParseInt(val, 10, 64)
					if err != nil {
						return err
					}
					*fPtr = proto.Int64(v1)
					return nil
				}
				fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
					fPtr := (**int64)(ptr)
					if *fPtr != nil {
						return **fPtr
					}
					return int64(0)
				}
				fMeta.IsNil = func(ptr unsafe.Pointer) bool {
					fPtr := (**int64)(ptr)
					return *fPtr == nil
				}
			case reflect.Uint32:
				fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
					fPtr := (**uint32)(ptr)
					if val == "" {
						*fPtr = proto.Uint32(0)
						return nil
					}
					v1, err := strconv.ParseUint(val, 10, 32)
					if err != nil {
						return err
					}
					*fPtr = proto.Uint32(uint32(v1))
					return nil
				}
				fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
					fPtr := (**uint32)(ptr)
					if *fPtr != nil {
						return **fPtr
					}
					return uint32(0)
				}
				fMeta.IsNil = func(ptr unsafe.Pointer) bool {
					fPtr := (**uint32)(ptr)
					return *fPtr == nil
				}
			case reflect.Int32:
				fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
					fPtr := (**int32)(ptr)
					if val == "" {
						*fPtr = proto.Int32(0)
						return nil
					}
					v1, err := strconv.ParseInt(val, 10, 32)
					if err != nil {
						return err
					}
					*fPtr = proto.Int32(int32(v1))
					return nil
				}
				fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
					fPtr := (**int32)(ptr)
					if *fPtr != nil {
						return **fPtr
					}
					return int32(0)
				}
				fMeta.IsNil = func(ptr unsafe.Pointer) bool {
					fPtr := (**int32)(ptr)
					return *fPtr == nil
				}
			case reflect.Bool:
				fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
					fPtr := (**bool)(ptr)
					if val == "" {
						*fPtr = proto.Bool(false)
						return nil
					}
					v1, err := strconv.ParseInt(val, 10, 64)
					if err != nil {
						return err
					}
					*fPtr = proto.Bool(v1 != 0)
					return nil
				}
				fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
					fPtr := (**bool)(ptr)
					if *fPtr != nil {
						return **fPtr
					}
					return false
				}
				fMeta.IsNil = func(ptr unsafe.Pointer) bool {
					fPtr := (**bool)(ptr)
					return *fPtr == nil
				}
			case reflect.Float64:
				fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
					fPtr := (**float64)(ptr)
					if val == "" {
						*fPtr = proto.Float64(0)
					}
					v1, err := strconv.ParseFloat(val, 64)
					if err != nil {
						return err
					}
					*fPtr = proto.Float64(v1)
					return nil
				}
				fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
					fPtr := (**float64)(ptr)
					if *fPtr != nil {
						return **fPtr
					}
					return float64(0)
				}
				fMeta.IsNil = func(ptr unsafe.Pointer) bool {
					fPtr := (**float64)(ptr)
					return *fPtr == nil
				}
			case reflect.Float32:
				fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
					fPtr := (**float32)(ptr)
					if val == "" {
						*fPtr = proto.Float32(0)
					}
					v1, err := strconv.ParseFloat(val, 32)
					if err != nil {
						return err
					}
					*fPtr = proto.Float32(float32(v1))
					return nil
				}
				fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
					fPtr := (**float32)(ptr)
					if *fPtr != nil {
						return **fPtr
					}
					return float32(0)
				}
				fMeta.IsNil = func(ptr unsafe.Pointer) bool {
					fPtr := (**float32)(ptr)
					return *fPtr == nil
				}
			case reflect.String:
				fMeta.Setter = func(ptr unsafe.Pointer, val string) error {
					fPtr := (**string)(ptr)
					if val == "" {
						*fPtr = proto.String("")
					}
					value, err := hex.DecodeString(val)
					if err != nil {
						return err
					}
					*fPtr = proto.String(string(value))
					return nil
				}
				fMeta.Getter = func(ptr unsafe.Pointer) interface{} {
					fPtr := (**string)(ptr)
					if *fPtr != nil {
						return **fPtr
					}
					return ""
				}
				fMeta.IsNil = func(ptr unsafe.Pointer) bool {
					fPtr := (**string)(ptr)
					return *fPtr == nil
				}
			}
		}

		meta.Fields[fName] = fMeta
	}
	return meta
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	//typ  *struct{}
	_    *struct{}
	word unsafe.Pointer
}

func TransformGDSRow(meta *ModelMeta, row map[string]string) (proto.Message, error) {
	msg, ok := reflect.New(meta.ElemType).Interface().(proto.Message)
	if !ok {
		return nil, ErrTypeNotMatch
	}
	msgPtr := (*emptyInterface)(unsafe.Pointer(&msg)).word
	for fName, fMeta := range meta.Fields {
		fVal, ok := row[strconv.Itoa(fMeta.Index)]
		if !ok {
			continue
		}
		fPtr := unsafe.Pointer(uintptr(msgPtr) + fMeta.Offset)
		err := fMeta.Setter(fPtr, fVal)
		if err != nil {
			return nil, fmt.Errorf("set %v error: %v", fName, err)
		}
	}
	return msg, nil
}

func MessageType(name string) reflect.Type {
	t := goProto.MessageType(name)
	if t == nil {
		// nolint: staticcheck
		t = proto.MessageType(name)
	}
	return t
}

func MessageName(x proto.Message) string {
	n := goProto.MessageName(x)
	if n == "" {
		// nolint: staticcheck
		n = proto.MessageName(x)
	}
	return n
}
