<template>
  <div class="main-container">
    <h2>{{ message }}</h2>
    <router-link class="back-link" :to="link">{{ linkText }}</router-link>
  </div>
</template>

<script>
export default {
  name: 'UpperErrorMessage',
  props: {
    message: {
      type: String,
      required: true
    },
    link: {
      type: String,
      required: true
    },
    linkText: {
      type: String,
      required: true
    }
  }
}
</script>

<style lang="sass" scoped>

.main-container
  & h2
    color: #FFBC00
    font-size: 3.6rem
    font-weight: 900
    margin-bottom: 60px
    text-align: center

  & .back-link
    color: #190862
    display: block
    font-size: 3rem
    font-weight: 500
    margin: 0 auto
    text-decoration: none
    width: fit-content

    &:hover
      text-decoration: underline

</style>
