<template>
  <BNavbar variant="light">
    <BNavbarBrand>amazon review</BNavbarBrand>
  </BNavbar>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { BNavbar, BNavbarBrand } from 'bootstrap-vue'

@Component({
  components: {
    BNavbar,
    BNavbarBrand,
  },
})
export default class GeneralHeader extends Vue {}
</script>
