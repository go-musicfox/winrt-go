// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package streams

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/storage"
)

const SignatureRandomAccessStreamReference string = "rc(Windows.Storage.Streams.RandomAccessStreamReference;{33ee3134-1dd6-4e3a-8067-d1c162e8642b})"

type RandomAccessStreamReference struct {
	ole.IUnknown
}

func (impl *RandomAccessStreamReference) OpenReadAsync() (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIRandomAccessStreamReference))
	defer itf.Release()
	v := (*IRandomAccessStreamReference)(unsafe.Pointer(itf))
	return v.OpenReadAsync()
}

const GUIDiRandomAccessStreamReferenceStatics string = "857309dc-3fbf-4e7d-986f-ef3b1a07a964"
const SignatureiRandomAccessStreamReferenceStatics string = "{857309dc-3fbf-4e7d-986f-ef3b1a07a964}"

type iRandomAccessStreamReferenceStatics struct {
	ole.IInspectable
}

type iRandomAccessStreamReferenceStaticsVtbl struct {
	ole.IInspectableVtbl

	RandomAccessStreamReferenceCreateFromFile   uintptr
	RandomAccessStreamReferenceCreateFromUri    uintptr
	RandomAccessStreamReferenceCreateFromStream uintptr
}

func (v *iRandomAccessStreamReferenceStatics) VTable() *iRandomAccessStreamReferenceStaticsVtbl {
	return (*iRandomAccessStreamReferenceStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func RandomAccessStreamReferenceCreateFromFile(file *storage.IStorageFile) (*RandomAccessStreamReference, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Streams.RandomAccessStreamReference", ole.NewGUID(GUIDiRandomAccessStreamReferenceStatics))
	if err != nil {
		return nil, err
	}
	v := (*iRandomAccessStreamReferenceStatics)(unsafe.Pointer(inspectable))

	var out *RandomAccessStreamReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().RandomAccessStreamReferenceCreateFromFile,
		0,                             // this is a static func, so there's no this
		uintptr(unsafe.Pointer(file)), // in storage.IStorageFile
		uintptr(unsafe.Pointer(&out)), // out RandomAccessStreamReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func RandomAccessStreamReferenceCreateFromUri(uri *foundation.Uri) (*RandomAccessStreamReference, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Streams.RandomAccessStreamReference", ole.NewGUID(GUIDiRandomAccessStreamReferenceStatics))
	if err != nil {
		return nil, err
	}
	v := (*iRandomAccessStreamReferenceStatics)(unsafe.Pointer(inspectable))

	var out *RandomAccessStreamReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().RandomAccessStreamReferenceCreateFromUri,
		0,                             // this is a static func, so there's no this
		uintptr(unsafe.Pointer(uri)),  // in foundation.Uri
		uintptr(unsafe.Pointer(&out)), // out RandomAccessStreamReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func RandomAccessStreamReferenceCreateFromStream(stream *IRandomAccessStream) (*RandomAccessStreamReference, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Streams.RandomAccessStreamReference", ole.NewGUID(GUIDiRandomAccessStreamReferenceStatics))
	if err != nil {
		return nil, err
	}
	v := (*iRandomAccessStreamReferenceStatics)(unsafe.Pointer(inspectable))

	var out *RandomAccessStreamReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().RandomAccessStreamReferenceCreateFromStream,
		0,                               // this is a static func, so there's no this
		uintptr(unsafe.Pointer(stream)), // in IRandomAccessStream
		uintptr(unsafe.Pointer(&out)),   // out RandomAccessStreamReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}