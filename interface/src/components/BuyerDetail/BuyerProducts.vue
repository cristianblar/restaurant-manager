<template>
  <section class="exp-container">
    <v-expansion-panels hover>
      <v-expansion-panel>
        <v-expansion-panel-header class="products-title">
          Recommended products
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-list-item
            v-for="(product, idx) in products"
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
  </section>
</template>

<script>
export default {
  name: 'BuyerProducts',
  props: {
    products: {
      type: Array,
      default: () => []
    }
  }
}
</script>

<style lang="sass" scoped>

.exp-container
  margin-bottom: 36px

.products-title
  color: #FFBC00
  font-size: 1.8rem
  font-weight: 500

.product-name
  font-size: 1.6rem
  font-weight: 300

.product-price
  font-size: 1.4rem
  font-weight: 300

</style>
