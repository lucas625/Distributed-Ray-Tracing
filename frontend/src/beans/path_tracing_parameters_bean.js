/**
 * Bean for the parameters of a path tracing run.
 */
export default class PathTracingParametersBean {
    /**
     * @param {Number} raysPerPixel - The number of rays per pixel.
     * @param {Number} recursions - The number of recursions of each ray.
     */
    constructor (raysPerPixel, recursions) {
        this.raysPerPixel = raysPerPixel
        this.recursions = recursions
    }
}
