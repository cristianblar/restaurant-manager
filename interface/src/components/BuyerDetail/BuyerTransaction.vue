<template>
  <article class="transaction-container">
    <v-expansion-panels hover>
      <v-expansion-panel>
        <v-expansion-panel-header class="transaction-title">
          <v-list-item-content>
            <v-list-item-title class="transaction-title-id">
              {{ transactionId }}
            </v-list-item-title>
            <v-list-item-subtitle class="transaction-cost">
              ${{ transactionCost.toFixed(2) }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-list-item
            v-for="(product, idx) in transactionProducts"
            :key="idx + product['Product.id']"
            two-line
          >
            <v-list-item-content>
              <v-list-item-title class="product-name">
                {{ product['Product.name'] }} ({{ product['Product.id'] }})
              </v-list-item-title>
              <v-list-item-subtitle class="product-price">
                ${{ product['Product.price'] }}
              </v-list-item-subtitle>
              <v-list-item-subtitle>
                <v-divider />
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </article>
</template>

<script>
export default {
  name: 'BuyerTransaction',
  props: {
    transactionId: {
      type: String,
      required: true
    },
    transactionProducts: {
      type: Array,
      required: true
    },
    transactionCost: {
      type: Number,
      required: true
    }
  }
}
</script>

<style lang="sass" scoped>

.transaction-container
  margin-bottom: 8px

.transaction-title-id
  font-size: 1.6rem
  font-weight: 300

.transaction-cost
  color: rgba(0, 0, 0, 0.6)
  font-size: 1.4rem
  font-weight: 300

.product-name
  font-size: 1.4rem
  font-weight: 300

.product-price
  font-size: 1.2rem
  font-weight: 300

</style>
