/**
 * Bean for the parameters of a path tracing run.
 */
export default class PathTracingParametersBean {
    /**
     * @param {Number} width - The pixels per line.
     * @param {Number} height - The lines.
     * @param {Number} raysPerPixel - The number of rays per pixel.
     * @param {Number} recursions - The number of recursions of each ray.
     */
    constructor (width, height, raysPerPixel, recursions) {
        this.width = width
        this.height = height
        this.raysPerPixel = raysPerPixel
        this.recursions = recursions
    }
}
