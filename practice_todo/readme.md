### リソース

- ⭐️チュートリアル

https://go-tour-jp.appspot.com/concurrency/10

- ⭐️go テスト駆動開発

https://andmorefine.gitbook.io/learn-go-with-tests

- ⭐️go ナイスな記事 web開発の基本を

https://zenn.dev/tara_is_ok/scraps/fcdd5bf4e44b75


- ⭐️アプリケーションを作成(気になる => ベースの知識が学べそうである)
https://go.dev/doc/articles/wiki/

https://go.dev/doc/articles/wiki/final.go
全部のコード
バリデーションはパスのgrep

=> /Users/akifumi_watanabe/work/writing_web_application/gowiki

- ⭐️マインドマップ
初めてのgo

https://www.mindmeister.com/app/map/3547997529?source=template

- effective go

https://go.dev/doc/effective_go

- 言語仕様について

https://go.dev/ref/spec#Introduction

### todoのtodo

- 後からのカラム変更migrate

### todo


- サーバー作成
- todo モデル
- docker-compose
- バリデーーション
- ユーザー作成
- todoと紐付ける
- 認証
- テストコード
- メール
- ジョブ、バッチの実行
- デプロイ
- ci cd



### dockerfile

- 参考
https://zenn.dev/ryouta26/articles/1aff585faa5084


- docker && docker-compose
https://zenn.dev/ryouta26/articles/1aff585faa5084

- docker-compose echo

https://zenn.dev/osushi02/articles/21e14cbcf71008

### air

- https://github.com/air-verse/air

```
// install
go install github.com/air-verse/air@latest
// init
air init
```


## docker gorm mysql

https://qiita.com/nao-United92/items/01a3b02b41b0c26fb56d


## gorm 参考になる

https://zenn.dev/ring_belle/articles/go-gorm-docker

## go api

https://zenn.dev/ohke/articles/go-echo-how-to-use-for-backend-api


## いい記事 go

https://zenn.dev/tara_is_ok/scraps/fcdd5bf4e44b75

- [x]プロジェクト作成
  -　ディレクトリ作成
  - go mod
- [x]mysql起動
  - postgresだけどmysql
    - docker-compose
- [x] 環境変数
- [x] models
  - todoとuserの定義
- [x] db
  - db gormを起動して、msyqlに接続
- [x] マイグレーション
  - migrate/migrate.go
    - コンテナの中に入って
- [x] ユーザーAPI作成
  - usecase
    - user_usecase.go
  - repogitory
    - interface
    - user_repository.go
  - controller
  - router
- [x] jwt install
- [x] TODO API
  - todo usecase
  - todo repogitory
- [x] validation
  - todo・userに適用
- [x] middleware
  - cors・csrfに追加
- [] メール
  - mailエンドポイントを作成してメールを送信
    - 作成中
  - docker-compose でdev mailを
    - https://github.com/maildev/maildev
  - メール実装
  - net/mail
    - https://pkg.go.dev/net/mail
    - https://github.com/jordan-wright/email

### データを消したい時

- ローカル環境ではこれおすすめ
docker-compose down --volumes

## jwt

https://github.com/golang-jwt/jwt

- echo jtw

https://echo.labstack.com/docs/cookbook/jwt

### テスト編集

### getUserIDFromJWTがうまくできなかった

- jwtのバージョンの競合が問題だったs

### 1対多のデータ取得

https://qiita.com/sedori/items/adc9d09e58c83823c6be

```
func GetOneItems(id int) (data Item, err error) {
    err = Db.Debug().Preload("Company").Preload("Player").Preload("Tags").Preload("Images").First(&data, id).Error

    return
}

```


### go tour


### go routine

#### 基本造作

### チャネル

- チャネル定義
  - chan := make(chan int)
    - chan := make(chan int, 10)
      - これはバッファあり
        - これの場合はデットロックにならない
- 送信・受信の書き方
  - 送信(値を入れる)
    - ch <- 1
  - 受信(出力)
    - <- ch
- 受信と送信の順番はどっちでも良い

- 基本的な基本的なチャネルの利用

```
// 送信が先の場合

func main() {
	c := make(chan int)

	go func() {
    // 送信
		c <- 1
		fmt.Println("送信する")

	}()

  // 受信
	fmt.Println("受信", <-c)

}

// 受信が先の場合
func main() {
	c := make(chan int)

	go func() {
		fmt.Println("受信", <-c)

	}()

	c <- 1
	fmt.Println("送信する")

}
⭐️いずれかは非同期(go routine)でいずれかは同期
⭐️デットロックは

// バッファありチャネル
⭐️順番は送信、受信(受信先にやるとデットロック)
func main() {
	c := make(chan int, 1)

	c <- 1
	fmt.Println("送信する")
	fmt.Println("受信", <-c)

}
```

- select(tour)
  - https://go-tour-jp.appspot.com/concurrency/5
    - 送信・受信が詰まっている
    - for文や無限ループの理解も必要
    - 受信が非同期で、送信が同期的だった

- timeはチャネル型を返すらしい
  - https://go-tour-jp.appspot.com/concurrency/6
  - time := time.Tick(100 * time.Millisecond)
    - ⭐️読み取り専用のチャネルを返してる

