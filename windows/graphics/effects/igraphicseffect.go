// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package effects

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const GUIDIGraphicsEffect string = "cb51c0ce-8fe6-4636-b202-861faa07d8f3"
const SignatureIGraphicsEffect string = "{cb51c0ce-8fe6-4636-b202-861faa07d8f3}"

type IGraphicsEffect struct {
	ole.IInspectable
}

type IGraphicsEffectVtbl struct {
	ole.IInspectableVtbl

	GetName uintptr
	SetName uintptr
}

func (v *IGraphicsEffect) VTable() *IGraphicsEffectVtbl {
	return (*IGraphicsEffectVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IGraphicsEffect) GetName() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetName,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

func (v *IGraphicsEffect) SetName(name string) error {
	nameHStr, err := ole.NewHString(name)
	if err != nil {
		return err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetName,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(nameHStr),          // in string
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
