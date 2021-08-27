<template>
  <GridLoader
    class="loader"
    color="#FFBC00"
    :size="60"
  />
</template>

<script>
import { GridLoader } from '@saeris/vue-spinners'

export default {
  name: 'Loader',
  components: { GridLoader }
}
</script>

<style lang="sass" scoped>

.loader
  margin: 60px auto 0

</style>
