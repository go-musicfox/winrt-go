// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package fileproperties

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/foundation/collections"
)

const SignatureVideoProperties string = "rc(Windows.Storage.FileProperties.VideoProperties;{719ae507-68de-4db8-97de-49998c059f2f})"

type VideoProperties struct {
	ole.IUnknown
}

func (impl *VideoProperties) GetRating() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetRating()
}

func (impl *VideoProperties) SetRating(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.SetRating(value)
}

func (impl *VideoProperties) GetKeywords() (*collections.IVector, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetKeywords()
}

func (impl *VideoProperties) GetWidth() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetWidth()
}

func (impl *VideoProperties) GetHeight() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetHeight()
}

func (impl *VideoProperties) GetDuration() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetDuration()
}

func (impl *VideoProperties) GetLatitude() (*foundation.IReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetLatitude()
}

func (impl *VideoProperties) GetLongitude() (*foundation.IReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetLongitude()
}

func (impl *VideoProperties) GetTitle() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetTitle()
}

func (impl *VideoProperties) SetTitle(value string) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.SetTitle(value)
}

func (impl *VideoProperties) GetSubtitle() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetSubtitle()
}

func (impl *VideoProperties) SetSubtitle(value string) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.SetSubtitle(value)
}

func (impl *VideoProperties) GetProducers() (*collections.IVector, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetProducers()
}

func (impl *VideoProperties) GetPublisher() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetPublisher()
}

func (impl *VideoProperties) SetPublisher(value string) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.SetPublisher(value)
}

func (impl *VideoProperties) GetWriters() (*collections.IVector, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetWriters()
}

func (impl *VideoProperties) GetYear() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetYear()
}

func (impl *VideoProperties) SetYear(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.SetYear(value)
}

func (impl *VideoProperties) GetBitrate() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetBitrate()
}

func (impl *VideoProperties) GetDirectors() (*collections.IVector, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
	defer itf.Release()
	v := (*iVideoProperties)(unsafe.Pointer(itf))
	return v.GetDirectors()
}

//func (impl *VideoProperties) GetOrientation() (VideoOrientation, error) {
//	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiVideoProperties))
//	defer itf.Release()
//	v := (*iVideoProperties)(unsafe.Pointer(itf))
//	return v.GetOrientation()
//}
//
//func (impl *VideoProperties) RetrievePropertiesAsync(propertiesToRetrieve *collections.IIterable) (*foundation.IAsyncOperation, error) {
//	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIStorageItemExtraProperties))
//	defer itf.Release()
//	v := (*IStorageItemExtraProperties)(unsafe.Pointer(itf))
//	return v.RetrievePropertiesAsync(propertiesToRetrieve)
//}
//
//func (impl *VideoProperties) SavePropertiesAsync(propertiesToSave *collections.IIterable) (*foundation.IAsyncAction, error) {
//	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIStorageItemExtraProperties))
//	defer itf.Release()
//	v := (*IStorageItemExtraProperties)(unsafe.Pointer(itf))
//	return v.SavePropertiesAsync(propertiesToSave)
//}
//
//func (impl *VideoProperties) SavePropertiesAsyncOverloadDefault() (*foundation.IAsyncAction, error) {
//	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIStorageItemExtraProperties))
//	defer itf.Release()
//	v := (*IStorageItemExtraProperties)(unsafe.Pointer(itf))
//	return v.SavePropertiesAsyncOverloadDefault()
//}

const GUIDiVideoProperties string = "719ae507-68de-4db8-97de-49998c059f2f"
const SignatureiVideoProperties string = "{719ae507-68de-4db8-97de-49998c059f2f}"

type iVideoProperties struct {
	ole.IInspectable
}

type iVideoPropertiesVtbl struct {
	ole.IInspectableVtbl

	GetRating      uintptr
	SetRating      uintptr
	GetKeywords    uintptr
	GetWidth       uintptr
	GetHeight      uintptr
	GetDuration    uintptr
	GetLatitude    uintptr
	GetLongitude   uintptr
	GetTitle       uintptr
	SetTitle       uintptr
	GetSubtitle    uintptr
	SetSubtitle    uintptr
	GetProducers   uintptr
	GetPublisher   uintptr
	SetPublisher   uintptr
	GetWriters     uintptr
	GetYear        uintptr
	SetYear        uintptr
	GetBitrate     uintptr
	GetDirectors   uintptr
	GetOrientation uintptr
}

func (v *iVideoProperties) VTable() *iVideoPropertiesVtbl {
	return (*iVideoPropertiesVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iVideoProperties) GetRating() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetRating,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) SetRating(value uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetRating,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iVideoProperties) GetKeywords() (*collections.IVector, error) {
	var out *collections.IVector
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetKeywords,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVector
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetWidth() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetWidth,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetHeight() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetHeight,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetDuration() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDuration,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetLatitude() (*foundation.IReference, error) {
	var out *foundation.IReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetLatitude,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetLongitude() (*foundation.IReference, error) {
	var out *foundation.IReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetLongitude,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetTitle() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetTitle,
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

func (v *iVideoProperties) SetTitle(value string) error {
	valueHStr, err := ole.NewHString(value)
	if err != nil {
		return err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetTitle,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(valueHStr),         // in string
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iVideoProperties) GetSubtitle() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSubtitle,
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

func (v *iVideoProperties) SetSubtitle(value string) error {
	valueHStr, err := ole.NewHString(value)
	if err != nil {
		return err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetSubtitle,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(valueHStr),         // in string
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iVideoProperties) GetProducers() (*collections.IVector, error) {
	var out *collections.IVector
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetProducers,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVector
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetPublisher() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetPublisher,
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

func (v *iVideoProperties) SetPublisher(value string) error {
	valueHStr, err := ole.NewHString(value)
	if err != nil {
		return err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetPublisher,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(valueHStr),         // in string
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iVideoProperties) GetWriters() (*collections.IVector, error) {
	var out *collections.IVector
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetWriters,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVector
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetYear() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetYear,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) SetYear(value uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetYear,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iVideoProperties) GetBitrate() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetBitrate,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iVideoProperties) GetDirectors() (*collections.IVector, error) {
	var out *collections.IVector
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDirectors,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVector
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

//func (v *iVideoProperties) GetOrientation() (VideoOrientation, error) {
//	var out VideoOrientation
//	hr, _, _ := syscall.SyscallN(
//		v.VTable().GetOrientation,
//		uintptr(unsafe.Pointer(v)),    // this
//		uintptr(unsafe.Pointer(&out)), // out VideoOrientation
//	)
//
//	if hr != 0 {
//		return VideoOrientationNormal, ole.NewError(hr)
//	}
//
//	return out, nil
//}
