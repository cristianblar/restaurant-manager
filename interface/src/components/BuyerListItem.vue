<template>
  <div class="buyer-container">
    <router-link :to="`/buyers/${buyerId}`">
      <h4>{{ name }} ({{ buyerId }})</h4>
    </router-link>
    <span>{{ age }} years old</span>
    <span>
      <span class="tran-number">{{ transactions }}</span> transactions
    </span>
  </div>
</template>

<script>
export default {
  name: 'BuyerListItem',
  props: {
    name: {
      type: String,
      required: true
    },
    buyerId: {
      type: String,
      required: true
    },
    age: {
      type: Number,
      required: true
    },
    transactions: {
      type: Number,
      required: true
    }
  }
}
</script>

<style lang="sass" scoped>

.buyer-container
  background-color: white
  border: 1px solid #190862
  border-radius: 16px
  color: black
  display: flex
  font-size: 1.8rem
  font-weight: 300
  justify-content: space-between
  margin: 18px 0
  padding: 12px 24px
  text-align: center
  transition: all 200ms

  &:hover
    background-color: #FFC933
    box-shadow: 0 6px 8px -4px #00000080, 5px 5px 15px 5px rgb(0 0 0 / 0%)
    color: #190862
    & a h4
      color: #190862
      text-decoration: underline

  & span
    cursor: default

  & a
    color: black
    text-decoration: none

.tran-number
  text-decoration: underline

</style>
