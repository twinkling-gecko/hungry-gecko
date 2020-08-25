<template>
  <ItemDetail v-if="isFetchSuccessed" :item="item" />
  <div v-else>通信に失敗しました</div>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import ItemDetail from '@/components/ItemDetail.vue'
import { Item } from '@/types/index'

@Component({
  components: {
    ItemDetail,
  },
  asyncData(context: Context) {
    const { route, $axios } = context
    const id = route.params.id
    return $axios
      .$get($axios.defaults.baseURL + 'items/' + id)
      .then((item: Item) => {
        return {
          item,
        }
      })
      .catch(() => {
        return {
          isFetchSuccessed: false,
        }
      })
  },
})
export default class Detail extends Vue {
  isFetchSuccessed = true
}
</script>
