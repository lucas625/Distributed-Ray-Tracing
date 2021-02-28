package path_tracing

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/line"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/camera"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/color_matrix"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/ray"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/screen"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/thread_locker"
	"math"
	"math/rand"
	"time"
)

// Controller is a class for controlling the path tracing algorithm.
//
// Members:
// 	none
//
type Controller struct {}

// RandomInSemiSphere is a function to find a ray for diffuse reflection in semisphere.
//
// Parameters:
//  normal - the normal.
//
// Returns:
// 	the vector.
//
func RandomInSemiSphere(normal utils.Vector, pos entity.Point) utils.Vector {

	found := false

	var random1,random2,random3 float64
	var v utils.Vector
	for !found {
		found = true
		random1 = rand.Float64()
		random2 = rand.Float64()
		random3 = rand.Float64()

		v = utils.Vector{Coordinates: []float64{random1, random2, random3}}
		v = utils.CMultVector(&v, 2)
		vaux := utils.Vector{Coordinates: []float64{1,1,1}}
		v = utils.SumVector(&v, &vaux, 1, -1)
		if math.Pow(utils.VectorNorm(&v),2) >= 1{
			found = false
		}
	}
	vect := utils.SumVector(&normal, &v, 1, 1)
	vect = utils.NormalizeVector(&vect)
	return vect
}

// RandomInSemiSphereSpecular is a function to find a ray for diffuse reflection in semisphere.
//
// Parameters:
//  none
//
// Returns:
// 	the vector.
//
func RandomInSemiSphereSpecular() utils.Vector {

	found := false

	var random1,random2,random3 float64
	var v utils.Vector
	for !found {
		found = true
		random1 = rand.Float64()
		random2 = rand.Float64()
		random3 = rand.Float64()

		v = utils.Vector{Coordinates: []float64{random1, random2, random3}}
		v = utils.CMultVector(&v, 2)
		vaux := utils.Vector{Coordinates: []float64{1,1,1}}
		v = utils.SumVector(&v, &vaux, 1, -1)
		if math.Pow(utils.VectorNorm(&v),2) >= 1{
			found = false
		}
	}
	return v
}

// FindNextRay is a function to find the next line.
//
// Parameters:
//  pos       - the point.
//  obj       - the object.
//  triangIdx - the index of the triangle.
//  bCoords   - baricentric coords for each respective normal.
//
// Returns:
// 	the line
//
func (controller *Controller) FindNextRay(ptracer *PathTracer, pos entity.Point, obj general.Object, triangleIdx int, bCoords []float64) entity.Line {
	normals := make([]utils.Vector, 3)
	for i := 0; i < 3; i++ {
		normals[i] = obj.Normals[obj.Triangles[triangleIdx].Normals[i]]
	}
	resultingNormal := utils.SumVector(&normals[0], &normals[1], bCoords[0], bCoords[1])
	resultingNormal = utils.SumVector(&resultingNormal, &normals[2], 1, bCoords[2])
	resultingNormal = utils.NormalizeVector(&resultingNormal)

	ktot := obj.DiffuseReflection + obj.SpecularReflection // + obj.TransReflection
	r := 0.0 + rand.Float64()*ktot
	vector := utils.Vector{Coordinates: []float64{1.0, 1.0, 1.0}}
	if r <= obj.DiffuseReflection {
		vector = RandomInSemiSphere(resultingNormal, pos)
	} else if r <= obj.DiffuseReflection+obj.SpecularReflection {
		lightPos := ptracer.Lgts.LightList[0].LightObject.GetCenter()
		Lvector := entity.ExtractVector(&pos, &lightPos)
		Lvector = utils.NormalizeVector(&Lvector)

		constantPart := 2 * utils.DotProduct(&resultingNormal, &Lvector)

		vector = utils.SumVector(&resultingNormal, &Lvector, constantPart, -1) // R = 2N(N.L) - L

		offsetVector := RandomInSemiSphereSpecular()
		offsetVector = utils.CMultVector(&offsetVector, obj.RoughNess)

		vector = utils.SumVector(&vector, &offsetVector, 1, 1)

	} else {
		// use transmission (unavailable)
	}
	vector = utils.NormalizeVector(&vector)
	line := entity.Line{Start: pos, Director: vector}
	return line
}

// TraceRayDepth is a function to trace a ray and return the resulting color.
//
// Parameters:
//  line       - the ray.
//  recursions - number of recursions.
//
// Returns:
// 	the rgb color at a given position.
//
func (ptracer *PathTracer) TraceRayDepth(line entity.Line, recursions int) []float64 {
	color := make([]float64, 3)

	closestT := math.MaxFloat64
	closestObjIdx := -1
	closestTriangleIndex := -1
	closestBCoords := make([]float64, 3)
	for objIdx, obj := range ptracer.Objs.ObjList { // iterating through all objects
		for triangIdx, triangle := range obj.Triangles { // iterating through all triangles of an object
			points := make([]entity.Point, 3)
			for pi := 0; pi < 3; pi++ { // getting triangle points
				points[pi] = obj.Vertices.Points[triangle.Vertices[pi]]
			}
			t, bCoords, intersected := line.IntersectTriangle(points)
			if intersected {
				if t > 0 && t < closestT {
					closestT = t
					closestObjIdx = objIdx
					closestTriangleIndex = triangIdx
					closestBCoords = bCoords
				}
			}
		}
	}

	// intersecting lights
	lightClosest := false
	for lgtIdx, lgt := range ptracer.Lgts.LightList {
		for _, triangle := range lgt.LightObject.Triangles { // iterating through all triangles of an object
			points := make([]entity.Point, 3)
			for pi := 0; pi < 3; pi++ { // getting triangle points
				points[pi] = lgt.LightObject.Vertices.Points[triangle.Vertices[pi]]
			}
			t, _, intersected := line.IntersectTriangle(points)
			if intersected {
				if t > 0 && t <= closestT {
					lightClosest = true
					closestT = t
					closestObjIdx = lgtIdx
				}
			}
		}
	}

	if !lightClosest {
		if closestObjIdx != -1 {
			colorAux := []float64{1, 1, 1}
			if recursions > 0 {
				newLine := ptracer.FindNextRay(line.FindPos(closestT), ptracer.Objs.ObjList[closestObjIdx], closestTriangleIndex, closestBCoords)
				colorAux = ptracer.TraceRayDepth(newLine, recursions-1)
			}
			for i := 0; i < 3; i++ {
				color[i] = ptracer.Objs.ObjList[closestObjIdx].Color[i] * colorAux[i]
			}
		}
	} else {
		for i := 0; i < 3; i++ {
			lgtAux := ptracer.Lgts.LightList[closestObjIdx]
			color[i] = lgtAux.Color[i] * lgtAux.LightIntensity
		}
	}
	return color
}

// intersectObjects uses a ray to intersect all objects.
//
// Parameters:
// 	pathTracer - The PathTracer.
//  currentRay - The current ray.
//
// Returns:
// 	If there is intersections with the objects.
// 	The closest line parameter.
// 	The closest object index.
// 	The closest triangle index.
// 	The closest triangle is barycentric coordinates.
//
func (controller *Controller) intersectObjects(pathTracer *PathTracer, currentRay *line.Line) (bool, float64, int, int,
	[]float64) {
	closestLineParameter := math.MaxFloat64
	closesObjectIndex := -1
	closestTriangleIndex := -1
	closestTriangleBarycentricCoordinates := make([]float64, 3)
	hasObjectIntersections := false
	rayController := ray.Controller{}

	for objectIndex, currentObject := range pathTracer.GetObjects() {
		for triangleIndex, currentTriangle := range currentObject.GetTriangles() {

			lineParameter, barycentricCoordinates, hasIntersection, _ := rayController.IntersectRayTriangle(
				currentRay, currentTriangle, currentObject.GetRepository())

			if hasIntersection {
				hasObjectIntersections = true
				if lineParameter >= 1 && lineParameter < closestLineParameter {
					closestLineParameter = lineParameter
					closesObjectIndex = objectIndex
					closestTriangleIndex = triangleIndex
					closestTriangleBarycentricCoordinates = barycentricCoordinates
				}
			}
		}
	}
	return hasObjectIntersections, closestLineParameter, closesObjectIndex, closestTriangleIndex,
	closestTriangleBarycentricCoordinates
}

// intersectLights uses a ray to intersect all lights.
//
// Parameters:
// 	pathTracer           - The PathTracer.
//  currentRay           - The current ray.
//  closestLineParameter - The closest line parameter found through objects intersections.
//
// Returns:
// 	If there is intersections with lights.
// 	The closest light is line parameter.
// 	The closest light index.
//
func (controller *Controller) intersectLights(pathTracer *PathTracer, currentRay *line.Line) (bool, float64, int) {
	closesLightLineParameterIndex := math.MaxFloat64
	closestLightIndex := -1
	hasLightIntersection := false
	rayController := ray.Controller{}

	for lightIndex, currentLight := range pathTracer.GetLights() {
		for _, currentTriangle := range currentLight.GetLightObject().GetTriangles() {
			lineParameter, _, hasIntersection, _ := rayController.IntersectRayTriangle(
				currentRay, currentTriangle, currentLight.GetLightObject().GetRepository())

			if hasIntersection {
				hasLightIntersection = true
				if lineParameter >= 1 && lineParameter <= closesLightLineParameterIndex {
					closesLightLineParameterIndex = lineParameter
					closestLightIndex = lightIndex
				}
			}
		}
	}

	return hasLightIntersection, closesLightLineParameterIndex, closestLightIndex
}

// iterateRay uses a ray to calculate the color.
//
// Parameters:
// 	pathTracer       - The PathTracer.
// 	currentIteration - The number of the current iteration.
//  depthIterations  - Number of depth rays recursions.
//  currentRay       - The current ray.
//
// Returns:
// 	The color found by the ray.
//
func (controller *Controller) iterateRay(pathTracer *PathTracer, currentIteration, depthIterations int,
	currentRay *line.Line) []float64 {
	color := make([]float64, 3)

	hasObjectIntersection, closestLineParameter, closesObjectIndex, closestTriangleIndex, closestTriangleBarycentricCoordinates :=
		controller.intersectObjects(pathTracer, currentRay)

	hasLightIntersection, closestLightLineParameter, closestLight := controller.intersectLights(
		pathTracer, currentRay)

	if hasLightIntersection && closestLightLineParameter <= closestLineParameter {
		intersectedLight := pathTracer.GetLights()[closestLight]
		for index := 0; index < 3; index++ {
			color[index] = intersectedLight.GetColor()[index] * intersectedLight.GetLightIntensity()
		}

	} else {
		if hasObjectIntersection {
			colorAux := []float64{0, 0, 0}
			if depthIterations > 0 {
				lineController := line.Controller{}
				newRayStartingPoint, _ := lineController.FindPoint(currentRay, closestLineParameter)
				newLine := controller.FindNextRay(pathTracer, newRayStartingPoint,
					pathTracer.GetObjects()[closesObjectIndex], closestTriangleIndex,
					closestTriangleBarycentricCoordinates)
				colorAux = pathTracer.TraceRayDepth(newLine, depthIterations-1)
			}
			objectColor := pathTracer.GetObjects()[closesObjectIndex].GetLightCharacteristics().GetColor()
			for index := 0; index < 3; index++ {
				color[index] = objectColor[index] * colorAux[index]
			}
		}
	}
	return color
}

// parseRaysColorsToRGB traces all primary rays of a pixel.
//
// Parameters:
// 	raysColors   - The colors found by the rays.
// 	numberOfRays - The number of rays.
//
// Returns:
// 	The average of the colors as RGB.
//
func (controller *Controller) parseRaysColorsToRGB(raysColors [][]float64, numberOfRays int) []int {
	color := make([]float64, 3)
	for rayIndex := 0; rayIndex < numberOfRays; rayIndex++ {
		for rayColorCoordinateIndex := 0; rayColorCoordinateIndex < 3; rayColorCoordinateIndex++ {
			color[rayColorCoordinateIndex] = color[rayColorCoordinateIndex] + raysColors[rayIndex][
			rayColorCoordinateIndex]
		}
	}

	rgbColor := make([]int, 3)
	for index := 0; index < 3; index++ {
		rgbColor[index] = int(math.Floor(math.Sqrt(color[index]/float64(numberOfRays)) * 255))
		if rgbColor[index] > 255 {
			rgbColor[index] = 255
		} else if rgbColor[index] < 0 {
			rgbColor[index] = 0
		}
	}
	return rgbColor
}

// traceFirstRays traces all primary rays of a pixel.
//
// Parameters:
// 	pathTracer      - The PathTracer.
// 	lineIndex       - Pixel line index.
//  columnIndex     - Pixel column index.
//  numberOfRays    - Number of rays per pixel.
//  depthIterations - Number of depth rays recursions.
//
// Returns:
// 	The RGB color of the pixel.
//
func (controller *Controller) traceFirstRays(pathTracer *PathTracer, lineIndex, columnIndex, numberOfRays,
	depthIterations int) []int {
	rand.Seed(time.Now().UnixNano())
	floatColors := make([][]float64, numberOfRays)
	lock := thread_locker.Init()
	cameraController := &camera.Controller{}
	cameraToWorldMatrix := cameraController.CameraToWorldMatrix(pathTracer.GetSceneCamera())
	for rayIndex := 0; rayIndex < numberOfRays; rayIndex++ {
		if lock.GetThreads() < 12 {
			lock.AddThread()
			go func(threadRayIndex int) {
				pixelLineOffset := rand.Float64()
				pixelColumnOffset := rand.Float64()

				screenController := &screen.Controller{}
				rayVectorDirector, _ := screenController.BuildRayVectorDirectorToPixel(lineIndex, columnIndex,
					pixelLineOffset, pixelColumnOffset, cameraToWorldMatrix, pathTracer.GetPixelScreen(),
					pathTracer.GetSceneCamera())

				currentRay, _ := line.Init(pathTracer.GetSceneCamera().GetPosition(), rayVectorDirector)
				currentRayReturnedColor := controller.iterateRay(pathTracer, 0, depthIterations, currentRay)
				floatColors[threadRayIndex] = currentRayReturnedColor
				lock.RemoveThread()
			}(rayIndex)
		} else {
			time.Sleep(20)
			rayIndex--
		}
	}
	ended := false
	for !ended {
		if lock.GetThreads() == 0 {
			ended = true
		} else {
			time.Sleep(25)
		}
	}


	return controller.parseRaysColorsToRGB(floatColors, numberOfRays)

}

// Run runs the path tracing.
//
// Parameters:
// 	pathTracer        - The PathTracer.
// 	raysPerPixel      - The number of rays per pixel.
// 	recursions        - The number recursions of each ray.
// 	windowStartLine   - The starting line index of the window of the screen to use the path tracing.
// 	windowStartColumn - The starting column index of the window of the screen to use the path tracing.
// 	windowEndLine     - The ending line index of the window of the screen to use the path tracing.
// 	windowEndColumn   - The ending column index of the window of the screen to use the path tracing.
//
// Returns:
// 	The color matrix representing the rendered image.
//
func (controller *Controller) Run(pathTracer *PathTracer, raysPerPixel, recursions, windowStartLine, windowStartColumn,
	windowEndLine, windowEndColumn int) (*color_matrix.ColorMatrix, error) {
	if windowStartLine < 0 ||
		windowStartLine > windowEndLine ||
		windowEndLine > pathTracer.GetPixelScreen().GetHeight() ||
		windowStartColumn < 0 ||
		windowStartColumn > windowEndColumn ||
		windowEndColumn > pathTracer.GetPixelScreen().GetWidth() {
		return nil, windowError(pathTracer, windowStartLine, windowStartColumn, windowEndLine, windowEndColumn)
	}

	colorMatrix := color_matrix.Init(pathTracer.pixelScreen)
	for lineIndex := windowStartLine; lineIndex < windowEndLine; lineIndex++ {
		for columnIndex := windowStartColumn; columnIndex < windowEndColumn; columnIndex++ {
			pixelColor := controller.traceFirstRays(pathTracer, lineIndex, columnIndex, raysPerPixel, recursions)
			_ = colorMatrix.SetColor(lineIndex, columnIndex, pixelColor)
		}
	}
	return colorMatrix, nil
}
