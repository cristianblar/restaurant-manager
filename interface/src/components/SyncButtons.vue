<template>
  <div>
    <v-btn
      class="sync-button"
      elevation="6"
      large
      rounded
      :loading="loading"
      :disabled="loading || syncedData"
      color="#190862"
      @click="syncClick"
    >
      SYNC
      <template v-slot:loader>
        <span class="custom-loader">
          <v-icon light>mdi-cached</v-icon>
        </span>
      </template>
    </v-btn>
    <v-btn
      :disabled="!syncedData"
      class="see-button"
      outlined
      large
      color="#190862"
      @click="goToBuyers"
    >
      SEE BUYERS
    </v-btn>
  </div>
</template>

<script>
export default {
  name: 'SyncButtons',
  props: {
    syncedData: {
      type: Boolean,
      default: false
    },
    loading: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    syncClick () {
      this.$emit('sync-click')
    },
    goToBuyers () {
      this.$router.push('/buyers')
    }
  }
}
</script>

<style lang="sass" scoped>

.sync-button
  color: #FFBC00
  display: block
  font-size: 1.6rem
  margin: 0 auto

.see-button
  display: block
  font-size: 1.6rem
  margin: 36px auto 0

.custom-loader
  animation: loader 1s infinite
  display: flex

@keyframes loader
  from
    transform: rotate(0)
  to
    transform: rotate(360deg)

</style>
