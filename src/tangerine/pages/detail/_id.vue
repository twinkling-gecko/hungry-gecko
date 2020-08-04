<template>
  <div>
    <ItemDetail :item="data" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import ItemDetail from '@/components/ItemDetail.vue'

@Component({
  components: {
    ItemDetail
  }
})
export default class Detail extends Vue {
  async asyncData(context: any) {
    const id = context.route.params.id
    const res = await context.$axios.get(
      context.$axios.defaults.baseURL + 'items/' + id
    )
    return {
      data: res.data
    }
  }
}
</script>
