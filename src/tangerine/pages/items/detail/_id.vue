<template>
  <ItemDetail v-if="!fetchFail" :item="item" />
  <div v-else>通信に失敗しました</div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import ItemDetail from '@/components/ItemDetail.vue'

@Component({
  components: {
    ItemDetail,
  },
  asyncData(context: any) {
    const id = context.route.params.id
    return context.$axios
      .get(context.$axios.defaults.baseURL + 'items/' + id)
      .then((res: any) => {
        const data = res.data
        return {
          item: {
            id: data.id,
            name: data.name,
            summary: data.summary,
            uri: data.uri,
            created_at: data.created_at,
            updated_at: data.updated_at,
          },
        }
      })
      .catch(() => {
        return {
          fetchFail: true,
        }
      })
  },
})
export default class Detail extends Vue {
  item = {
    id: 0,
    name: '',
    summary: '',
    uri: '',
    created_at: Date.now(),
    updated_at: Date.now(),
  }

  fetchFail = false
}
</script>
