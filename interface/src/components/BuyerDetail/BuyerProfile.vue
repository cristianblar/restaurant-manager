<template>
  <section>
    <v-card
      class="mx-auto card-container"
      elevation="6"
      shaped
    >
      <v-card-title class="card-title">{{ name }}</v-card-title>
      <v-card-subtitle class="card-subtitle">{{ buyerId }}</v-card-subtitle>
      <v-card-text class="card-text"><strong>Age: </strong>{{ age }}</v-card-text>
    </v-card>
  </section>
</template>

<script>
export default {
  name: 'BuyerProfile',
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
    }
  }
}
</script>

<style lang="sass" scoped>

.card-container
  margin-bottom: 36px

.card-title
  color: #190862
  font-size: 2rem
  font-weight: 900
  margin-bottom: 12px

.card-subtitle
  font-size: 1.6rem
  font-weight: 500
  margin-bottom: 24px

.card-text
  font-size: 1.8rem
  font-weight: 300
  padding-bottom: 24px

</style>
