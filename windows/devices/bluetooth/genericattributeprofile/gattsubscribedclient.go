// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package genericattributeprofile

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureGattSubscribedClient string = "rc(Windows.Devices.Bluetooth.GenericAttributeProfile.GattSubscribedClient;{736e9001-15a4-4ec2-9248-e3f20d463be9})"

type GattSubscribedClient struct {
	ole.IUnknown
}

func (impl *GattSubscribedClient) GetSession() (*GattSession, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSubscribedClient))
	defer itf.Release()
	v := (*iGattSubscribedClient)(unsafe.Pointer(itf))
	return v.GetSession()
}

func (impl *GattSubscribedClient) GetMaxNotificationSize() (uint16, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSubscribedClient))
	defer itf.Release()
	v := (*iGattSubscribedClient)(unsafe.Pointer(itf))
	return v.GetMaxNotificationSize()
}

func (impl *GattSubscribedClient) AddMaxNotificationSizeChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSubscribedClient))
	defer itf.Release()
	v := (*iGattSubscribedClient)(unsafe.Pointer(itf))
	return v.AddMaxNotificationSizeChanged(handler)
}

func (impl *GattSubscribedClient) RemoveMaxNotificationSizeChanged(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSubscribedClient))
	defer itf.Release()
	v := (*iGattSubscribedClient)(unsafe.Pointer(itf))
	return v.RemoveMaxNotificationSizeChanged(token)
}

const GUIDiGattSubscribedClient string = "736e9001-15a4-4ec2-9248-e3f20d463be9"
const SignatureiGattSubscribedClient string = "{736e9001-15a4-4ec2-9248-e3f20d463be9}"

type iGattSubscribedClient struct {
	ole.IInspectable
}

type iGattSubscribedClientVtbl struct {
	ole.IInspectableVtbl

	GetSession                       uintptr
	GetMaxNotificationSize           uintptr
	AddMaxNotificationSizeChanged    uintptr
	RemoveMaxNotificationSizeChanged uintptr
}

func (v *iGattSubscribedClient) VTable() *iGattSubscribedClientVtbl {
	return (*iGattSubscribedClientVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iGattSubscribedClient) GetSession() (*GattSession, error) {
	var out *GattSession
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSession,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out GattSession
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattSubscribedClient) GetMaxNotificationSize() (uint16, error) {
	var out uint16
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMaxNotificationSize,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint16
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattSubscribedClient) AddMaxNotificationSizeChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddMaxNotificationSizeChanged,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattSubscribedClient) RemoveMaxNotificationSizeChanged(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveMaxNotificationSizeChanged,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
