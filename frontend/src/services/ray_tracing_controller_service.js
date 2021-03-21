import axios from 'axios'

/**
 * A access layer to the ray tracing controller service.
 * @constructor
 */
export default class RayTracingControllerService {
  /**
   * {RayTracingControllerService} constructor.
   */
  constructor () {
    this.client = axios.create({ baseURL: `${process.env.VUE_APP_RAY_TRACING_CONTROLLER_URL}api/path-tracing` })
  }

  /**
   * Downloads a png image of the scene rendered with path tracing.
   * @param {object} data - The parameters to the analysis.
   * @param {function} successCallBack - The function to be performed on success.
   * @param {function} errorCallback - The function to be performed on error.
   * @param {function} finallyCallback - The function to be performed after the success/error callback.
   */
  renderWithPathTracing (data, successCallBack, errorCallback, finallyCallback) {
    this.client.post('', data, { responseType: 'arraybuffer' })
      .then(successCallBack)
      .catch(errorCallback)
      .then(finallyCallback)
  }
}
