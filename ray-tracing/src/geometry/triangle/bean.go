package triangle

// Triangle is a class for triangles.
//
// Members:
// 	verticesIndexes        - The 3 indexes of the Triangle is vertices on the point list.
// 	verticesNormalsIndexes - The 3 indexes of the Triangle is vertices normals on the normals list.
//
type Triangle struct {
	verticesIndexes []int
	verticesNormalsIndexes []int
}

// GetVertexIndex gets a vertex index of the Triangle.
//
// Parameters:
// 	index - The index of the vertex on the Triangle vertices indexes.
//
// Returns:
// 	The vertex index.
//  An error.
//
func (triangle *Triangle) GetVertexIndex(index int) (int, error) {
	if index < 0 || index >= 3 {
		return 0, indexError(index)
	}
	return triangle.verticesIndexes[index], nil
}

// GetVertexNormalIndex gets a normal index of the Triangle.
//
// Parameters:
// 	index - The index of the vertex normal on the Triangle vertices normals indexes.
//
// Returns:
// 	The vertex normal index.
//  An error.
//
func (triangle *Triangle) GetVertexNormalIndex(index int) (int, error) {
	if index < 0 || index >= 3 {
		return 0, indexError(index)
	}
	return triangle.verticesNormalsIndexes[index], nil
}

// Init initializes a Triangle.
//
// Parameters:
// 	verticesIndexes        - The 3 indexes of the Triangle is vertices on the point list.
// 	verticesNormalsIndexes - The 3 indexes of the Triangle is vertices normals on the normals list.
//
// Returns:
// 	A Triangle
//  An error
//
func Init(verticesIndexes, verticesNormalsIndexes []int) (*Triangle, error) {
	if len(verticesIndexes) != 3 || len(verticesNormalsIndexes) != 3 {
		return nil, initializationError(verticesIndexes, verticesNormalsIndexes)
	}
	return &Triangle{verticesIndexes: verticesIndexes, verticesNormalsIndexes: verticesNormalsIndexes}, nil
}
