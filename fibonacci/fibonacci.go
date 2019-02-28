// n 番目のフィボナッチ数列を O(n) 時間で計算する動的計画アルゴリズムを設計せよ

package main

import "fmt"

func fibonacci(n int, m []int) int {
	if m[n] > 0 {
		return m[n]
	}
	res := fibonacci(n-1, m) + fibonacci(n-2, m)
	m[n] = res
	return res
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	// フィボナッチ数の結果を格納する配列を負の数で初期化する
	m := make([]int, n+1)
	for i := 0; i <= n; i++ {
		m[i] = -10000
	}
	m[1], m[2] = 1, 1

	fmt.Println(fibonacci(n, m))
}
