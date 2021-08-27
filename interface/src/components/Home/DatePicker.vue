<template>
  <div class="main-container">
    <h2>Pick a date to sync the buyers&apos; data!</h2>
    <v-row class="calendar-container" justify="center">
      <v-date-picker
        :disabled="loading"
        color="#190862"
        v-model="date"
        elevation="6"
        full-width
        @input="dateChange"
        :show-current="false"
        :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)"
      />
    </v-row>
  </div>
</template>

<script>
import { API_URL } from '@/constants'

export default {
  name: 'DatePicker',
  props: {
    syncedDate: {
      type: String,
      default: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)
    },
    loading: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      date: this.syncedDate
    }
  },
  watch: {
    syncedDate () {
      this.date = this.syncedDate
      this.dateChange()
    }
  },
  methods: {
    dateChange () {
      if (this.date !== this.syncedDate) this.$emit('date-change', false)
      else this.$emit('date-change', true)
    },
    async syncData () {
      const unixTimestamp = new Date(this.date).getTime() / 1000
      const result = await fetch(`${API_URL}/load-data?date=${unixTimestamp}`)
      if (!result.ok) {
        const jsonResult = await result.json()
        if (jsonResult.result === 'Date already synced') return 'current'
      }
      return [result.ok, this.date]
    }
  }
}
</script>

<style lang="sass" scoped>

.main-container
  margin-top: 24px

h2
  cursor: default
  font-size: 2rem
  font-weight: 400
  margin-bottom: 36px
  text-align: center

.calendar-container
  margin: 0 auto 36px
  width: 90%

</style>
