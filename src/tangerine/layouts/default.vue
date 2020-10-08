<template>
  <div>
    <Header />
    <div class="content">
      <b-row>
        <b-col cols="3" class="pl-2 pr-0">
          <!-- TODO: Sidebarのコンポーネント切り出し -->
          <b-container>
            <b-card no-body class="my-2">
              <b-list-group flush>
                <b-list-group-item>
                  <n-link to="/items/">Item一覧</n-link>
                </b-list-group-item>
                <b-list-group-item>
                  <n-link to="/items/new">Item登録</n-link>
                </b-list-group-item>
              </b-list-group>
            </b-card>
          </b-container>
        </b-col>
        <b-col cols="9" class="pl-0 pr-2">
          <Nuxt />
        </b-col>
      </b-row>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Header from '@/components/GeneralHeader.vue'

@Component({
  components: {
    Header,
  },
})
export default class Default extends Vue {}
</script>

<style lang="scss" scoped>
.content {
  margin-top: 80px;
}
</style>
