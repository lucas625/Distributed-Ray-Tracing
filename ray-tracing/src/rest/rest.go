package rest

import (
	"encoding/json"
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/marshaller"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/path_tracing"
	"io/ioutil"
	"net/http"
)

// parsePathTracingRequest parses the request for a path tracing run to the corresponding classes.
//
// Parameters:
// 	request - The request.
//
// Returns:
// 	The PathTracer.
// 	The number of rays per pixel.
// 	The number recursions of each ray.
// 	The starting line index of the window of the screen to use the path tracing.
// 	The starting column index of the window of the screen to use the path tracing.
// 	The ending line index of the window of the screen to use the path tracing.
// 	The ending column index of the window of the screen to use the path tracing.
// 	An error.
//
func parsePathTracingRequest(request *http.Request) (*path_tracing.PathTracer, int, int, int, int, int, int, error) {
	var data map[string]interface{}
	bodyAsBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, errors.New("failed to decode your request")
	}
	err = json.Unmarshal(bodyAsBytes, &data)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, errors.New("failed to parse your request")
	}
	marshallerController := &marshaller.Controller{}
	return marshallerController.ParsePathTracingParametersFromMap(data)
}

// RunPathTracing runs the requested path tracing, sending a matrix of colors as response.
//
// Parameters:
// 	responseWriter - The response writer.
// 	request        - The request.
//
// Returns:
// 	none
//
func RunPathTracing(responseWriter http.ResponseWriter, request *http.Request) {
	pathTracer, raysPerPixel, recursions, windowStartLine, windowStartColumn, windowEndLine, windowEndColumn, err :=
		parsePathTracingRequest(request)

	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
	}

	pathTracingController := path_tracing.Controller{}
	colorMatrix, err := pathTracingController.Run(pathTracer, raysPerPixel, recursions, windowStartLine,
		windowStartColumn, windowEndLine, windowEndColumn)

	if err != nil {
		http.Error(responseWriter, "failed to run the path tracing.", 500)
	}

	colorMatrixAsBytes, err := json.Marshal(colorMatrix)

	if err != nil {
		http.Error(responseWriter, "failed to serialize the response", 500)
	}

	_, err = responseWriter.Write(colorMatrixAsBytes)
	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
	}
}
