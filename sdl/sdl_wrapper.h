#if defined(_WIN32)
	#include <SDL2/SDL.h>
#if !(SDL_VERSION_ATLEAST(2,0,6))
	#include <SDL2/SDL_vulkan.h>
#endif
	#include <stdlib.h>
#else
	#include <SDL.h>
#if (SDL_VERSION_ATLEAST(2,0,6))
	#include <SDL_vulkan.h>
#endif
#endif

#if !defined(SDL_WINDOW_ALLOW_HIGHDPI)
	#define SDL_WINDOW_ALLOW_HIGHDPI (0x00002000)
#endif
