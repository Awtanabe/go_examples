package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// ユーザー情報の構造体
type User struct {
	ID    int
	Email string
}

// ユーザーリスト（仮データ）
var users = []User{
	{ID: 1, Email: "user1@example.com"},
	{ID: 2, Email: "user2@example.com"},
	{ID: 3, Email: "user3@example.com"},
	{ID: 4, Email: "user4@example.com"},
}

// 疑似的なメール送信関数
func sendEmail(user User) {
	// ユーザーIDが3の場合にパニックを発生させる
	if user.ID == 3 {
		panic(errors.New("failed to send email: simulated panic"))
	}
	fmt.Printf("メール送信成功: %s\n", user.Email)
}

// 1人分の処理を行う関数（panicをrecoverする）
func processUser(user User) {
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("ユーザーID %d へのメール送信中にエラー発生: %v\n", user.ID, rec)
		}
	}()

	// メール送信処理
	sendEmail(user)
}

func main() {
	log.Println("一括メール送信処理を開始します...")
	startTime := time.Now()

	// ユーザーごとにメール送信
	for _, user := range users {
		processUser(user) // ユーザーごとにpanicをrecover
	}

	log.Printf("一括メール送信処理が完了しました。経過時間: %s\n", time.Since(startTime))
}
