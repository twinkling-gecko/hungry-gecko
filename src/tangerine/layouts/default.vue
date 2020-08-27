<template>
  <div>
    <Header />
    <div class="content">
      <Nuxt />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Header from '@/components/GeneralHeader.vue'

@Component({
  components: {
    Header,
  },
})
export default class Default extends Vue {}
</script>

<style lang="scss" scoped>
.content {
  margin-top: 60px;
}
</style>
