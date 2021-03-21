<template>
  <v-dialog
      v-model="isLoading"
      persistent
  >
    <v-card
        color="primary"
        dark
    >
      <v-card-title>
        {{ loadingMessage }}
      </v-card-title>
      <v-card-text>
        <v-progress-linear
            indeterminate
            color="white"
        />
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'LoadingBar',
  props: {
    loadingMessage: {
      type: String,
      required: true
    },
    isLoadingProps: {
      type: Boolean,
      required: true
}
  },
  data: function () {
    return {
      isLoading: this.isLoadingProps
    }
  },
  watch: {
    isLoadingProps: function (newIsLoadingProps) {
      this.isLoading = newIsLoadingProps
    }
  }
}
</script>
