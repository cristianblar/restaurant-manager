<template>
  <div>
    <div v-if="!syncedData || error" class="no-sync">
      <div v-if="!syncedData">
        <h2>No data synced yet! Please pick a date to sync</h2>
        <router-link class="back-link" to="/">Go to sync</router-link>
      </div>
      <div v-else>
        <h2>Something went wrong... Please, try again</h2>
        <router-link class="back-link" to="/">Go back</router-link>
      </div>
    </div>
    <div v-else-if="loading">
      <GridLoader
        class="loader"
        color="#FFBC00"
        :size="60"
      ></GridLoader>
    </div>
    <div class="list-container" v-else>
      <h2>Buyers |<strong>| {{ syncedDate }}</strong></h2>
      <BuyerList :buyers="buyers"></BuyerList>
    </div>
    <v-btn
      v-if="!allBuyersRetrieved && totalBuyers"
      :loading="loadingButton"
      :disabled="loadingButton"
      class="load-button"
      rounded
      color="#FFBC00"
      @click="handleMoreLoad"
    >
      Load more buyers...
    </v-btn>
    <h4 v-show="moreLoadError" class="more-error">
      Something went wrong... Please, try again
    </h4>
  </div>
</template>

<script>
import { GridLoader } from '@saeris/vue-spinners'
import BuyerList from '@/components/BuyerList'
import { API_URL } from '@/constants'

export default {
  name: 'Buyers',
  props: {
    syncedData: {
      type: Boolean,
      default: false
    },
    syncedDate: {
      type: String,
      required: true
    }
  },
  components: { GridLoader, BuyerList },
  data () {
    return {
      buyers: [],
      totalBuyers: null,
      nextPage: 1,
      loading: false,
      loadingButton: false,
      error: false,
      moreLoadError: false
    }
  },
  computed: {
    allBuyersRetrieved: function () {
      return this.buyers.length === this.totalBuyers
    }
  },
  created () {
    if (this.syncedData) {
      this.error = false
      this.loading = true
      fetch(`${API_URL}/api/buyers`)
        .then(response => {
          if (response.ok) return response.json()
          else throw Error('Fetch data error')
        })
        .then(jsonResponse => {
          this.buyers = jsonResponse.results
          this.totalBuyers = jsonResponse['Buyers.total']
          ++this.nextPage
          this.loading = false
        })
        .catch(() => {
          this.loading = false
          this.error = true
        })
    }
  },
  methods: {
    handleMoreLoad () {
      this.loadingButton = true
      fetch(`${API_URL}/api/buyers?page=${this.nextPage}`)
        .then(response => {
          if (response.ok) return response.json()
          else throw Error('Fetch data error')
        })
        .then(jsonResponse => {
          this.buyers = this.buyers.concat(jsonResponse.results)
          this.totalBuyers = jsonResponse['Buyers.total']
          ++this.nextPage
          this.loadingButton = false
        })
        .catch(() => {
          this.loadingButton = false
          this.moreLoadError = true
          setTimeout(() => (this.moreLoadError = false), 3000)
        })
    }
  }
}
</script>

<style lang="sass" scoped>

.no-sync
  & h2
    color: #FFBC00
    font-size: 3.6rem
    font-weight: 900
    margin-bottom: 60px
    text-align: center

  & .back-link
    color: #190862
    display: block
    font-size: 3rem
    font-weight: 500
    margin: 0 auto
    text-decoration: none
    width: fit-content

    &:hover
      text-decoration: underline

.loader
  margin: 60px auto 0

.list-container
  margin: 0 auto
  width: 90%

  & h2
    color: #FFBC00
    cursor: default
    font-size: 3rem
    font-weight: 500
    margin-bottom: 36px
    text-align: center

    & strong
      color: #190862

.load-button
  color: #190862
  display: block
  font-size: 1.4rem
  font-weight: 400
  margin: 36px auto

.more-error
  color: #FFBC00
  cursor: default
  font-size: 2rem
  font-weight: 400
  margin: -12px 0 24px
  text-align: center

</style>
