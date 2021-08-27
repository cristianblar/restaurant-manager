<template>
  <main>
    <template v-if="!syncedData || error">
      <upper-error-message
        v-if="!syncedData"
        link-text="Go to sync"
        link="/"
        message="No data synced yet! Please pick a date to sync"
      />
      <upper-error-message
        v-else
        link-text="Go back"
        link="/"
        message="Something went wrong... Please, try again"
      />
    </template>
    <Loader v-else-if="loading" />
    <section v-else class="list-container">
      <h2>Buyers |<strong>| {{ syncedDate }}</strong></h2>
      <BuyerList :buyers="buyers" />
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
    </section>
  </main>
</template>

<script>
import BuyerList from '@/components/Buyers/BuyerList'
import UpperErrorMessage from '@/components/UpperErrorMessage'
import Loader from '@/components/Loader'
import { fetchData } from '@/utils'

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
  components: { BuyerList, UpperErrorMessage, Loader },
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
      fetchData('buyers')
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
      fetchData(`buyers?page=${this.nextPage}`)
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
