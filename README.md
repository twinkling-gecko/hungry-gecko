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
- yarn

## pre-commitの設定

プロジェクトルートで

`$ yarn`

## github packages の設定

`read:package` の権限のある token でログインする。  
表示されるプロンプトには**パスワードではなく** token を入力する。

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
バックグラウンドで起動した場合は、`$ docker-compose logs -f (サービス名)` で任意のサービスのログにつなげる。

`http://localhost` でアプリに接続できる。  
`/` が tangerine、`/api`が macksnow

### アプリケーション単体で立ち上げ

`$ docker-compose run --service-ports tangerine`

or

`$ docker-compose run --service-ports macksnow`

tangerine は `http://localhost:3000`、
macksnow は `http://localhost:4000` で繋がる。

# DBのマイグレーション

ネイティブにインストールされたgoose(presslyによるfork版)を利用する。  
コンテナは利用せず、外からマイグレーションを行う。  
ネイティブにgoの動作する環境と、$GOBINにpathが通っている必要がある。

## gooseのインストール

`$ go get -u github.com/pressly/goose/cmd/goose`

## DBのマイグレーションを適用する

`$ cd src/macksnow/migrations/`  
`$ goose mysql "giant:leopard@/macksnow?parseTime=true" up`

## DBのロールバック

`$ cd src/macksnow/migrations/`  
`$ goose mysql "giant:leopard@/macksnow?parseTime=true" down`

## パラメータの省略

マイグレーション適用とロールバックの際に毎回ドライバと接続項目を打つのは面倒なので  
環境変数に入れておくとよい。

`$ export GOOSE_DRIVER=mysql`  
`$ export GOOSE_DBSTRING="giant:leopard@/macksnow?parseTime=true"`

これを入れておくと、マイグレーション適用とロールバックをそれぞれ

`$ goose up`  
`$ goose down`

で実行可能。

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

# heroku の検証環境について - prudent-gecko

- prudent: 用心深い、慎重な

## この環境について

| 項目         | 説明                      |
| ------------ | ------------------------- |
| バックエンド | heroku (stack: container) |
| DB           | mysql 8.0, 5MB まで       |

この環境は３つのサーバ（nginx, echo, nuxt）を強引に一つの Docker イメージに詰め込んでいるので、多少癖がある。

コンテナの起動時、echo, nuxt を起動した後に nginx のプロセスを起動する（`src/eclipse/start.sh` を参照）が、  
heroku の制約上 web アクセスを listen する nginx がまっさきに起動する必要があり裏のサーバの起動を待機しない。  
そのため起動直後には 502 エラーが発生する場合がある。アプリケーションの実装上の問題ではないのでリロードでカバー。

DB は事前に共有してる認証情報から直に接続できるので、各自適当なクライアントからメンテが可能。

## デプロイ方法

基本的には`prudent`へマージして push するだけで反映可能。  
ただし、このブランチは定期的に破棄して develop と一致させるためここで作業はしないこと。
（意図しないタイミングで作業が失われる事がある）

ただし DB 情報は基本的に失われない。5MB を超えてしまった場合は破棄する。

1. 反映したい変更を `prudent` ブランチにマージする
   - `git switch prudent` => `git merge other-branch`
2. github に push する
   - `git push`
3. push 後自動的に heroku でビルドとデプロイが始まる
   - heroku のダッシュボードからビルドの進捗を確認可能
   - Docker イメージのビルドから始まるため、そこそこ時間がかかる

## マイグレーション

マイグレーションは heroku のデプロイ時に自動で行われる。

prudent にマージした作業を取り消した場合などでマイグレーション取り消しを行いたい場合は、
**マイグレーションファイルを削除する前に**ローカルから prudent の db 宛に `goose down` を実行する。

