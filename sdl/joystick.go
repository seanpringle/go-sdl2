package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

const (
	HAT_CENTERED  = C.SDL_HAT_CENTERED
	HAT_UP        = C.SDL_HAT_UP
	HAT_RIGHT     = C.SDL_HAT_RIGHT
	HAT_DOWN      = C.SDL_HAT_DOWN
	HAT_LEFT      = C.SDL_HAT_LEFT
	HAT_RIGHTUP   = C.SDL_HAT_RIGHTUP
	HAT_RIGHTDOWN = C.SDL_HAT_RIGHTDOWN
	HAT_LEFTUP    = C.SDL_HAT_LEFTUP
	HAT_LEFTDOWN  = C.SDL_HAT_LEFTDOWN
)

const (
    JOYSTICK_TYPE_UNKNOWN JoystickType = 0
    JOYSTICK_TYPE_GAMECONTROLLER
    JOYSTICK_TYPE_WHEEL
    JOYSTICK_TYPE_ARCADE_STICK
    JOYSTICK_TYPE_FLIGHT_STICK
    JOYSTICK_TYPE_DANCE_PAD
    JOYSTICK_TYPE_GUITAR
    JOYSTICK_TYPE_DRUM_KIT
    JOYSTICK_TYPE_ARCADE_PAD
    JOYSTICK_TYPE_THROTTLE
)

type Joystick C.SDL_Joystick
type JoystickGUID C.SDL_JoystickGUID
type JoystickID C.SDL_JoystickID
type JoystickType C.SDL_JoystickType

func (joy *Joystick) cptr() *C.SDL_Joystick {
	return (*C.SDL_Joystick)(unsafe.Pointer(joy))
}

func (guid JoystickGUID) c() C.SDL_JoystickGUID {
	return C.SDL_JoystickGUID(guid)
}

// NumJoysticks (https://wiki.libsdl.org/SDL_NumJoysticks)
func NumJoysticks() int {
	return (int)(C.SDL_NumJoysticks())
}

// JoystickNameForIndex (https://wiki.libsdl.org/SDL_JoystickNameForIndex)
func JoystickNameForIndex(index int) string {
	return (C.GoString)(C.SDL_JoystickNameForIndex(C.int(index)))
}

// JoystickOpen (https://wiki.libsdl.org/SDL_JoystickOpen)
func JoystickOpen(index JoystickID) *Joystick {
	return (*Joystick)(C.SDL_JoystickOpen(C.int(index)))
}

// JoystickName (https://wiki.libsdl.org/SDL_JoystickName)
func (joy *Joystick) Name() string {
	return (C.GoString)(C.SDL_JoystickName(joy.cptr()))
}

// TODO
func (joy *Joystick) Vendor() uint16 {
	return 0
}

// TODO
func (joy *Joystick) Product() uint16 {
	return 0
}

// TODO
func (joy *Joystick) ProductVersion() uint16 {
	return 0
}

// TODO
func (joy *Joystick) Type() JoystickType {
	return 0
}

// JoystickGetDeviceGUID (https://wiki.libsdl.org/SDL_JoystickGetDeviceGUID)
func JoystickGetDeviceGUID(index int) JoystickGUID {
	return (JoystickGUID)(C.SDL_JoystickGetDeviceGUID(C.int(index)))
}

// TODO
func JoystickGetDeviceVendor(index int) uint16 {
	return (uint16)(C.SDL_JoystickGetDeviceVendor(C.int(index)))
}

// TODO
func JoystickGetDeviceProduct(index int) uint16 {
	return (uint16)(C.SDL_JoystickGetDeviceProduct(C.int(index)))
}

// TODO
func JoystickGetDeviceProductVersion(index int) uint16 {
	return (uint16)(C.SDL_JoystickGetDeviceProductVersion(C.int(index)))
}

// TODO
func JoystickGetDeviceType(index int) JoystickType {
	return (JoystickType)(C.SDL_JoystickGetDeviceType(C.int(index)))
}

// Joystick (https://wiki.libsdl.org/SDL_JoystickGetGUID)
func (joy *Joystick) GetGUID() JoystickGUID {
	return (JoystickGUID)(C.SDL_JoystickGetGUID(joy.cptr()))
}

// JoystickGetGUIDString (https://wiki.libsdl.org/SDL_JoystickGetGUIDString)
func JoystickGetGUIDString(guid JoystickGUID) string {
	_pszGUID := make([]rune, 1024)
	pszGUID := C.CString(string(_pszGUID[:]))
	defer C.free(unsafe.Pointer(pszGUID))
	C.SDL_JoystickGetGUIDString(guid.c(), pszGUID, C.int(unsafe.Sizeof(_pszGUID)))
	return C.GoString(pszGUID)
}

// JoystickGetGUIDFromString (https://wiki.libsdl.org/SDL_JoystickGetGUIDFromString)
func JoystickGetGUIDFromString(pchGUID string) JoystickGUID {
	_pchGUID := C.CString(pchGUID)
	defer C.free(unsafe.Pointer(_pchGUID))
	return (JoystickGUID)(C.SDL_JoystickGetGUIDFromString(_pchGUID))
}

// JoystickGetAttached (https://wiki.libsdl.org/SDL_JoystickGetAttached)
func (joy *Joystick) GetAttached() bool {
	return C.SDL_JoystickGetAttached(joy.cptr()) > 0
}

// JoystickInstanceID (https://wiki.libsdl.org/SDL_JoystickInstanceID)
func (joy *Joystick) InstanceID() JoystickID {
	return (JoystickID)(C.SDL_JoystickInstanceID(joy.cptr()))
}

// JoystickNumAxes (https://wiki.libsdl.org/SDL_JoystickNumAxes)
func (joy *Joystick) NumAxes() int {
	return (int)(C.SDL_JoystickNumAxes(joy.cptr()))
}

// JoystickNumBalls (https://wiki.libsdl.org/SDL_JoystickNumBalls)
func (joy *Joystick) NumBalls() int {
	return (int)(C.SDL_JoystickNumBalls(joy.cptr()))
}

// JoystickNumHats (https://wiki.libsdl.org/SDL_JoystickNumHats)
func (joy *Joystick) NumHats() int {
	return (int)(C.SDL_JoystickNumHats(joy.cptr()))
}

// JoystickNumButtons (https://wiki.libsdl.org/SDL_JoystickNumButtons)
func (joy *Joystick) NumButtons() int {
	return (int)(C.SDL_JoystickNumButtons(joy.cptr()))
}

// JoystickUpdate (https://wiki.libsdl.org/SDL_JoystickUpdate)
func JoystickUpdate() {
	C.SDL_JoystickUpdate()
}

// JoystickEventState (https://wiki.libsdl.org/SDL_JoystickEventState)
func JoystickEventState(state int) int {
	return (int)(C.SDL_JoystickEventState(C.int(state)))
}

// JoystickGetAxis (https://wiki.libsdl.org/SDL_JoystickGetAxis)
func (joy *Joystick) GetAxis(axis int) int16 {
	return (int16)(C.SDL_JoystickGetAxis(joy.cptr(), C.int(axis)))
}

// TODO
func (joy *Joystick) GetAxisInitialState(axis int) (initialState int16, hasInitialValue bool) {
	hasInitialValue = C.SDL_JoystickGetAxisInitialState(joy.cptr(), C.int(axis), (*C.Sint16)(&initialState)) == C.SDL_TRUE
	return
}

// JoystickGetHat (https://wiki.libsdl.org/SDL_JoystickGetHat)
func (joy *Joystick) GetHat(hat int) byte {
	return (byte)(C.SDL_JoystickGetHat(joy.cptr(), C.int(hat)))
}

// JoystickGetBall (https://wiki.libsdl.org/SDL_JoystickGetBall)
func (joy *Joystick) GetBall(ball int, dx, dy *int) int {
	_dx := (*C.int)(unsafe.Pointer(dx))
	_dy := (*C.int)(unsafe.Pointer(dy))
	return (int)(C.SDL_JoystickGetBall(joy.cptr(), C.int(ball), _dx, _dy))
}

// JoystickGetButton (https://wiki.libsdl.org/SDL_JoystickGetButton)
func (joy *Joystick) GetButton(button int) byte {
	return (byte)(C.SDL_JoystickGetButton(joy.cptr(), C.int(button)))
}

// JoystickClose (https://wiki.libsdl.org/SDL_JoystickClose)
func (joy *Joystick) Close() {
	C.SDL_JoystickClose(joy.cptr())
}
