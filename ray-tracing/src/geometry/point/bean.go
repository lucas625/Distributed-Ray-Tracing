package point

// Point is a class for points.
//
// Members:
// 	coordinates - List of coordinates.
//
type Point struct {
	coordinates []float64
}

// GetCoordinate gets a coordinate of the Point.
//
// Parameters:
// 	index - The index of the coordinate.
//
// Returns:
// 	The coordinate.
//  An error.
//
func (point *Point) GetCoordinate(index int) (float64, error) {
	if index < 0 || index >= point.Dimension() {
		return 0, indexError(point, index)
	}
	return point.coordinates[index], nil
}

// SetCoordinate sets a coordinate of the Point.
//
// Parameters:
// 	index         - The index of the coordinate.
//  newCoordinate - The new coordinate.
//
// Returns:
// 	An error.
//
func (point *Point) SetCoordinate(index int, newCoordinate float64) error {
	if index < 0 || index >= point.Dimension() {
		return indexError(point, index)
	}
	point.coordinates[index] = newCoordinate
	return nil
}

// Dimension gets the dimension of the Point.
//
// Parameters:
// 	none
//
// Returns:
// 	The dimension of the Point.
//
func (point *Point) Dimension() int {
	return len(point.coordinates)
}

// Init initializes a Point.
//
// Parameters:
// 	dimension - The dimension of the Point.
//
// Returns:
// 	A Point
//  An error
//
func Init(dimension int) (*Point, error) {
	if dimension < 0 {
		return nil, invalidDimensionError(dimension)
	}
	return &Point{coordinates: make([]float64, dimension)}, nil
}
