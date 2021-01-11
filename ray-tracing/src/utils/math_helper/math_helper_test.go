package math_helper

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestMathHelper_Init tests the instantiation of a MathHelper.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMathHelper_Init(t *testing.T) {
	tolerance := 10E-10
	mathHelper := Init(tolerance)
	test_helpers.AssertEqual(t, tolerance, mathHelper.tolerance)
}

// TestMathHelper_GetTolerance tests the get tolerance of a MathHelper.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMathHelper_GetTolerance(t *testing.T) {
	tolerance := 10E-10
	mathHelper := Init(tolerance)
	test_helpers.AssertEqual(t, tolerance, mathHelper.GetTolerance())
}

// TestMathHelper_IsEqualWithTolerance tests if two values are equal based on a tolerance.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMathHelper_IsEqualWithTolerance(t *testing.T) {
	tolerance := 3.0
	mathHelper := Init(tolerance)
	firstValue := 5.0
	secondValue := 8.0

	areEqual := mathHelper.IsEqualWithTolerance(firstValue, secondValue)
	test_helpers.AssertEqual(t, true, areEqual)
}

// TestMathHelper_IsEqualWithTolerance_Different tests if two values are different based on a tolerance.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMathHelper_IsEqualWithTolerance_Different(t *testing.T) {
	tolerance := 3.0
	mathHelper := Init(tolerance)
	firstValue := 5.0
	secondValue := 8.1

	areEqual := mathHelper.IsEqualWithTolerance(firstValue, secondValue)
	test_helpers.AssertEqual(t, false, areEqual)
}
