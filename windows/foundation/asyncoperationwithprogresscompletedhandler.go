// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package foundation

import (
	"sync"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/internal/delegate"
	"github.com/saltosystems/winrt-go/internal/kernel32"
)

const GUIDAsyncOperationWithProgressCompletedHandler string = "e85df41d-6aa7-46e3-a8e2-f009d840c627"
const SignatureAsyncOperationWithProgressCompletedHandler string = "delegate({e85df41d-6aa7-46e3-a8e2-f009d840c627})"

type AsyncOperationWithProgressCompletedHandler struct {
	ole.IUnknown
	sync.Mutex
	refs uint64
	IID  ole.GUID
}

type AsyncOperationWithProgressCompletedHandlerVtbl struct {
	ole.IUnknownVtbl
	Invoke uintptr
}

type AsyncOperationWithProgressCompletedHandlerCallback func(instance *AsyncOperationWithProgressCompletedHandler, asyncInfo *IAsyncOperationWithProgress, asyncStatus AsyncStatus)

var callbacksAsyncOperationWithProgressCompletedHandler = &asyncOperationWithProgressCompletedHandlerCallbacks{
	mu:        &sync.Mutex{},
	callbacks: make(map[unsafe.Pointer]AsyncOperationWithProgressCompletedHandlerCallback),
}

var releaseChannelsAsyncOperationWithProgressCompletedHandler = &asyncOperationWithProgressCompletedHandlerReleaseChannels{
	mu:    &sync.Mutex{},
	chans: make(map[unsafe.Pointer]chan struct{}),
}

func NewAsyncOperationWithProgressCompletedHandler(iid *ole.GUID, callback AsyncOperationWithProgressCompletedHandlerCallback) *AsyncOperationWithProgressCompletedHandler {
	size := unsafe.Sizeof(*(*AsyncOperationWithProgressCompletedHandler)(nil))
	instPtr := kernel32.Malloc(size)
	inst := (*AsyncOperationWithProgressCompletedHandler)(instPtr)

	callbacks := delegate.RegisterCallbacks(instPtr, inst)

	// Initialize all properties: the malloc may contain garbage
	inst.RawVTable = (*interface{})(unsafe.Pointer(&AsyncOperationWithProgressCompletedHandlerVtbl{
		IUnknownVtbl: ole.IUnknownVtbl{
			QueryInterface: callbacks.QueryInterface,
			AddRef:         callbacks.AddRef,
			Release:        callbacks.Release,
		},
		Invoke: callbacks.Invoke,
	}))
	inst.IID = *iid // copy contents
	inst.Mutex = sync.Mutex{}
	inst.refs = 0

	callbacksAsyncOperationWithProgressCompletedHandler.add(unsafe.Pointer(inst), callback)

	// See the docs in the releaseChannelsAsyncOperationWithProgressCompletedHandler struct
	releaseChannelsAsyncOperationWithProgressCompletedHandler.acquire(unsafe.Pointer(inst))

	inst.addRef()
	return inst
}

func (r *AsyncOperationWithProgressCompletedHandler) GetIID() *ole.GUID {
	return &r.IID
}

// addRef increments the reference counter by one
func (r *AsyncOperationWithProgressCompletedHandler) addRef() uint64 {
	r.Lock()
	defer r.Unlock()
	r.refs++
	return r.refs
}

// removeRef decrements the reference counter by one. If it was already zero, it will just return zero.
func (r *AsyncOperationWithProgressCompletedHandler) removeRef() uint64 {
	r.Lock()
	defer r.Unlock()

	if r.refs > 0 {
		r.refs--
	}

	return r.refs
}

func (instance *AsyncOperationWithProgressCompletedHandler) Invoke(instancePtr, rawArgs0, rawArgs1, rawArgs2, rawArgs3, rawArgs4, rawArgs5, rawArgs6, rawArgs7, rawArgs8 unsafe.Pointer) uintptr {
	asyncInfoPtr := rawArgs0
	asyncStatusRaw := (int32)(uintptr(rawArgs1))

	// See the quote above.
	asyncInfo := (*IAsyncOperationWithProgress)(asyncInfoPtr)
	asyncStatus := (AsyncStatus)(asyncStatusRaw)
	if callback, ok := callbacksAsyncOperationWithProgressCompletedHandler.get(instancePtr); ok {
		callback(instance, asyncInfo, asyncStatus)
	}
	return ole.S_OK
}

func (instance *AsyncOperationWithProgressCompletedHandler) AddRef() uint64 {
	return instance.addRef()
}

func (instance *AsyncOperationWithProgressCompletedHandler) Release() uint64 {
	rem := instance.removeRef()
	if rem == 0 {
		// We're done.
		instancePtr := unsafe.Pointer(instance)
		callbacksAsyncOperationWithProgressCompletedHandler.delete(instancePtr)

		// stop release channels used to avoid
		// https://github.com/golang/go/issues/55015
		releaseChannelsAsyncOperationWithProgressCompletedHandler.release(instancePtr)

		kernel32.Free(instancePtr)
	}
	return rem
}

type asyncOperationWithProgressCompletedHandlerCallbacks struct {
	mu        *sync.Mutex
	callbacks map[unsafe.Pointer]AsyncOperationWithProgressCompletedHandlerCallback
}

func (m *asyncOperationWithProgressCompletedHandlerCallbacks) add(p unsafe.Pointer, v AsyncOperationWithProgressCompletedHandlerCallback) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.callbacks[p] = v
}

func (m *asyncOperationWithProgressCompletedHandlerCallbacks) get(p unsafe.Pointer) (AsyncOperationWithProgressCompletedHandlerCallback, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.callbacks[p]
	return v, ok
}

func (m *asyncOperationWithProgressCompletedHandlerCallbacks) delete(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.callbacks, p)
}

// typedEventHandlerReleaseChannels keeps a map with channels
// used to keep a goroutine alive during the lifecycle of this object.
// This is required to avoid causing a deadlock error.
// See this: https://github.com/golang/go/issues/55015
type asyncOperationWithProgressCompletedHandlerReleaseChannels struct {
	mu    *sync.Mutex
	chans map[unsafe.Pointer]chan struct{}
}

func (m *asyncOperationWithProgressCompletedHandlerReleaseChannels) acquire(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	c := make(chan struct{})
	m.chans[p] = c

	go func() {
		// we need a timer to trick the go runtime into
		// thinking there's still something going on here
		// but we are only really interested in <-c
		t := time.NewTimer(time.Minute)
		for {
			select {
			case <-t.C:
				t.Reset(time.Minute)
			case <-c:
				t.Stop()
				return
			}
		}
	}()
}

func (m *asyncOperationWithProgressCompletedHandlerReleaseChannels) release(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if c, ok := m.chans[p]; ok {
		close(c)
		delete(m.chans, p)
	}
}
