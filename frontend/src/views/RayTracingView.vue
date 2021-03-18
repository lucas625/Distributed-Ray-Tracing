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
              </v-col>

              <v-col cols="6">

              </v-col>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-btn @click="submit">
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
      pathTracingParameters: new PathTracingParametersBean(1, 2)
    }
  },
  methods: {
    /**
     * Peforms the submit of all scenes and downloads the png images.
     */
    submit: function () {
      if (this.$refs.analysisForm.validate()) {

        const rayTracingParameters = {
          ...JSON.parse(this.selectedFiles[0]),
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
      }
    }
  }
}
</script>
