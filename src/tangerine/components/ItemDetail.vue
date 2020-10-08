<template>
  <div class="my-2">
    <h2>{{ item.name }}</h2>
    <p class="small">更新日：{{ item.updated_at }}</p>
    <p>{{ item.summary }}</p>
    <b-link :href="item.uri">{{ item.uri }}</b-link>
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
