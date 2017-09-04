package main

import (
	"fmt"
	"github.com/tetiteti/helloworld/sub"
)
type Vertex struct {
	X, Y int
}
func main() {
	// 標準出力
	fmt.Println("ハローワールド")

	// 配列
	p := []int{2, 3, 5, 7, 11, 13}
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	// 連想配列
	var maps = map[string]int{
		"Bell Labs": 1,
		"Google": 2 ,
	}
	fmt.Println(maps)

	var m = map[string]Vertex{
		"aiueo": {
			X: 11,
			Y: 12,
		},
	}
	fmt.Println(m)


	// 外部ファンクションの呼び出し
	sub.SomeFunc()

}
