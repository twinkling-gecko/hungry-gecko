# hungry-gecko

```
.
├── .go
│   └── pkg             # go mod のパッケージキャッシュ
├── README.md
├── docker-compose.yml
└── src
    ├── enigma          # nginx の設定ファイル
    ├── macksnow        # go のサーバサイドプロジェクト
    └── tangerine       # nuxtのフロントエンドプロジェクト
```

# 環境構築

## required

- docker
- docker-compose

## github packages の設定

`read:package` の権限のある token でログインする（パスワードは打たない）

`$ docker login https://docker.pkg.github.com -u USERNAME`

## 開発用コンテナの立ち上げ

### tangerine の依存性

依存ライブラリを足した場合や、はじめて起動する場合などに必要。  
`/src/tangerine/node_modules` に保存される。

`$ docker-compose run tangerine yarn`

### macksnow の依存性

初回起動時に勝手に解決される。

手動でやりたい場合は

`$ docker-compose run macksnow go mod download`

`/.go/pkg` に保存される

### まとめて立ち上げ

`$ docker-compose up` （-d でバックグラウンド起動）

`http://localhost` でアプリに接続できる。  
`/` が tangerine、`/api`が macksnow

### アプリケーション単体で立ち上げ

`$ docker-compose run --service-ports tangerine`

or

`$ docker-compose run --service-ports macksnow`

tangerine は `http://localhost:3000`、
macksnow は `http://localhost:4000` で繋がる。

# 運用ルール

## git

git-flow をベースとする。

開発時は最新の develop から作業ブランチを切り、プルリクエストを出す。
master ブランチは常に production の状態となる。
**master, develop で直にコミットを行わない**。

1. `$ git switch develop`
1. `$ git pull`
1. `$ git switch -c feat/new-featuer`
1. `$ git push -u origin feat/new-featuer`

### ブランチ名

| 区分                               | 命名                     |
| ---------------------------------- | ------------------------ |
| 新機能、機能変更                   | feat/wakari-yasui-namae  |
| リリース済みのバグ修正             | fix/wakari-yasui-namae   |
| 機能追加を伴わないドキュメント追加 | doc/wakari-yasui-namae   |
| 機能追加を伴わないテスト追加       | test/wakari-yasui-namae  |
| その他                             | chore/wakari-yasui-namae |

### コミットメッセージ

| 区分                               | 命名                        |
| ---------------------------------- | --------------------------- |
| 新機能、機能変更                   | feat: wakari yasui message  |
| リリース済みのバグ修正             | fix: wakari yasui message   |
| 機能追加を伴わないドキュメント追加 | doc: wakari yasui message   |
| 機能追加を伴わないテスト追加       | test: wakari yasui message  |
| その他                             | chore: wakari yasui message |
