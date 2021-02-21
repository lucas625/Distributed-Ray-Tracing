package path_tracing

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/color_matrix"
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
func (ptracer *PathTracer) FindNextRay(pos entity.Point, obj general.Object, triangleIdx int, bCoords []float64) entity.Line {
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

// TraceRay is a function to trace a ray through a pixel.
//
// Parameters:
// 	lp             - pixel line index.
//  cp             - pixel column index.
//  rays           - number of rays per pixel.
//  recursions     - number of recursions.
//
// Returns:
// 	the colored screen painted at that position.
//
func (ptracer *PathTracer) TraceRay(lp, cp, rays, recursions int) []int {
	rand.Seed(time.Now().UnixNano())
	floatColors := make([][]float64, rays)
	lock := Locker{threads:0}
	for ray := 0; ray < rays; ray++ {
		if lock.getThreads() < 12 {
			lock.addThreads()
			go func(inRay int) {
				offx := rand.Float64()
				offy := rand.Float64()

				screenV := ptracer.PixelScreen.PixelToWorld(lp, cp, 1.0, offx, offy, ptracer.Cam.FieldOfView)
				line := entity.Line{Start: ptracer.Cam.Pos, Director: screenV}

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
							if t >= 1 && t < closestT {
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
							if t >= 1 && t <= closestT {
								lightClosest = true
								closestT = t
								closestObjIdx = lgtIdx
							}
						}
					}
				}

				if !lightClosest {
					if closestObjIdx != -1 {
						colorAux := []float64{0, 0, 0}
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
				floatColors[inRay] = color
				lock.removeThreads()
			}(ray)
		} else {
			time.Sleep(20)
			ray--
		}
	}
	ended := false
	for !ended {
		if lock.getThreads() == 0 {
			ended = true
		} else {
			time.Sleep(25)
		}
	}

	// calculating average
	color := make([]float64, 3)
	for i := 0; i < rays; i++ {
		for j := 0; j < 3; j++ {
			color[j] = color[j] + floatColors[i][j]
		}
	}

	intColor := make([]int, 3)
	for i := 0; i < 3; i++ {
		intColor[i] = int(math.Floor(math.Sqrt(color[i]/float64(rays)) * 255))
		if intColor[i] > 255 {
			intColor[i] = 255
		} else if intColor[i] < 0 {
			intColor[i] = 0
		}
	}
	return intColor

}

// Run is a function to run the path tracing.
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
			pixelColor := controller.TraceRay(pathTracer, lineIndex, columnIndex, raysPerPixel, recursions)
			colorMatrix.SetColor(lineIndex, columnIndex, pixelColor)
		}
	}
	return colorMatrix, nil
}
