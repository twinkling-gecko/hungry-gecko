<template>
  <div class="Con">
    <div v-for="item in data.items" :key="item.id" class="Card">
      <b-card :title="item.name">
        <b-card-text>
          {{ item.summary }}
        </b-card-text>
        <a :href="item.uri" class="card-link">商品リンク</a>
      </b-card>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { BLink, BCard, BCardText } from 'bootstrap-vue'

@Component({
  components: {
    BLink,
    BCard,
    BCardText,
  },
})
export default class List extends Vue {
  async asyncData(context: any) {
    const res = await context.$axios.get(
      'http://localhost/api/v1/items'
    )
    return {
      data: res.data,
    }
  }
}
console.log()
</script>
