<template>
  <div class="A">
    <div class="registerscreen">
      <p>商品登録画面</p>
    </div>
    <div class="titles">
      <p>商品名</p>
      <input v-model="title" placeholder="商品名を入力してね" />
    </div>
    <div class="overview">
      <p>商品概要</p>
      <input v-model="overview" placeholder="商品概要を入力" />
    </div>
    <div class="link">
      <p>商品リンク</p>
      <input v-model="link" placeholder="商品のリンクを貼り付けるところ" />
    </div>
    <button @click="onClickButton">登録</button>
  </div>
</template>
<!--typeScript-->

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import axios from 'axios'

@Component
export default class List extends Vue {
  // DATAの定義
  title = ''
  overview = ''
  link = ''

  // クリック時のアクション
  onClickButton() {
    if (this.checkSendItem()) {
      const res = axios
        .post('http://localhost/api/v1/items', {
          name: this.title,
          summary: this.overview,
          uri: this.link,
        })
        .then(() => {
          alert('登録完了しました。')
          this.deleteForm()
        })
        .catch((error) => {
          alert('登録失敗')
          console.log(error)
        })
      return {
        data: res.data,
      }
    } else {
      alert('登録情報を入力してください')
    }
  }

  checkSendItem() {
    if (this.title && this.overview && this.link) {
      return true
    } else {
      return false
    }
  }

  deleteForm() {
    this.title = ''
    this.overview = ''
    this.link = ''
  }
}
</script>
<!--CSS styleだけだとERRORが起きる-->
