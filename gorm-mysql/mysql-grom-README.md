### 下準備

- mysqlをdocker-composeで

### 接続

### 


# 参考
https://midorigame-jo.com/docker-mysql/

###  docker-compose mysql

https://qiita.com/ysd_marrrr/items/e8a50c43cff87951385c
※volumeの解説あり
- volume指定しないとデータが保存されないよね


```
// アクセス
docker-compose exec mysql bash

// 
CREATE TABLE fruit (
    id INT NOT NULL PRIMARY KEY,
    apple VARCHAR(255) NOT NULL,
    price INT NOT NULL
);

insert into fruit (id, apple, price) values(1, "hoge", 1100);
```

- docker でvolumeを設定
```
// volumeの確認
docker volume ls

local     gorm-mysql_db-store
```


### gorm

https://gorm.io/ja_JP/docs/index.html

### クエリ検証


- 構造体

大文字小文字で、検索のヒットするしないが出る

```
type Fruit struct {
	ID    int    `gorm:"primaryKey"`
	Apple string
	Price int	
}
```

- where で簡単なクエリ

```
	var fruit Fruit
	db.Where("price = ?", 1100).First(&fruit)

```

- テーブルの連携 TableName

```
type Fruit struct {
	ID    int    `gorm:"primaryKey"`
	Apple string
	Price int	
}

func (Fruit) TableName() string {
	return "fruit"
}
```

- レコード作成



```
type User struct {
  Id int
  Name string
}

func  (User) TableName() string {
    return "user"
}

user := User { Id: 1, Name: "name"}

// ここは値渡せば良いだけだもんな
db.Create(user)
```

- 一括作成

※トランザクションが開始されているらしい
```
var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
db.Create(&users)


// バッチ処理

var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}

// batch size 100
db.CreateInBatches(users, 100)
```


- 作成時のhook
・before save
・after save

構造体を引数にBeforeCreateとかすれば良いだけ

```
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}
```

### gorm Model

- イメージしてる db.Model(&use)ではない

https://gorm.io/ja_JP/docs/models.html#gorm-Model


- イメージしてたもの

構造体にTableNameをつけない場合はこれか

```
type User struct {
    ID   uint
    Name string
    Age  int
}

db.Model(&User{}).Where("id = ?", 1).Update("Name", "New Name")

```

### アソシエーション

https://gorm.io/ja_JP/docs/associations.html#tags


- リレーションの種類
  - 1対1
  - 1対多
  - 多対1
  - 多対多

- タグを使用
  - foreignKey
  - references
  - gorm:"many2many:user_groups;"

- dbレベルでの関連性を表現する制約
  - REFERENCES => これだけ使う場合整合性チェックはしない
  - FOREIGN KEY

- belongs to

https://gorm.io/ja_JP/docs/belongs_to.html

- has many

https://gorm.io/ja_JP/docs/has_many.html

### Context

https://gorm.io/ja_JP/docs/context.html

```
// タイムアウトさせる
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

db.WithContext(ctx).Find(&users)
```

### 更新・削除・トランザクション

https://gorm.io/ja_JP/docs/query.html

- エラーハンドリング
- トランザクションなど
- migration
  - alter user 


https://gorm.io/ja_JP/docs/transactions.html

```go
// 一般的
// Error関数
db.Transaction(func(tx *gorm.DB) error {
  // トランザクション内でのデータベース処理を行う(ここでは `db` ではなく `tx` を利用する)
  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
    // 何らかのエラーを返却するとロールバックされる
    return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
    return err
  }

  // nilが返却されるとトランザクション内の全処理がコミットされる
  return nil
})

// 手動
tx := db.Begin() // トランザクション開始
defer func() {
    if r := recover(); r != nil {
        tx.Rollback() // パニックが発生した場合、ロールバック
        panic(r) // 再度パニック
    }
}()

if err := tx.First(&user, 1).Error; err != nil {
    tx.Rollback() // エラー時にロールバック
    return err
}

if err := tx.Model(&user).Update("role", "admin").Error; err != nil {
    tx.Rollback() // エラー時にロールバック
    return err
}

if err := tx.Commit().Error; err != nil { // トランザクションをコミット
    return err
}
```

### エラーハンドリング

https://gorm.io/ja_JP/docs/error_handling.html


```
if err := db.Where("name = ?", "jinzhu").First(&user).Error; err != nil {
  // Handle error...
}


// エラー
err := db.First(&user, 100).Error
if errors.Is(err, gorm.ErrRecordNotFound) {
  // Handle record not found error...
}
```

### 制約

- DB作成時、automigration時
https://gorm.io/ja_JP/docs/constraints.html


### インデックス

- db作成、automigration時に
https://gorm.io/ja_JP/docs/indexes.html