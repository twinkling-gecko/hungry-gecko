<template>
  <div v-if="!fetchFail">
    <b-card v-for="item in items" :key="item.id" class="Card">
      <n-link :to="detailUri(item.id)" class="detail-link">
        <b-card-title>{{ item.name }}</b-card-title>
        <b-card-text>{{ item.summary }}</b-card-text>
      </n-link>
      <b-link :href="item.uri" class="card-link">商品リンク</b-link>
    </b-card>
  </div>
  <div v-else>通信に失敗しました</div>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { Item } from '@/types/index'

@Component({
  asyncData(context: Context) {
    const { $axios } = context
    return $axios
      .$get($axios.defaults.baseURL + 'items')
      .then((res: any) => {
        const items: Item[] = res.items
        return {
          items,
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
  fetchFail = false

  get detailUri() {
    return (id: any) => {
      return '/items/detail/' + id
    }
  }
}
</script>

<style lang="scss" scoped>
.detail-link {
  color: black;
  &:hover {
    color: #007bff;
    text-decoration: none;
  }
}
</style>
