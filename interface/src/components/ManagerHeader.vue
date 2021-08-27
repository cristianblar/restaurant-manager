<template>
  <header>
    <div class="d-flex align-center">
      <router-link to="/">
        <h1>Restaurant manager 🍔</h1>
      </router-link>
    </div>

    <nav class="search-box">
      <v-text-field
        :disabled="!syncedData"
        v-model="inputId"
        class="search-input"
        dark
        label="Search buyer by id"
      />
      <v-btn :disabled="!syncedData" @click="handleSearch" icon color="white">
        <v-icon>mdi-magnify</v-icon>
      </v-btn>
    </nav>
  </header>
</template>

<script>
export default {
  name: 'ManagerHeader',
  props: {
    syncedData: {
      type: Boolean,
      default: true
    }
  },
  data: () => ({
    inputId: ''
  }),
  methods: {
    handleSearch () {
      this.$router.push(`/buyers/${this.inputId.toLowerCase()}`)
      this.inputId = ''
    }
  }
}
</script>

<style lang="sass" scoped>

header
  align-items: center
  background-color: #190862
  box-shadow: 0 8px 8px -1px rgba(0, 0, 0, 0.3)
  display: flex
  min-height: 100px
  justify-content: space-between
  margin-bottom: 36px
  padding: 12px 24px

a
  text-decoration: none

h1
  color: white
  font-size: 2.4rem
  font-weight: 500

  &:hover
    text-decoration: underline

.search-box
  align-items: center
  display: flex

</style>
