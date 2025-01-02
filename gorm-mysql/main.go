package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id int
	Name string
}

func (User) TableName() string {
	return "user"
}

type Fruit struct {
	ID    int    `gorm:"primaryKey"`
	Apple string
	Price int
	UserId int
	User User `gorm:"foreignKey:UserId"`
}

func (Fruit) TableName() string {
	return "fruit"
}


type FruitWithUser struct {
	ID int
	Apple   string
	Price   int
	UserId  int
	UserName string
}


func main() {

  dsn := "user:password@tcp(localhost:3306)/test-db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("errr")
	}

	var fruit Fruit
	// db.Where("price = ?", 1100).First(&fruit)
	db.Preload("User").Where("price = ?", 300).First(&fruit)

	// Updateは単一カラム
	// 複数からむはUpadates は構造体
	// db.Model(&fruit).Updates(Fruit{ Apple: "hoge", Price: 300})

	fmt.Printf("fruit id %d fruit apple %s user name %s", fruit.ID, fruit.Apple, fruit.User.Name)


	// サブクエリ利用

	// db.Table("users as u").Where("name = ?", "jinzhu").Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))


	// var result FruitWithUser
  //   db.Table("fruit").
  //   Select("fruit.id, fruit.apple, fruit.price, fruit.user_id, user.name as user_name").
  //   Joins("JOIN user ON user.id = fruit.user_id").
  //   Where("fruit.price = ?", 1100).
  //   Scan(&result)

	// fmt.Printf("Fruit ID: %d, Apple: %s, Price: %d, User ID: %d, User Name: %s\n",
	// 	result.ID, result.Apple, result.Price, result.UserId, result.UserName)


	// db.Joins("User").Where("fruit.price = ?", 1100).First(&fruit)

	// db.Model(&Fruit{}).
	//   Joins("INNER JOIN user ON user.id = fruit.user_id").
	// 	Where("fruit.price = ?", 1100)

	// fmt.Printf("Fruit: %+v, User: %+v\n", fruit, fruit.User)


	// var results []FruitWithUser
	// db.Model(&Fruit{}).
	// 	Select("fruit.id as fruit_id, fruit.apple, fruit.price, fruit.user_id, user.id as user_id, user.name as user_name").
	// 	Joins("JOIN user ON user.id = fruit.user_id").
	// 	Where("fruit.price = ?", 1100).
	// 	Scan(&results)

	// // 結果の出力
	// for _, result := range results {
	// 	fmt.Printf("Fruit: ID=%d, Apple=%s, Price=%d, User: ID=%d, Name=%s\n",
	// 		result.FruitID, result.Apple, result.Price, result.UserID, result.UserName)
	// }

}