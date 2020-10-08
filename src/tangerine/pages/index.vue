<template>
  <b-container>
    <div class="text-center">
      <!-- TODO: ブランドロゴの作成 -->
      <Logo />
      <h1 class="display-1">Hungry Gecko</h1>
    </div>
    <div class="text-center">
      <b-button class="mr-2" variant="outline-primary" size="lg" to="/items">
        ItemList
      </b-button>
      <b-button
        class="ml-2"
        variant="outline-secondary"
        size="lg"
        to="/items/new"
      >
        NewItem
      </b-button>
    </div>
  </b-container>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  layout: 'top',
})
</script>
