package path_tracing

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/camera"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/light"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/object"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/screen"
)

// PathTracer is a class for path tracing algorithm.
//
// Members:
// 	objects     - The list of objects.
//  pixelScreen - The screen.
//  sceneCamera - The camera on the scene.
//  lights      - The list of light objects.
//
type PathTracer struct {
	objects     []*object.Object
	pixelScreen *screen.Screen
	sceneCamera *camera.Camera
	lights      []*light.Light
}

// GetObjects gets the objects of the PathTracer.
//
// Parameters:
// 	none
//
// Returns:
// 	The objects of the PathTracer.
//
func (pathTracer *PathTracer) GetObjects() []*object.Object {
	return pathTracer.objects
}

// GetPixelScreen gets the pixel screen of the PathTracer.
//
// Parameters:
// 	none
//
// Returns:
// 	The pixel screen of the PathTracer.
//
func (pathTracer *PathTracer) GetPixelScreen() *screen.Screen {
	return pathTracer.pixelScreen
}

// GetSceneCamera gets the camera on the scene of the PathTracer.
//
// Parameters:
// 	none
//
// Returns:
// 	The camera on the scene of the PathTracer.
//
func (pathTracer *PathTracer) GetSceneCamera() *camera.Camera {
	return pathTracer.sceneCamera
}

// GetLights gets the list of light objects of the PathTracer.
//
// Parameters:
// 	none
//
// Returns:
// 	The list of light objects of the PathTracer.
//
func (pathTracer *PathTracer) GetLights() []*light.Light {
	return pathTracer.lights
}

// Init initializes a PathTracer.
//
// Parameters:
// 	objects     - The list of objects.
//  pixelScreen - The screen.
//  sceneCamera - The camera on the scene.
//  lights      - The list of light objects.
//
// Returns:
// 	a PathTracer.
//
func Init(objects []*object.Object, pixelScreen *screen.Screen, sceneCamera *camera.Camera,
	lights []*light.Light) *PathTracer {
	return &PathTracer{objects: objects, pixelScreen: pixelScreen, sceneCamera: sceneCamera, lights: lights}
}
