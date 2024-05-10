// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/media/core"
)

const SignatureMediaPlaybackItem string = "rc(Windows.Media.Playback.MediaPlaybackItem;{047097d2-e4af-48ab-b283-6929e674ece2})"

type MediaPlaybackItem struct {
	ole.IUnknown
}

func (impl *MediaPlaybackItem) AddAudioTracksChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.AddAudioTracksChanged(handler)
}

func (impl *MediaPlaybackItem) RemoveAudioTracksChanged(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.RemoveAudioTracksChanged(token)
}

func (impl *MediaPlaybackItem) AddVideoTracksChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.AddVideoTracksChanged(handler)
}

func (impl *MediaPlaybackItem) RemoveVideoTracksChanged(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.RemoveVideoTracksChanged(token)
}

func (impl *MediaPlaybackItem) AddTimedMetadataTracksChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.AddTimedMetadataTracksChanged(handler)
}

func (impl *MediaPlaybackItem) RemoveTimedMetadataTracksChanged(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.RemoveTimedMetadataTracksChanged(token)
}

func (impl *MediaPlaybackItem) GetSource() (*core.MediaSource, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.GetSource()
}

func (impl *MediaPlaybackItem) GetAudioTracks() (*MediaPlaybackAudioTrackList, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.GetAudioTracks()
}

func (impl *MediaPlaybackItem) GetVideoTracks() (*MediaPlaybackVideoTrackList, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.GetVideoTracks()
}

func (impl *MediaPlaybackItem) GetTimedMetadataTracks() (*MediaPlaybackTimedMetadataTrackList, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem))
	defer itf.Release()
	v := (*iMediaPlaybackItem)(unsafe.Pointer(itf))
	return v.GetTimedMetadataTracks()
}

func (impl *MediaPlaybackItem) GetBreakSchedule() (*MediaBreakSchedule, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem2))
	defer itf.Release()
	v := (*iMediaPlaybackItem2)(unsafe.Pointer(itf))
	return v.GetBreakSchedule()
}

func (impl *MediaPlaybackItem) GetStartTime() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem2))
	defer itf.Release()
	v := (*iMediaPlaybackItem2)(unsafe.Pointer(itf))
	return v.GetStartTime()
}

func (impl *MediaPlaybackItem) GetDurationLimit() (*foundation.IReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem2))
	defer itf.Release()
	v := (*iMediaPlaybackItem2)(unsafe.Pointer(itf))
	return v.GetDurationLimit()
}

func (impl *MediaPlaybackItem) GetCanSkip() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem2))
	defer itf.Release()
	v := (*iMediaPlaybackItem2)(unsafe.Pointer(itf))
	return v.GetCanSkip()
}

func (impl *MediaPlaybackItem) SetCanSkip(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem2))
	defer itf.Release()
	v := (*iMediaPlaybackItem2)(unsafe.Pointer(itf))
	return v.SetCanSkip(value)
}

func (impl *MediaPlaybackItem) GetDisplayProperties() (*MediaItemDisplayProperties, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem2))
	defer itf.Release()
	v := (*iMediaPlaybackItem2)(unsafe.Pointer(itf))
	return v.GetDisplayProperties()
}

func (impl *MediaPlaybackItem) ApplyDisplayProperties(value *MediaItemDisplayProperties) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem2))
	defer itf.Release()
	v := (*iMediaPlaybackItem2)(unsafe.Pointer(itf))
	return v.ApplyDisplayProperties(value)
}

func (impl *MediaPlaybackItem) GetIsDisabledInPlaybackList() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem3))
	defer itf.Release()
	v := (*iMediaPlaybackItem3)(unsafe.Pointer(itf))
	return v.GetIsDisabledInPlaybackList()
}

func (impl *MediaPlaybackItem) SetIsDisabledInPlaybackList(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem3))
	defer itf.Release()
	v := (*iMediaPlaybackItem3)(unsafe.Pointer(itf))
	return v.SetIsDisabledInPlaybackList(value)
}

func (impl *MediaPlaybackItem) GetTotalDownloadProgress() (float64, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem3))
	defer itf.Release()
	v := (*iMediaPlaybackItem3)(unsafe.Pointer(itf))
	return v.GetTotalDownloadProgress()
}

func (impl *MediaPlaybackItem) GetAutoLoadedDisplayProperties() (AutoLoadedDisplayPropertyKind, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem3))
	defer itf.Release()
	v := (*iMediaPlaybackItem3)(unsafe.Pointer(itf))
	return v.GetAutoLoadedDisplayProperties()
}

func (impl *MediaPlaybackItem) SetAutoLoadedDisplayProperties(value AutoLoadedDisplayPropertyKind) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackItem3))
	defer itf.Release()
	v := (*iMediaPlaybackItem3)(unsafe.Pointer(itf))
	return v.SetAutoLoadedDisplayProperties(value)
}

const GUIDiMediaPlaybackItem string = "047097d2-e4af-48ab-b283-6929e674ece2"
const SignatureiMediaPlaybackItem string = "{047097d2-e4af-48ab-b283-6929e674ece2}"

type iMediaPlaybackItem struct {
	ole.IInspectable
}

type iMediaPlaybackItemVtbl struct {
	ole.IInspectableVtbl

	AddAudioTracksChanged            uintptr
	RemoveAudioTracksChanged         uintptr
	AddVideoTracksChanged            uintptr
	RemoveVideoTracksChanged         uintptr
	AddTimedMetadataTracksChanged    uintptr
	RemoveTimedMetadataTracksChanged uintptr
	GetSource                        uintptr
	GetAudioTracks                   uintptr
	GetVideoTracks                   uintptr
	GetTimedMetadataTracks           uintptr
}

func (v *iMediaPlaybackItem) VTable() *iMediaPlaybackItemVtbl {
	return (*iMediaPlaybackItemVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaPlaybackItem) AddAudioTracksChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddAudioTracksChanged,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem) RemoveAudioTracksChanged(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveAudioTracksChanged,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaPlaybackItem) AddVideoTracksChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddVideoTracksChanged,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem) RemoveVideoTracksChanged(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveVideoTracksChanged,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaPlaybackItem) AddTimedMetadataTracksChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddTimedMetadataTracksChanged,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem) RemoveTimedMetadataTracksChanged(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveTimedMetadataTracksChanged,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaPlaybackItem) GetSource() (*core.MediaSource, error) {
	var out *core.MediaSource
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSource,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out core.MediaSource
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem) GetAudioTracks() (*MediaPlaybackAudioTrackList, error) {
	var out *MediaPlaybackAudioTrackList
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetAudioTracks,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaPlaybackAudioTrackList
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem) GetVideoTracks() (*MediaPlaybackVideoTrackList, error) {
	var out *MediaPlaybackVideoTrackList
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetVideoTracks,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaPlaybackVideoTrackList
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem) GetTimedMetadataTracks() (*MediaPlaybackTimedMetadataTrackList, error) {
	var out *MediaPlaybackTimedMetadataTrackList
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetTimedMetadataTracks,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaPlaybackTimedMetadataTrackList
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiMediaPlaybackItem2 string = "d859d171-d7ef-4b81-ac1f-f40493cbb091"
const SignatureiMediaPlaybackItem2 string = "{d859d171-d7ef-4b81-ac1f-f40493cbb091}"

type iMediaPlaybackItem2 struct {
	ole.IInspectable
}

type iMediaPlaybackItem2Vtbl struct {
	ole.IInspectableVtbl

	GetBreakSchedule       uintptr
	GetStartTime           uintptr
	GetDurationLimit       uintptr
	GetCanSkip             uintptr
	SetCanSkip             uintptr
	GetDisplayProperties   uintptr
	ApplyDisplayProperties uintptr
}

func (v *iMediaPlaybackItem2) VTable() *iMediaPlaybackItem2Vtbl {
	return (*iMediaPlaybackItem2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaPlaybackItem2) GetBreakSchedule() (*MediaBreakSchedule, error) {
	var out *MediaBreakSchedule
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetBreakSchedule,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaBreakSchedule
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem2) GetStartTime() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetStartTime,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem2) GetDurationLimit() (*foundation.IReference, error) {
	var out *foundation.IReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDurationLimit,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem2) GetCanSkip() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCanSkip,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem2) SetCanSkip(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetCanSkip,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaPlaybackItem2) GetDisplayProperties() (*MediaItemDisplayProperties, error) {
	var out *MediaItemDisplayProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDisplayProperties,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaItemDisplayProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem2) ApplyDisplayProperties(value *MediaItemDisplayProperties) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().ApplyDisplayProperties,
		uintptr(unsafe.Pointer(v)),     // this
		uintptr(unsafe.Pointer(value)), // in MediaItemDisplayProperties
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiMediaPlaybackItem3 string = "0d328220-b80a-4d09-9ff8-f87094a1c831"
const SignatureiMediaPlaybackItem3 string = "{0d328220-b80a-4d09-9ff8-f87094a1c831}"

type iMediaPlaybackItem3 struct {
	ole.IInspectable
}

type iMediaPlaybackItem3Vtbl struct {
	ole.IInspectableVtbl

	GetIsDisabledInPlaybackList    uintptr
	SetIsDisabledInPlaybackList    uintptr
	GetTotalDownloadProgress       uintptr
	GetAutoLoadedDisplayProperties uintptr
	SetAutoLoadedDisplayProperties uintptr
}

func (v *iMediaPlaybackItem3) VTable() *iMediaPlaybackItem3Vtbl {
	return (*iMediaPlaybackItem3Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaPlaybackItem3) GetIsDisabledInPlaybackList() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsDisabledInPlaybackList,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem3) SetIsDisabledInPlaybackList(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetIsDisabledInPlaybackList,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaPlaybackItem3) GetTotalDownloadProgress() (float64, error) {
	var out float64
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetTotalDownloadProgress,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out float64
	)

	if hr != 0 {
		return 0.0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem3) GetAutoLoadedDisplayProperties() (AutoLoadedDisplayPropertyKind, error) {
	var out AutoLoadedDisplayPropertyKind
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetAutoLoadedDisplayProperties,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out AutoLoadedDisplayPropertyKind
	)

	if hr != 0 {
		return AutoLoadedDisplayPropertyKindNone, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackItem3) SetAutoLoadedDisplayProperties(value AutoLoadedDisplayPropertyKind) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetAutoLoadedDisplayProperties,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in AutoLoadedDisplayPropertyKind
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiMediaPlaybackItemStatics string = "4b1be7f4-4345-403c-8a67-f5de91df4c86"
const SignatureiMediaPlaybackItemStatics string = "{4b1be7f4-4345-403c-8a67-f5de91df4c86}"

type iMediaPlaybackItemStatics struct {
	ole.IInspectable
}

type iMediaPlaybackItemStaticsVtbl struct {
	ole.IInspectableVtbl

	MediaPlaybackItemFindFromMediaSource uintptr
}

func (v *iMediaPlaybackItemStatics) VTable() *iMediaPlaybackItemStaticsVtbl {
	return (*iMediaPlaybackItemStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func MediaPlaybackItemFindFromMediaSource(source *core.MediaSource) (*MediaPlaybackItem, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Playback.MediaPlaybackItem", ole.NewGUID(GUIDiMediaPlaybackItemStatics))
	if err != nil {
		return nil, err
	}
	v := (*iMediaPlaybackItemStatics)(unsafe.Pointer(inspectable))

	var out *MediaPlaybackItem
	hr, _, _ := syscall.SyscallN(
		v.VTable().MediaPlaybackItemFindFromMediaSource,
		0,                               // this is a static func, so there's no this
		uintptr(unsafe.Pointer(source)), // in core.MediaSource
		uintptr(unsafe.Pointer(&out)),   // out MediaPlaybackItem
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiMediaPlaybackItemFactory2 string = "d77cdf3a-b947-4972-b35d-adfb931a71e6"
const SignatureiMediaPlaybackItemFactory2 string = "{d77cdf3a-b947-4972-b35d-adfb931a71e6}"

type iMediaPlaybackItemFactory2 struct {
	ole.IInspectable
}

type iMediaPlaybackItemFactory2Vtbl struct {
	ole.IInspectableVtbl

	MediaPlaybackItemCreateWithStartTime                 uintptr
	MediaPlaybackItemCreateWithStartTimeAndDurationLimit uintptr
}

func (v *iMediaPlaybackItemFactory2) VTable() *iMediaPlaybackItemFactory2Vtbl {
	return (*iMediaPlaybackItemFactory2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func MediaPlaybackItemCreateWithStartTime(source *core.MediaSource, startTime foundation.TimeSpan) (*MediaPlaybackItem, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Playback.MediaPlaybackItem", ole.NewGUID(GUIDiMediaPlaybackItemFactory2))
	if err != nil {
		return nil, err
	}
	v := (*iMediaPlaybackItemFactory2)(unsafe.Pointer(inspectable))

	var out *MediaPlaybackItem
	hr, _, _ := syscall.SyscallN(
		v.VTable().MediaPlaybackItemCreateWithStartTime,
		0,                               // this is a static func, so there's no this
		uintptr(unsafe.Pointer(source)), // in core.MediaSource
		uintptr(startTime.Duration),     // in foundation.TimeSpan
		uintptr(unsafe.Pointer(&out)),   // out MediaPlaybackItem
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func MediaPlaybackItemCreateWithStartTimeAndDurationLimit(source *core.MediaSource, startTime foundation.TimeSpan, durationLimit foundation.TimeSpan) (*MediaPlaybackItem, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Playback.MediaPlaybackItem", ole.NewGUID(GUIDiMediaPlaybackItemFactory2))
	if err != nil {
		return nil, err
	}
	v := (*iMediaPlaybackItemFactory2)(unsafe.Pointer(inspectable))

	var out *MediaPlaybackItem
	hr, _, _ := syscall.SyscallN(
		v.VTable().MediaPlaybackItemCreateWithStartTimeAndDurationLimit,
		0,                               // this is a static func, so there's no this
		uintptr(unsafe.Pointer(source)), // in core.MediaSource
		uintptr(startTime.Duration),     // in foundation.TimeSpan
		uintptr(durationLimit.Duration), // in foundation.TimeSpan
		uintptr(unsafe.Pointer(&out)),   // out MediaPlaybackItem
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiMediaPlaybackItemFactory string = "7133fce1-1769-4ff9-a7c1-38d2c4d42360"
const SignatureiMediaPlaybackItemFactory string = "{7133fce1-1769-4ff9-a7c1-38d2c4d42360}"

type iMediaPlaybackItemFactory struct {
	ole.IInspectable
}

type iMediaPlaybackItemFactoryVtbl struct {
	ole.IInspectableVtbl

	MediaPlaybackItemCreate uintptr
}

func (v *iMediaPlaybackItemFactory) VTable() *iMediaPlaybackItemFactoryVtbl {
	return (*iMediaPlaybackItemFactoryVtbl)(unsafe.Pointer(v.RawVTable))
}

func MediaPlaybackItemCreate(source *core.MediaSource) (*MediaPlaybackItem, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Playback.MediaPlaybackItem", ole.NewGUID(GUIDiMediaPlaybackItemFactory))
	if err != nil {
		return nil, err
	}
	v := (*iMediaPlaybackItemFactory)(unsafe.Pointer(inspectable))

	var out *MediaPlaybackItem
	hr, _, _ := syscall.SyscallN(
		v.VTable().MediaPlaybackItemCreate,
		0,                               // this is a static func, so there's no this
		uintptr(unsafe.Pointer(source)), // in core.MediaSource
		uintptr(unsafe.Pointer(&out)),   // out MediaPlaybackItem
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
