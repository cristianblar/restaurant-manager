<template>
  <main>
    <date-picker
      ref="datePicker"
      :loading="loading"
      :synced-date="syncedDate"
      v-on="$listeners"
    />
    <span class="current-date">
      <strong>Current synced date:</strong> {{ syncedDate }}
    </span>
    <sync-buttons
      :synced-data="syncedData"
      :loading="loading"
      @sync-click="handleSyncClick"
    />
    <h4 v-show="resultMessage.active" :class="!resultMessage.result && 'wrong'">
      {{ resultMessage.result ? resultMessage.goodMessage : resultMessage.errorMessage }}
    </h4>
  </main>
</template>

<script>
import DatePicker from '@/components/Home/DatePicker'
import SyncButtons from '@/components/Home/SyncButtons'

export default {
  name: 'Home',
  props: {
    syncedData: {
      type: Boolean,
      default: true
    },
    syncedDate: {
      type: String,
      default: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)
    }
  },
  data: () => ({
    loading: false,
    resultMessage: {
      active: false,
      result: false,
      goodMessage: 'Data synced successfully',
      errorMessage: 'Something went wrong... Please, try again'
    }
  }),
  components: { DatePicker, SyncButtons },
  methods: {
    handleSyncClick () {
      this.loading = true
      this.$refs.datePicker.syncData()
        .then(([result, date]) => {
          if (result) {
            this.loading = false
            this.resultMessage.result = true
            this.resultMessage.active = true
            setTimeout(() => (this.resultMessage.active = false), 4000)
            this.$emit('data-synced', date)
          } else throw Error('Fetch data error')
        })
        .catch((err) => {
          console.error(err)
          this.loading = false
          this.resultMessage.result = false
          this.resultMessage.active = true
          setTimeout(() => (this.resultMessage.active = false), 4000)
        })
    }
  }
}
</script>

<style lang="sass" scoped>

.current-date
  color: #190862
  cursor: default
  display: block
  font-size: 1.8rem
  font-weight: 300
  margin-bottom: 36px
  text-align: center

h4
  color: #190862
  cursor: default
  font-size: 2rem
  font-weight: 300
  margin-top: 24px
  text-align: center

h4.wrong
  color: #FFBC00

</style>
