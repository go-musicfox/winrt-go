// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package core

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const GUIDISingleSelectMediaTrackList string = "77206f1f-c34f-494f-8077-2bad9ff4ecf1"
const SignatureISingleSelectMediaTrackList string = "{77206f1f-c34f-494f-8077-2bad9ff4ecf1}"

type ISingleSelectMediaTrackList struct {
	ole.IInspectable
}

type ISingleSelectMediaTrackListVtbl struct {
	ole.IInspectableVtbl

	AddSelectedIndexChanged    uintptr
	RemoveSelectedIndexChanged uintptr
	SetSelectedIndex           uintptr
	GetSelectedIndex           uintptr
}

func (v *ISingleSelectMediaTrackList) VTable() *ISingleSelectMediaTrackListVtbl {
	return (*ISingleSelectMediaTrackListVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ISingleSelectMediaTrackList) AddSelectedIndexChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddSelectedIndexChanged,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *ISingleSelectMediaTrackList) RemoveSelectedIndexChanged(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveSelectedIndexChanged,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISingleSelectMediaTrackList) SetSelectedIndex(value int32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetSelectedIndex,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in int32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISingleSelectMediaTrackList) GetSelectedIndex() (int32, error) {
	var out int32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSelectedIndex,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out int32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}
