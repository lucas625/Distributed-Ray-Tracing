package math_helper

// MathHelper is a class for vectors.
//
// Members:
// 	tolerance - The tolerance to be used on checks.
//
type MathHelper struct {
	tolerance float64
}

// GetTolerance gets the tolerance of the MathHelper.
//
// Parameters:
// 	none
//
// Returns:
// 	The tolerance of the MathHelper.
//
func (mathHelper *MathHelper) GetTolerance() float64 {
	return mathHelper.tolerance
}

// IsEqualWithTolerance gets a coordinate of a Vector.
//
// Parameters:
// 	firstValue  - The first value.
// 	secondValue - The second value.
//
// Returns:
// 	If both values are equal.
//
func (mathHelper *MathHelper) IsEqualWithTolerance(firstValue, secondValue float64) bool {
	return firstValue - mathHelper.GetTolerance() <= secondValue &&
		secondValue <= firstValue + mathHelper.GetTolerance()
}

// Init initiates a MathHelper.
//
// Parameters:
// 	tolerance  - The tolerance.
//
// Returns:
// 	A MathHelper.
//
func Init(tolerance float64) *MathHelper {
	return &MathHelper{tolerance: tolerance}
}
