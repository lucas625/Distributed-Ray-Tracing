package object

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// Object is a class for all data of an Object.
//
// Members:
//  name                 - The name of the Object.
// 	repository           - The point repository.
// 	triangles            - The triangles that form the Object.
//  normals              - The normals of the vertices.
//  lightCharacteristics - The light characteristics of the Object.
//
type Object struct {
	name               string
	repository         *point_repository.PointRepository
	triangles          []*triangle.Triangle
	normals            []*vector.Vector
	lightCharacteristics *lightCharacteristics
}

// GetName gets the name of the Object.
//
// Parameters:
// 	none
//
// Returns:
// 	The name of the Object.
//
func (object *Object) GetName() string {
	return object.name
}

// GetRepository gets the point repository of the Object.
//
// Parameters:
// 	none
//
// Returns:
// 	The points of the Object.
//
func (object *Object) GetRepository() *point_repository.PointRepository {
	return object.repository
}

// GetTriangles gets the triangles of the Object.
//
// Parameters:
// 	none
//
// Returns:
// 	The triangles of the Object.
//
func (object *Object) GetTriangles() []*triangle.Triangle {
	return object.triangles
}

// GetNormals gets the normals of the vertices of the Object.
//
// Parameters:
// 	none
//
// Returns:
// 	The vertices normals of the Object.
//
func (object *Object) GetNormals() []*vector.Vector {
	return object.normals
}

// GetLightCharacteristics gets the light characteristics of the Object.
//
// Parameters:
// 	none
//
// Returns:
// 	The light characteristics of the Object.
//
func (object *Object) GetLightCharacteristics() *lightCharacteristics {
	return object.lightCharacteristics
}

// IsEqual checks if a Object object is equal to another.
//
// Parameters:
// 	other - The other Object.
//
// Returns:
// 	If the objects are equal.
//
func (object *Object) IsEqual(other *Object) bool {
	if len(object.GetTriangles()) != len(other.GetTriangles()) {
		return false
	}
	for triangleIndex := 0; triangleIndex < len(object.GetTriangles()); triangleIndex++ {
		if !object.GetTriangles()[triangleIndex].IsEqual(other.GetTriangles()[triangleIndex]) {
			return false
		}
	}
	if len(object.GetNormals()) != len(other.GetNormals()) {
		return false
	}
	for normalIndex := 0; normalIndex < len(object.GetNormals()); normalIndex++ {
		if !object.GetNormals()[normalIndex].IsEqual(other.GetNormals()[normalIndex]) {
			return false
		}
	}
	return object.GetName() == other.GetName() &&
		object.GetRepository().IsEqual(other.GetRepository()) &&
		object.GetLightCharacteristics().IsEqual(other.GetLightCharacteristics())
}

// Init initializes an Object.
//
// Parameters:
//  name                   - The name of the Object.
// 	repository             - The point repository.
// 	triangles              - The triangles that form the Object.
//  normals                - The normals of the vertices.
//  color                  - RGB for the color of the object.
//  specularDecay          - Constant for how fast the specular component decays.
//  specularReflection     - Percentage of specular rays.
//  roughNess              - How much reflections rays get distorted.
//  transmissionReflection - Percentage of transmission rays.
//  diffuseReflection      - Percentage of diffuse rays.
//
// Returns:
// 	An Object.
// 	An error.
//
func Init(name string, repository *point_repository.PointRepository, triangles []*triangle.Triangle,
	normals []*vector.Vector, color []float64, specularDecay, specularReflection, roughNess, transmissionReflection,
	diffuseReflection float64) (*Object, error) {
	lightCharacteristics, err := initLightCharacteristics(color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	if err != nil {
		return nil, err
	}
	return &Object{name: name, repository: repository, triangles: triangles, normals: normals,
		lightCharacteristics: lightCharacteristics}, nil
}
