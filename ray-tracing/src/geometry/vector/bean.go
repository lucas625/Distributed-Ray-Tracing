package vector

// Vector is a class for vectors.
//
// Members:
// 	coordinates - List of Coordinates.
//
type Vector struct {
	coordinates []float64
}

// GetCoordinate gets a coordinate of a Vector.
//
// Parameters:
// 	index - The index of the coordinate.
//
// Returns:
// 	The coordinates of the vector.
//
func (vector *Vector) GetCoordinate(index int) (float64, error) {
	if index < 0 || index >= vector.Dimension() {
		return 0, indexError(vector, index)
	}
	return vector.coordinates[index], nil
}

// SetCoordinate is the setter for a Vector's coordinates.
//
// Parameters:
// 	index          - the index of the new coordinate.
// 	newCoordinates - the new coordinate.
//
// Returns:
// 	none
//
func (vector *Vector) SetCoordinate(index int, newCoordinate float64) error {
	if index < 0 || index >= vector.Dimension() {
		return indexError(vector, index)
	}
	vector.coordinates[index] = newCoordinate
	return nil
}

// Dimension gets the dimension of the vector.
//
// Parameters:
// 	none
//
// Returns:
// 	the number of dimensions of the vector.
//
func (vector *Vector) Dimension() int {
	return len(vector.coordinates)
}

// IsEqualDimension checks if two vectors are of the same dimension.
//
// Parameters:
// 	other - The second vector.
//
// Returns:
// 	If the vectors are of the same dimension.
//
func (vector *Vector) IsEqualDimension(other *Vector) bool {
	return vector.Dimension() == other.Dimension()
}

// IsEqual checks if two vectors are equal.
//
// Parameters:
// 	other - The second vector.
//
// Returns:
// 	If the two vectors are equal.
//
func (vector *Vector) IsEqual(other *Vector) bool {
	if !vector.IsEqualDimension(other) {
		return false
	}
	for index := 0; index < vector.Dimension(); index++ {
		vectorCoordinate, _ := vector.GetCoordinate(index)
		otherCoordinate, _ := other.GetCoordinate(index)
		if vectorCoordinate != otherCoordinate {
			return false
		}
	}
	return true
}

// CopyAllCoordinates gets all values of the Vector as a copy.
//
// Parameters:
// 	none
//
// Returns:
// 	All the values of the vector.
//
func (vector *Vector) CopyAllCoordinates() []float64 {
	copiedVectorValues := make([]float64, vector.Dimension())
	for index := 0; index < vector.Dimension(); index++ {
		vectorValue, _ := vector.GetCoordinate(index)
		copiedVectorValues[index] = vectorValue
	}
	return copiedVectorValues
}

// Init is a function to initialize a Vector.
//
// Parameters:
// 	dimension - The dimension of the Vector.
//
// Returns:
// 	A Vector.
//	An error.
//
func Init(dimension int) (*Vector, error) {
	if dimension < 0 {
		return nil, negativeDimensionError(dimension)
	}
	return &Vector{coordinates: make([]float64, dimension)}, nil
}
