<template>
  <b-navbar variant="light" fixed="top" class="nav">
    <b-navbar-brand to="/">
      <Gecko class="gecko-icon" />
      Hungry Gecko
    </b-navbar-brand>
    <b-navbar-nav class="ml-auto">
      <nuxt-link to="/items/new">
        <b-button size="sm" class="my-2 my-sm-0"> NewItems </b-button>
      </nuxt-link>
      <nuxt-link to="/items/">
        <b-button size="sm" class="my-2 my-sm-0"> アイテム一覧 </b-button>
      </nuxt-link>
    </b-navbar-nav>
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
