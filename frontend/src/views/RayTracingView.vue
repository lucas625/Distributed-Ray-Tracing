<template>
  <v-container
    class="fill-height"
    fluid
  >
    <loading-bar
      loading-message="Uploading Scenes"
      :is-loading-props="isUploading"
    />
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
                          :rules="[ruleRequiredField]"
                          type="Number"
                      />
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="pathTracingParameters.height"
                          label="Height"
                          :rules="[ruleRequiredField]"
                          type="Number"
                      />
                    </v-col>
                  </v-row>
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
          link.download = 'scene.png'
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
    }
  }
}
</script>
