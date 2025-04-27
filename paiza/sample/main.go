package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 標準入力を取得する
	sc := bufio.NewScanner(os.Stdin)

	arr := []string{}

	// 1行目
	sc.Scan()

	arr = append(arr, sc.Text())

	// 2行目以降
	for len(sc.Text()) > 0 {
		sc.Scan()

		if len(sc.Text()) > 0 {
			arr = append(arr, sc.Text())
		} else {
			break
		}
	}

	front, e1 := strconv.Atoi(arr[0])
	behind, e2 := strconv.Atoi(arr[1])

	if e1 != nil || e2 != nil {
		return
	}

	logic(front, behind)
}

func logic(front int, behind int) {
	fmt.Println(front, behind)
}
