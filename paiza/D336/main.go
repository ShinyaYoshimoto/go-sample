package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 標準入力を取得する
	sc := bufio.NewScanner(os.Stdin)

	// 1行目（前に並んでいる人数 n ）
	sc.Scan()
	n, e1 := strconv.Atoi(sc.Text())

	// 2行目（後ろに並んでいる人数 m ）
	sc.Scan()
	m, e2 := strconv.Atoi(sc.Text())

	if e1 != nil || e2 != nil {
		return
	}

	total, err := handle(n, m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(total)
	}
}

func handle(front int, behind int) (int, error) {
	if !validate(front, behind) {
		return 0, errors.New("invalid input")
	}

	total := front + behind + 1
	return total, nil
}

func validate(n int, m int) bool {
	return n >= 1 && n <= 10 && m >= 1 && m <= 10
}
