<template>
  <b-form @submit="onSubmit" @reset="onReset">
    <b-form-group label="商品名" label-for="item-name">
      <b-form-input
        id="item-name"
        v-model="form.name"
        type="text"
        placeholder="商品名を入力してね"
        required
      />
    </b-form-group>
    <b-form-group label="商品概要" label-for="item-summary">
      <b-form-input
        id="item-summary"
        v-model="form.summary"
        type="text"
        placeholder="商品概要を入力してね"
        required
      />
    </b-form-group>
    <b-form-group label="商品リンク" label-for="item-uri">
      <b-form-input
        id="item-uri"
        v-model="form.uri"
        type="text"
        placeholder="商品のリンクを入力してね"
        required
      />
    </b-form-group>
    <b-button type="submit">登録</b-button>
  </b-form>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'

@Component
export default class ItemForm extends Vue {
  @Prop({
    default: {
      name: '',
      summary: '',
      uri: '',
    },
  })
  form!: { name: string; summary: string; uri: string }

  @Prop({
    default: (event: Event) => {
      event.preventDefault()
      return false
    },
  })
  onSubmit!: Function

  @Prop({
    default: (event: Event) => {
      event.preventDefault()
      return false
    },
  })
  onReset!: Function
}
</script>
