<template>
  <b-container>
    <h1>商品登録画面</h1>

    <b-container>
      <b-form @submit="onSubmit" @reset="onReset">
        <b-form-group label="商品名" label-for="item-name">
          <b-form-input
            v-model="form.title"
            id="item-name"
            type="text"
            placeholder="商品名を入力してね"
            required
          />
        </b-form-group>
        <b-form-group label="商品概要" label-for="item-summary">
          <b-form-input
            v-model="form.summary"
            id="item-summary"
            type="text"
            placeholder="商品概要を入力してね"
            required
          />
        </b-form-group>
        <b-form-group label="商品リンク" label-for="item-link">
          <b-form-input
            v-model="form.link"
            id="item-link"
            type="text"
            placeholder="商品のリンクを入力してね"
            required
          />
        </b-form-group>
        <b-button type="submit">登録</b-button>
      </b-form>
    </b-container>
  </b-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class ItemList extends Vue {
  // DATAの定義
  form = {
    title: '',
    summary: '',
    link: '',
  }

  // submit時のアクション
  onSubmit() {
    this.$axios
      .post(this.$axios.defaults.baseURL + 'items', {
        name: this.form.title,
        summary: this.form.summary,
        uri: this.form.link,
      })
      .then(() => {
        alert('登録完了しました。')
        this.onReset()
      })
      .catch(() => {
        alert('登録失敗')
      })
  }

  onReset() {
    this.form.title = ''
    this.form.summary = ''
    this.form.link = ''
  }
}
</script>
