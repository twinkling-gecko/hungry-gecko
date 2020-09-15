<template>
  <b-container>
    <h1>商品編集画面</h1>

    <b-container>
      <ItemForm :form="form" :on-submit="onSubmit" :on-reset="onReset" />
    </b-container>
  </b-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class ItemUpdate extends Vue {
  id = ''
  form = {
    name: '',
    summary: '',
    uri: '',
  }

  created() {
    this.id = this.$route.params.id
  }

  onSubmit(event: Event) {
    event.preventDefault()

    this.$axios
      .patch(this.$axios.defaults.baseURL + 'items/' + this.id, {
        name: this.form.name,
        summary: this.form.summary,
        uri: this.form.uri,
      })
      .then(() => {
        alert('編集完了しました。')
        this.onReset(event)
      })
      .catch(() => {
        alert('編集失敗')
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
