<template>
  <v-app>
    <manager-header :synced-data="syncedData"></manager-header>
    <v-main>
      <router-view
        :synced-data="syncedData"
        :synced-date="syncedDate"
        @data-synced="handleDataSync"
        @date-change="handleDateChange"
      />
    </v-main>
  </v-app>
</template>

<script>
import ManagerHeader from '@/components/ManagerHeader'

export default {
  name: 'App',
  components: { ManagerHeader },
  data: () => ({
    syncedData: true,
    syncedDate: sessionStorage.getItem('syncedDate') || (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)
  }),
  methods: {
    handleDateChange (e) {
      this.syncedData = e
    },
    handleDataSync (e) {
      this.syncedDate = e
      // this.syncedData = true
      sessionStorage.setItem('syncedDate', e)
    }
  }
}
</script>

<style lang="sass">

*
  box-sizing: border-box
  margin: 0
  padding: 0

html
  font-size: 62.5%

body
  font-family: 'Roboto', sans-serif

</style>
