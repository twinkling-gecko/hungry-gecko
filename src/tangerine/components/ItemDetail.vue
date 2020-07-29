<template>
  <div class="container">
    <h2>{{item.name}}</h2>
    <p>{{item.summary}}</p>
    <BLink :href="item.uri">{{item.uri}}</BLink>
    <p>更新日：{{item.updated_at}}</p>
  </div>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { BLink } from 'bootstrap-vue'

interface Item {
  created_at: string
  id: number
  name: string
  summary: string
  updated_at: string
  uri: string
}

@Component({
  components: {
    BLink,
  },
})
export default class ItemDetail extends Vue {
  @Prop({ type: Object as PropType<Item>, default: {} })
  item!: Item
}
</script>

<style lang="scss" scoped>
.container {
  padding: 8px;
}
</style>
