package path_tracing

import (
	"errors"
	"fmt"
)

// windowError is the error where the window of the screen upon which we are trying to calculate the path tracing
// is invalid.
//
// Parameters:
// 	pathTracer        - The PathTracer.
// 	windowStartLine   - The starting line index of the window of the screen to use the path tracing.
// 	windowStartColumn - The starting column index of the window of the screen to use the path tracing.
// 	windowEndLine     - The ending line index of the window of the screen to use the path tracing.
// 	windowEndColumn   - The ending column index of the window of the screen to use the path tracing.
//
// Returns:
//  An Error.
//
func windowError(pathTracer *PathTracer, windowStartLine, windowStartColumn, windowEndLine, windowEndColumn int) error {
	errorMessage := fmt.Sprintf("Window error. Expected from [(0,0), (%d,%d)] and got [(%d,%d), (%d,%d)]",
		pathTracer.GetPixelScreen().GetHeight(), pathTracer.GetPixelScreen().GetWidth(),
		windowStartLine, windowStartColumn, windowEndLine, windowEndColumn)
	return errors.New(errorMessage)
}
