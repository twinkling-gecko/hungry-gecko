<template>
  <div>
    <h2>{{ item.name }}</h2>
    <p>{{ item.summary }}</p>
    <b-link :href="item.uri">{{ item.uri }}</b-link>
    <p>更新日：{{ item.updated_at }}</p>
  </div>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { Item } from '@/types/index'

@Component
export default class ItemDetail extends Vue {
  @Prop({
    type: Object as PropType<Item>,
    default: (): Item => ({
      id: 0,
      name: '',
      summary: '',
      uri: '',
      created_at: Date.now(),
      updated_at: Date.now(),
    }),
  })
  item!: Item
}
</script>
