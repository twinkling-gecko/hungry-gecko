<template>
  <b-container>
    <h1>商品登録画面</h1>

    <b-container>
      <ItemForm :form="form" :on-submit="onSubmit" :on-reset="onReset" />
    </b-container>
  </b-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class ItemList extends Vue {
  // DATAの定義
  form = {
    name: '',
    summary: '',
    uri: '',
  }

  // submit時のアクション
  onSubmit(event: Event) {
    event.preventDefault()

    this.$axios
      .post(this.$axios.defaults.baseURL + 'items', {
        name: this.form.name,
        summary: this.form.summary,
        uri: this.form.uri,
      })
      .then(() => {
        alert('登録完了しました。')
        this.onReset(event)
      })
      .catch(() => {
        alert('登録失敗')
      })
  }

  onReset(event: Event) {
    event.preventDefault()

    this.form.name = ''
    this.form.summary = ''
    this.form.uri = ''
  }
}
</script>
