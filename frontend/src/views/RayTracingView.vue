<template>
  <v-container
    class="fill-height"
    fluid
  >
    <loading-bar
      loading-message="Uploading Scenes"
      :is-loading-props="isUploading"
    />
    <h1>
      Render scenes with path tracing
    </h1>
    <v-row>
      <v-col>
        <v-card>
          <v-card-title>
            <span>Upload Scenes</span>
          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col
                  lg="6"
                  cols="12"
              >
                <drop-files-zone
                    :file-type="'.json'"
                    :selected-files="selectedFiles"
                    @push-files="selectedFiles.push(...$event)"
                />
                <file-list
                    v-if="selectedFiles.length > 0"
                    :file-list="selectedFiles"
                    @clear-all="selectedFiles=[]"
                    @remove="selectedFiles.splice($event, 1)"
                />
              </v-col>

              <v-col
                  lg="6"
                  cols="12"
              >
                <v-form v-model="isFormValid">
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="pathTracingParameters.width"
                          label="Width"
                          :rules="[ruleRequiredField, ruleMinMaxValueField(100, 1080)]"
                          type="Number"
                      />
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="pathTracingParameters.height"
                          label="Height"
                          :rules="[ruleRequiredField, ruleMinMaxValueField(100, 768)]"
                          type="Number"
                      />
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="pathTracingParameters.raysPerPixel"
                          label="Rays Per Pixel"
                          :rules="[ruleRequiredField, ruleMinMaxValueField(1, 400)]"
                          type="Number"
                      />
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="pathTracingParameters.recursions"
                          label="Recursions"
                          :rules="[ruleRequiredField, ruleMinMaxValueField(1, 5)]"
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
                :disabled="selectedFiles.length === 0 || !isFormValid || isUploading"
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

import FileHelper from "@/common/file_helper"
import PathTracingParametersBean from "@/beans/path_tracing_parameters_bean"
import RayTracingControllerService from "@/services/ray_tracing_controller_service"
import DropFilesZone from "@/components/DropFilesZone"
import LoadingBar from "@/components/LoadingBar";
import FileList from "@/components/FileList";

const rayTracingControllerService = new RayTracingControllerService()

export default {
  name: 'RayTracingView',
  components: {FileList, LoadingBar, DropFilesZone},
  data: function () {
    return {
      selectedFiles: [],
      pathTracingParameters: new PathTracingParametersBean(400, 400, 1, 2),
      isFormValid: false,
      uploadingCount: 0
    }
  },
  computed: {
    isUploading: function () {
      return this.uploadingCount > 0
    }
  },
  methods: {
    /**
     * Performs the submit of all scenes and downloads the png images.
     */
    submit: async function () {
      this.uploadingCount = this.selectedFiles.length

      for (const selectedFile of this.selectedFiles) {
        const jsonObj = await FileHelper.readJsonFile(selectedFile)

        const rayTracingParameters = {
          ...jsonObj,
          pathTracingParameters: {
            raysPerPixel: Number(this.pathTracingParameters.raysPerPixel),
            recursions: Number(this.pathTracingParameters.recursions)
          },
          pixelScreen: {
            width: Number(this.pathTracingParameters.width),
            height: Number(this.pathTracingParameters.height)
          }
        }

        const successCallBack = (response) => {
          let blob = new Blob([response.data], { type: 'image/png' })
          let link = document.createElement('a')
          link.href = window.URL.createObjectURL(blob)
          link.download = `${FileHelper.getFilenameWithoutExtension(selectedFile)}-w${this.pathTracingParameters.width}-h${this.pathTracingParameters.height}-rpp${this.pathTracingParameters.raysPerPixel}-r${this.pathTracingParameters.recursions}.png`
          link.click()
        }

        const errorCallBack = (error) => {
          alert('Failed to run path tracing.')
        }

        const finallyCallBack = () => {
          this.uploadingCount--
        }

        rayTracingControllerService.renderWithPathTracing(
            rayTracingParameters, successCallBack, errorCallBack, finallyCallBack)
      }
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
    /**
     * Forces the form field to provide a value between the min and the max.
     * @param {Number} min
     * @param {Number} max
     * @return {function}
     */
    ruleMinMaxValueField: function (min, max) {
      return function(value) {
        let msg
        let is_valid = true
        if (value < min) {
          is_valid = false
          msg = `The value must be at least ${min}.`
        } else if (value > max) {
          is_valid = false
          msg = `The value must not be greater than ${max}.`
        }
        return is_valid || msg
      }
    }
  }
}
</script>
