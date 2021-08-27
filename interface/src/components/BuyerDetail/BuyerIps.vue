<template>
  <section>
    <v-card
      class="card-main-container"
      elevation="6"
      width="100%"
    >
      <span class="table-title">Other buyers using the same devices</span>
      <v-virtual-scroll
        :bench="6"
        :items="otherBuyers"
        height="250"
        item-height="50"
      >
        <template v-slot:default="{ item }">
          <v-list-item :key="`IP${item['Buyer.id']}`">
            <v-list-item-content>
              <v-list-item-title class="ip-buyer">
                {{ item['Buyer.name'] }} ({{ item['Buyer.id'] }})
              </v-list-item-title>
            </v-list-item-content>
            <v-list-item-action class="open-button">
              <a :href="`/buyers/${item['Buyer.id']}`" target="_blank">
                <v-icon small>
                  mdi-open-in-new
                </v-icon>
              </a>
            </v-list-item-action>
          </v-list-item>
          <v-divider />
        </template>
      </v-virtual-scroll>
    </v-card>
  </section>
</template>

<script>
export default {
  name: 'BuyerIps',
  props: {
    otherBuyers: {
      type: Array,
      default: () => []
    }
  }
}
</script>

<style lang="sass" scoped>

a
  text-decoration: none

.card-main-container
  padding-top: 16px

.table-title
  color: #FFBC00
  display: block
  font-size: 1.8rem
  font-weight: 500
  padding: 0 24px 28px

.ip-buyer
  color: #190862
  font-size: 1.6rem
  font-weight: 300

.open-button
  cursor: pointer

</style>
