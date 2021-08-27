<template>
  <section class="exp-container">
    <v-expansion-panels hover>
      <v-expansion-panel>
        <v-expansion-panel-header class="transactions-title">
          Transactions history
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <BuyerTransaction
            v-for="transaction in fixedTransactionsArray"
            :key="transaction['Transaction.id']"
            :transaction-id="transaction['Transaction.id']"
            :transaction-products="transaction['Transaction.products']"
            :transaction-cost="transaction.totalCost"
          ></BuyerTransaction>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </section>
</template>

<script>
import BuyerTransaction from '@/components/BuyerTransaction'

export default {
  name: 'BuyerTransactions',
  props: {
    transactions: {
      type: Array,
      default: () => []
    }
  },
  components: { BuyerTransaction },
  computed: {
    fixedTransactionsArray: function () {
      const transactionsCopy = JSON.parse(JSON.stringify(this.transactions))
      transactionsCopy.sort((a, b) => {
        if (a['Transaction.id'].toUpperCase() > b['Transaction.id'].toUpperCase()) {
          return -1
        }
        if (a['Transaction.id'].toUpperCase() < b['Transaction.id'].toUpperCase()) {
          return 1
        }
        return 0
      })
      for (const transaction of transactionsCopy) {
        transaction.totalCost = transaction['Transaction.products']
          .reduce((totalCost, product) => totalCost + product['Product.price'], 0)
      }
      return transactionsCopy
    }
  }
}
</script>

<style lang="sass" scoped>

.exp-container
  margin-bottom: 24px

.transactions-title
  color: #FFBC00
  font-size: 1.8rem
  font-weight: 500

</style>
