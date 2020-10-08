<template>
  <b-navbar variant="light" fixed="top" class="nav">
    <b-navbar-brand to="/">
      <Gecko class="gecko-icon" />
      Hungry Gecko
    </b-navbar-brand>
  </b-navbar>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Gecko from '@/assets/svg/gecko.svg'

@Component({
  components: {
    Gecko,
  },
})
export default class GeneralHeader extends Vue {}
</script>

<style lang="scss" scoped>
.nav {
  height: 60px;
}
.gecko-icon {
  height: 40px;
  width: auto;
  transform: rotate(-45deg);
}
</style>
