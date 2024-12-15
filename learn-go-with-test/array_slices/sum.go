package main

// TODO: 配列の大きさが固定されるのは厄介
// func Sum(numbers [5]int) int {
// 	sum := 0
// 	// for i :=0; i < 5; i++ {
// 	// 	sum += numbers[i]
// 	// }
// 	// こっちの方がインデックス指定しないで良いから綺麗か
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	return sum
// }


func Sum(numbers []int) int {
	sum := 0
	// for i :=0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	// こっちの方がインデックス指定しないで良いから綺麗か
	for _, number := range numbers {
		sum += number
	}
	return sum
}


func SumAll(numbersToSum ...[]int) []int {
	// numbersToSum 可変長引数のスライスの数を数える
	lengthOfNumbers := len(numbersToSum)
	// スライスを lengthOfNumbers個作成
	// このmakeはスライスを作っている。lengthOfNumbersは要素数 
	sums := make([]int, lengthOfNumbers)

	// それぞれのスライスをループ
	for i, numbers := range numbersToSum {
		  // 箱にそれぞれのスライスの合計を入れる
			sums[i] = Sum(numbers)
	}

	return sums
}

// これだとテスト失敗する
// func SumAllTails(numbersToSum ...[]int) []int {
// 	// スライス
// 	// 要素数を決めていない
// 	var sums []int

// 	for _, numbers := range numbersToSum {
// 		// numbers[1:] は1から最後まで
// 		tail := numbers[1:]
// 		sums = append(sums, Sum(tail))
// 	}
// 	return sums
// }

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
			if len(numbers) == 0 {
					sums = append(sums, 0)
			} else {
					tail := numbers[1:]
					sums = append(sums, Sum(tail))
			}
	}

	return sums
}