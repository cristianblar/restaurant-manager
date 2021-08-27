<template>
  <div>
    <UpperErrorMessage
      v-if="error"
      link-text="Go to buyers' list"
      link="/buyers"
      message="The ID doesn't exist or the buyer doesn't have transactions for the synced date"
    ></UpperErrorMessage>
    <Loader v-else-if="loading"></Loader>
    <div v-else class="main-container">
      <BuyerProfile
        :name="buyerInfo.owner[0]['Buyer.name']"
        :buyer-id="buyerInfo.owner[0]['Buyer.id']"
        :age="buyerInfo.owner[0]['Buyer.age']"
      ></BuyerProfile>
      <BuyerTransactions
        :transactions="buyerInfo.owner[0]['Buyer.transactions']"
      ></BuyerTransactions>
      <BuyerProducts
        :products="filteredRecommendedProducts"
      ></BuyerProducts>
      <BuyerIps
        :other-buyers="filteredOtherBuyers"
      ></BuyerIps>
    </div>
  </div>
</template>

<script>
import BuyerProfile from '@/components/BuyerProfile'
import BuyerTransactions from '@/components/BuyerTransactions'
import BuyerProducts from '@/components/BuyerProducts'
import BuyerIps from '@/components/BuyerIps'
import UpperErrorMessage from '@/components/UpperErrorMessage'
import Loader from '@/components/Loader'
import { API_URL } from '@/constants'

export default {
  name: 'Buyer',
  components: {
    BuyerProfile,
    BuyerTransactions,
    BuyerProducts,
    BuyerIps,
    UpperErrorMessage,
    Loader
  },
  props: {
    syncedData: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      buyerInfo: {
        owner: [],
        otherBuyers: [],
        otherProducts: []
      },
      loading: false,
      error: false
    }
  },
  computed: {
    filteredRecommendedProducts: function () {
      const productsMap = new Map()
      for (const product of this.buyerInfo.otherProducts) {
        if (!productsMap.has(product['Product.id'])) {
          productsMap.set(product['Product.id'], product)
        }
      }
      const productsCopy = Array.from(productsMap.values())
      productsCopy.sort((a, b) => b['Product.price'] - a['Product.price'])
      return productsCopy.slice(0, 10)
    },
    filteredOtherBuyers: function () {
      const buyersMap = new Map()
      for (const buyer of this.buyerInfo.otherBuyers) {
        if (!buyersMap.has(buyer['Buyer.id'])) {
          buyersMap.set(buyer['Buyer.id'], buyer)
        }
      }
      return Array.from(buyersMap.values())
    }
  },
  created () {
    if (this.syncedData) {
      this.error = false
      this.loading = true
      fetch(`${API_URL}/api/buyers/${this.$route.params.id}`)
        .then(response => {
          if (response.ok) return response.json()
          else throw Error('Fetch data error')
        })
        .then(jsonResponse => {
          this.buyerInfo = jsonResponse
          this.loading = false
        })
        .catch(() => {
          this.loading = false
          this.error = true
        })
    }
  }
}
</script>

<style lang="sass" scoped>

.main-container
  margin: 0 auto
  padding-bottom: 36px
  width: 90%

</style>
