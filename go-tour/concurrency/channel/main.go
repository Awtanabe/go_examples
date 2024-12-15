package main

// 目次
// 送信が先の場合
// 受信が先の場合
// バッファありのチャネルの場合(デットロック回避のために)

// 送信が先の場合
// func main() {
// 	c := make(chan int)

// 	go func() {
//     // 送信
// 		c <- 1
// 		fmt.Println("送信する")
// 	}()
//   // 受信
// 	fmt.Println("受信", <-c)
// }

// 受信が先の場合
// func main() {
// 	c := make(chan int)
// 	go func() {
// 		fmt.Println("受信", <-c)
// 	}()
// 	c <- 1
// 	fmt.Println("送信する")
// }

// ⭐️いずれかは非同期(go routine)でいずれかは同期
// ⭐️デットロックは

// バッファありチャネル
// ⭐️順番は送信、受信(受信先にやるとデットロック)
// func main() {
// 	c := make(chan int, 1)
// 	c <- 1
// 	fmt.Println("送信する")
// 	fmt.Println("受信", <-c)
// }