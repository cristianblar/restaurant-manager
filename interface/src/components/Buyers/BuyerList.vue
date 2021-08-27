<template>
  <ul>
    <li v-for="(buyer, idx) in buyers" :key="idx + buyer['Buyer.id']">
      <BuyerListItem
        :name="buyer['Buyer.name']"
        :buyer-id="buyer['Buyer.id']"
        :age="buyer['Buyer.age']"
        :transactions="buyer['Buyer.transactionsAmount']"
      />
    </li>
  </ul>
</template>

<script>
import BuyerListItem from '@/components/Buyers/BuyerListItem'

export default {
  name: 'BuyerList',
  props: {
    buyers: {
      type: Array,
      default: () => []
    }
  },
  components: { BuyerListItem }
}
</script>

<style lang="sass" scoped>

ul
  list-style: none
  padding: 0

</style>
