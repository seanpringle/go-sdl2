package sdl

/*
#include "sdl_wrapper.h"

#if !SDL_VERSION_ATLEAST(2,0,6)
#pragma message("SDL_VulkanLoadLibrary is not supported before SDL 2.0.6")

static inline int SDL_Vulkan_LoadLibrary(const char *path)
{
	return -2;
}

static inline void SDL_Vulkan_GetDrawableSize(SDL_Window *window, int *w, int *h)
{
	if (w)
		w = 0;
	if (h)
		h = 0;
}
#endif
*/
import "C"
import (
	"errors"
)

type VkInstance C.VkInstance
type VkSurfaceKHR C.VkSurfaceKHR

// TODO
func Vulkan_LoadLibrary(path string) (err error) {
	ret := C.SDL_Vulkan_LoadLibrary(C.CString(path))
	if ret == -1 {
		err = GetError()
	} else if ret == -2 {
		err = errors.New("SDL_VulkanLoadLibrary is not supported before SDL 2.0.6")
	}

	return
}

// TODO
func Vulkan_GetVkGetInstanceProcAddr() uintptr {
	return 0
}

// TODO
func Vulkan_UnloadLibrary() {
	C.SDL_Vulkan_UnloadLibrary()
}

// TODO
func Vulkan_GetInstanceExtensions(window *Window) (count uint, names []string, err error) {
	return
}

// TODO
func Vulkan_CreateSurface(window *Window, instance VkInstance, surface *VkSurfaceKHR) (err error) {
	return
}

// TODO
func Vulkan_GetDrawableSize(window *Window) (w, h int) {
	var _w, _h C.int

	C.SDL_Vulkan_GetDrawableSize(window.cptr(), &_w, &_h)
	w = int(_w)
	h = int(_h)

	return
}
