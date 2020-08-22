<template>
  <div class="Con">
    <div v-if="!fetchFail">
      <div v-for="item in items" :key="item.id" class="Card">
        <b-card :title="item.name">
          <b-card-text>
            {{ item.summary }}
          </b-card-text>
          <a :href="item.uri" class="card-link">商品リンク</a>
        </b-card>
      </div>
    </div>
    <div v-else>通信に失敗しました</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component({
  asyncData(context: any) {
    return context.$axios
      .get(context.$axios.defaults.baseURL + 'items')
      .then((res: any) => {
        const data = res.data
        return {
          items: data.items,
        }
      })
      .catch(() => {
        return {
          fetchFail: true,
        }
      })
  },
})
export default class List extends Vue {
  item = {
    id: 0,
    name: '',
    summary: '',
    uri: '',
  }

  fetchFail = false
}
</script>
