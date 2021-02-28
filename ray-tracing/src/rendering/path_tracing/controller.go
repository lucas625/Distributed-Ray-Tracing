package path_tracing

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/line"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/camera"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/color_matrix"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/object"
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

// findNormal finds the resulting normal of a object intersection.
//
// Parameters:
//  intersectedObject      - The object that has the next ray is origin.
//  triangleIndex          - The index of the triangle of the intersected object that hast the point.
//  barycentricCoordinates - The barycentric coordinates of the next ray origin relative to the triangle.
//
// Returns:
// 	The resulting normal.
//
func (controller *Controller) findNormal(intersectedObject *object.Object, triangleIndex int,
	barycentricCoordinates []float64) *vector.Vector {
	normals := make([]*vector.Vector, 3)
	for index := 0; index < 3; index++ {
		normalIndex, _ := intersectedObject.GetTriangles()[triangleIndex].GetVertexNormalIndex(index)
		normals[index] = intersectedObject.GetNormals()[normalIndex]
	}
	vectorController := &vector.Controller{}
	firstPlusSecondNormal, _ := vectorController.Sum(normals[0], normals[1], barycentricCoordinates[0],
		barycentricCoordinates[1])
	sumNormals, _ := vectorController.Sum(firstPlusSecondNormal, normals[2], 1, barycentricCoordinates[2])
	return vectorController.Normalize(sumNormals)
}

// findSpecularReflectionVector finds the resulting normal.
//
// Parameters:
//  pathTracer        - The PathTracer.
//  nextRayOrigin     - The origin of the next ray.
//  intersectedObject - The object that has the next ray is origin.
//  normalVector      - The resulting normal of a object intersection.
//
// Returns:
// 	The specular vector.
//
func (controller *Controller) findSpecularReflectionVector(pathTracer *PathTracer, nextRayOrigin *point.Point,
	intersectedObject *object.Object, normalVector *vector.Vector) *vector.Vector {

	vectorController := &vector.Controller{}
	pointController := &point.Controller{}
	objectController := &object.Controller{}

	lightCenter := objectController.GetCenter(pathTracer.GetLights()[0].GetLightObject())
	lightVector, _ := pointController.ExtractVector(nextRayOrigin, lightCenter)
	normalizedLightVector := vectorController.Normalize(lightVector)

	normalDotProductLight, _ := vectorController.DotProduct(normalVector, normalizedLightVector)

	// R = 2N(N.L) - L
	specularVector, _ := vectorController.Sum(normalVector, normalizedLightVector, 2 * normalDotProductLight, -1)

	offsetVector := RandomInSemiSphereSpecular()
	offsetVectorWithRoughness := vectorController.ScalarMultiplication(offsetVector,
		intersectedObject.GetLightCharacteristics().GetRoughNess())

	resultingSpecularVector, _ := vectorController.Sum(specularVector, offsetVectorWithRoughness, 1, 1)

	return resultingSpecularVector
}

// findNextRay finds the next ray.
//
// Parameters:
//  pathTracer             - The PathTracer.
//  nextRayOrigin          - The origin of the next ray.
//  intersectedObject      - The object that has the next ray is origin.
//  triangleIndex          - The index of the triangle of the intersected object that hast the point.
//  barycentricCoordinates - The barycentric coordinates of the next ray origin relative to the triangle.
//
// Returns:
// 	The next ray.
//
func (controller *Controller) findNextRay(pathTracer *PathTracer, nextRayOrigin *point.Point,
	intersectedObject *object.Object, triangleIndex int, barycentricCoordinates []float64) *line.Line {

	normalVector := controller.findNormal(intersectedObject, triangleIndex, barycentricCoordinates)

	sumOfTotalReflections := intersectedObject.GetLightCharacteristics().GetDiffuseReflection() +
		intersectedObject.GetLightCharacteristics().GetSpecularReflection() // + intersectedObject.TransReflection
	selectedRandomValue := 0.0 + rand.Float64()*sumOfTotalReflections
	var newRayVectorDirector *vector.Vector
	if selectedRandomValue <= intersectedObject.GetLightCharacteristics().GetDiffuseReflection() {
		newRayVectorDirector = RandomInSemiSphere(normalVector, nextRayOrigin)
	} else if selectedRandomValue <= intersectedObject.GetLightCharacteristics().GetDiffuseReflection() +
		intersectedObject.GetLightCharacteristics().GetSpecularReflection() {
		newRayVectorDirector = controller.findSpecularReflectionVector(
			pathTracer, nextRayOrigin, intersectedObject, normalVector)
	} else {
		// TODO: Transmission reflexion.
	}

	vectorController := &vector.Controller{}
	newRayVectorDirector = vectorController.Normalize(newRayVectorDirector)
	newRay, _ := line.Init(nextRayOrigin, newRayVectorDirector)
	return newRay
}

// intersectObjects uses a ray to intersect all objects.
//
// Parameters:
// 	pathTracer          - The PathTracer.
//  currentRay          - The current ray.
//  minimumRayParameter - The minimum value for the ray parametric is parameter.
//
// Returns:
// 	If there is intersections with the objects.
// 	The closest line parameter.
// 	The closest object index.
// 	The closest triangle index.
// 	The closest triangle is barycentric coordinates.
//
func (controller *Controller) intersectObjects(pathTracer *PathTracer, currentRay *line.Line,
	minimumRayParameter float64) (bool, float64, int, int, []float64) {
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

			if hasIntersection && lineParameter >= minimumRayParameter {
				hasObjectIntersections = true
				if  lineParameter < closestLineParameter {
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
// 	pathTracer          - The PathTracer.
//  currentRay          - The current ray.
//  minimumRayParameter - The minimum value for the ray parametric is parameter.
//
// Returns:
// 	If there is intersections with lights.
// 	The closest light is line parameter.
// 	The closest light index.
//
func (controller *Controller) intersectLights(pathTracer *PathTracer, currentRay *line.Line,
	minimumRayParameter float64) (bool, float64, int) {
	closesLightLineParameterIndex := math.MaxFloat64
	closestLightIndex := -1
	hasLightIntersection := false
	rayController := ray.Controller{}

	for lightIndex, currentLight := range pathTracer.GetLights() {
		for _, currentTriangle := range currentLight.GetLightObject().GetTriangles() {
			lineParameter, _, hasIntersection, _ := rayController.IntersectRayTriangle(
				currentRay, currentTriangle, currentLight.GetLightObject().GetRepository())

			if hasIntersection && lineParameter >= minimumRayParameter {
				hasLightIntersection = true
				if lineParameter <= closesLightLineParameterIndex {
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

	var minimumRayParameter float64

	if currentIteration == 0 {
		minimumRayParameter = 0
	} else {
		minimumRayParameter = 1
	}

	hasObjectIntersection, closestLineParameter, closesObjectIndex, closestTriangleIndex,
		closestTriangleBarycentricCoordinates := controller.intersectObjects(pathTracer, currentRay,
			minimumRayParameter)
	hasLightIntersection, closestLightLineParameter, closestLight := controller.intersectLights(
		pathTracer, currentRay, minimumRayParameter)

	if hasLightIntersection && closestLightLineParameter <= closestLineParameter {
		intersectedLight := pathTracer.GetLights()[closestLight]
		for index := 0; index < 3; index++ {
			color[index] = intersectedLight.GetColor()[index] * intersectedLight.GetLightIntensity()
		}

	} else {
		if hasObjectIntersection {
			var colorAux []float64
			if currentIteration == 0 {
				colorAux = []float64{0, 0, 0}
			} else {
				colorAux = []float64{1, 1, 1}
			}
			if depthIterations > 0 {
				lineController := line.Controller{}
				newRayStartingPoint, _ := lineController.FindPoint(currentRay, closestLineParameter)
				newRay := controller.findNextRay(pathTracer, newRayStartingPoint,
					pathTracer.GetObjects()[closesObjectIndex], closestTriangleIndex,
					closestTriangleBarycentricCoordinates)
				colorAux = controller.iterateRay(pathTracer, currentIteration+1, depthIterations, newRay)
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
