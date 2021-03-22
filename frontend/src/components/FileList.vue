<template>
  <v-card flat>
    <v-card-title>
      Selected Files
    </v-card-title>

    <v-card-subtitle>
      <span>Size of the file list: {{ sumFileSizes() + 'MB' }}</span>
    </v-card-subtitle>

    <v-card-text>
      <v-list rounded>
        <v-list-item-group>
          <v-list-item
              v-for="(file, index) in fileList"
              :key="index"
          >
            <v-list-item-icon>
              <v-chip>
                {{ index + 1 }}
              </v-chip>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>
                {{ file.name }}
              </v-list-item-title>
              <v-list-item-subtitle>
                {{ convertBytesToMegaBytes(file.size) + 'MB' }}
              </v-list-item-subtitle>
            </v-list-item-content>

            <v-list-item-action>
              <v-btn
                  fab
                  x-small
                  @click="$emit('remove', index)"
              >
                <v-icon>
                  mdi-delete
                </v-icon>
              </v-btn>
            </v-list-item-action>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-card-text>

    <v-card-actions>
      <v-spacer />
      <v-btn @click="$emit('clear-all')">
        <v-icon left>
          mdi-delete-sweep
        </v-icon>
        Clear
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import FileHelper from '@/common/file_helper'

export default {
  name: 'FileList',
  props: {
    fileList: {
      type: Array[File],
      required: true
    }
  },
  methods: {
    /**
     * Sums the sizes of the files in MBs.
     * @return {Number}
     */
    sumFileSizes: function () {
      const size = FileHelper.sumFileSizes(this.fileList)
      return this.convertBytesToMegaBytes(size)
    },
    convertBytesToMegaBytes: FileHelper.convertBytesToMegaBytes
  }
}
</script>
