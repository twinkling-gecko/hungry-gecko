<template>
  <div>
    <p>商品登録画面</p>
    <p>商品名</p>
    <input v-model="title" placeholder="商品名を入力してね" />
    <p>商品概要</p>
    <input v-model="overview" placeholder="商品概要を入力" />
    <p>商品リンク</p>
    <input v-model="link" placeholder="商品のリンクを貼り付けるところ" />
    <button @click="onClickButton">登録</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class ItemList extends Vue {
  // DATAの定義
  title = ''
  overview = ''
  link = ''

  // クリック時のアクション
  onClickButton() {
    if (this.validateSendItem()) {
      const res = this.$axios
        .post(this.$axios.defaults.baseURL + 'items', {
          name: this.title,
          summary: this.overview,
          uri: this.link,
        })
        .then(() => {
          alert('登録完了しました。')
          this.clearItemForm()
        })
        .catch(() => {
          alert('登録失敗')
        })
      return {
        data: res.data,
      }
    } else {
      alert('登録情報を入力してください')
    }
  }

  validateSendItem() {
    if (this.title && this.overview && this.link) {
      return true
    } else {
      return false
    }
  }

  clearItemForm() {
    this.title = ''
    this.overview = ''
    this.link = ''
  }
}
</script>
