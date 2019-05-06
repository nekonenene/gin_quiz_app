# Gin quiz app

Under development...（出来上がらないかも……むずかしい……）

## How to develop

Install [Go](https://golang.org) and [Docker](https://docs.docker.com/install).  
Then, init.

```sh
export GO111MODULE=on
make init
```

Fill `.env` .

Start web server and database.

```sh
docker-compose up
```

Then, Go => http://localhost:8013

### Build JavaScript

```sh
make js_watch
```


## Using...

* [Gin](https://github.com/gin-gonic/gin)
* [GORM](https://github.com/jinzhu/gorm)
* [Air](https://github.com/cosmtrek/air)
* [sqldef](https://github.com/k0kubun/sqldef)
* [webpack](https://webpack.js.org)
* [TypeScript](https://www.typescriptlang.org)
* [React](https://reactjs.org)
* [Materialize.css](https://materializecss.com)


## ディレクトリ構成

* assets: HTML, CSS, JS が入る。フロント側
* assets_src: JS の元となる TypeScript ファイルなどがある。webpack を通して assets/js に吐き出される
* common: 使い勝手がよく、他との依存が薄いもの。 response.go を引き離したい……
* model: DB のテーブルの元となる
* registry: 設定などを維持しておく倉庫。必ず最初に `Init()` する
* repository: データの操作をおこなう。repository 同士で循環 import になりそうなら、repository を register に持たせ DI 的になるよう設計を再考すること
* router: 受け付ける path とそれを受けての操作を記述。DB は repository のメソッドを通して操作する

router -> repository -> model の関係性は崩さないよう、気を付けて実装する


## Thanks!

* [gothinkster/golang-gin-realworld-example-app](https://github.com/gothinkster/golang-gin-realworld-example-app)
* [voyagegroup/gin-boilerplate](https://github.com/voyagegroup/gin-boilerplate)
* [How to do Google sign-in with Go](https://skarlso.github.io/2016/06/12/google-signin-with-go)
