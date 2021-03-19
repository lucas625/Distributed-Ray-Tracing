<template>
  <v-card
      flat
      color="grey lighten-4"
      @drop.prevent="receiveFiles($event.dataTransfer.files)"
      @dragover.prevent
      @click="openFileInputNavigator()"
  >
    <v-card-title class="justify-center">
      <span>Drop here the {{ fileType }} files</span>
    </v-card-title>
    <v-card-text class="text-center">
      <v-divider />
      <v-row>
        <v-col>
          <v-icon large>
            mdi-folder-open
          </v-icon>
          <v-file-input
              id="file-input"
              :value="selectedFiles"
              multiple
              :accept="fileType"
              class="d-none"
              @change="receiveFiles($event)"
          />
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script>
import FileHelper from '@/common/file_helper'

export default {
  name: 'DropFilesZone',
  props: {
    selectedFiles: {
      type: Array[File],
      required: true
    },
    fileType: {
      type: String,
      required: true
    }
  },
  methods: {
    /**
     * Opens the file input navigator.
     */
    openFileInputNavigator: function () {
      document.getElementById('file-input').click()
    },

    /**
     * Emits an event to push the file to the file list.
     * @param {[File]} files.
     */
    receiveFiles: function (files) {
      let validFiles = []
      for (const file of files) {
        if (FileHelper.isOfValidType(file, [this.fileType])) { // Add check to see if file is already selected
          validFiles.push(file)
        }
      }
      if (validFiles.length > 0) {
        this.$emit('push-files', validFiles)
      }
    }
  }
}
</script>
