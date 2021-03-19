<template>
  <v-container
    class="fill-height"
    fluid
  >
    <v-row>
      <v-col>
        <v-card>
          <v-card-title>
            <span>Upload Scenes</span>
          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="6">
                <drop-files-zone
                    :file-type="'.json'"
                    :selected-files="selectedFiles"
                    @push-files="selectedFiles.push(...$event)"
                />
                // Add here the list of json UI
              </v-col>

              <v-col cols="6">
                <v-form v-model="isFormValid">
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="pathTracingParameters.raysPerPixel"
                          label="Rays Per Pixel"
                          :rules="[ruleRequiredField]"
                          type="Number"
                      />
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="pathTracingParameters.recursions"
                          label="Recursions"
                          :rules="[ruleRequiredField]"
                          type="Number"
                      />
                    </v-col>
                  </v-row>
                </v-form>
              </v-col>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-btn
                color="primary"
                :disabled="selectedFiles.length === 0 || !isFormValid"
                @click="submit"
            >
              Submit
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

import PathTracingParametersBean from "@/beans/path_tracing_parameters_bean";
import RayTracingControllerService from "@/services/ray_tracing_controller_service";
import DropFilesZone from "@/components/DropFilesZone";

const rayTracingControllerService = new RayTracingControllerService()

export default {
  name: 'RayTracingView',
  components: {DropFilesZone},
  data: function () {
    return {
      selectedFiles: [],
      pathTracingParameters: new PathTracingParametersBean(1, 2),
      isFormValid: false
    }
  },
  methods: {
    /**
     * Peforms the submit of all scenes and downloads the png images.
     */
    submit: function () {

      const reader = new FileReader();
      reader.onload = function(event) {
        return JSON.parse(event.target.result)
      }

      const jsonObj = reader.readAsText(this.selectedFiles[0])

      const rayTracingParameters = {
        jsonObj,
        pathTracingParameters: this.pathTracingParameters
      }

      const successCallBack = (response) => {
        let blob = new Blob([response.data], { type: 'image/png' })
        let link = document.createElement('a')
        link.href = window.URL.createObjectURL(blob)
        link.download = 'scene.png'
        link.click()
      }

      const errorCallBack = (error) => {
        alert('Failed to run path tracing.')
      }

      const finallyCallBack = () => {}

      rayTracingControllerService.renderWithPathTracing(
          rayTracingParameters, successCallBack, errorCallBack, finallyCallBack)
    },
    /**
     * Forces the form field to provide a value.
     * @param {String|Number} value
     * @return {Boolean|String}
     */
    ruleRequiredField: function (value) {
      const msg = 'This field is required'
      let is_valid
      if (typeof value === 'number') {
        is_valid = Boolean(value)
      } else {
        is_valid = value && value.length > 0
      }
      return is_valid || msg
    },
  }
}
</script>
