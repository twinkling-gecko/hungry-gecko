<template>
  <b-container>
    <h2>{{ item.name }}</h2>
    <p>{{ item.summary }}</p>
    <b-link :href="item.uri">{{ item.uri }}</b-link>
    <p>更新日：{{ item.updated_at }}</p>
  </b-container>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Component, Vue, Prop } from 'nuxt-property-decorator'

interface Item {
  created_at: string
  id: number
  name: string
  summary: string
  updated_at: string
  uri: string
}

@Component
export default class ItemDetail extends Vue {
  @Prop({ type: Object as PropType<Item>, default: {} })
  item!: Item
}
</script>
