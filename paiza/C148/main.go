package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 標準入力を取得する
	sc := bufio.NewScanner(os.Stdin)

	battleList := []int{}

	// 1行目
	sc.Scan()

	info := strings.Split(sc.Text(), " ")
	_, err := strconv.Atoi(info[0])
	if err != nil {
		return
	}
	initialLevel, err := strconv.Atoi(info[1])
	if err != nil {
		return
	}

	// 2行目以降
	for len(sc.Text()) > 0 {
		sc.Scan()

		if len(sc.Text()) == 0 {
			break
		}

		level, err := strconv.Atoi(sc.Text())
		if err != nil {
			return
		}

		battleList = append(battleList, level)
	}

	result := handle(handleParam{
		initialLevel: initialLevel,
		battleList:  battleList,
	})
	fmt.Println(result)
}

// ・1 ≦ N ≦ 100,000
// ・1 ≦ x_i ≦ 10,000 (1 ≦ i ≦ N)
// ・1 ≦ L ≦ 10,000
// 　上記の条件をhandleに対するバリデーションでチェックするのはちょっとイメージ違う気がする
// 「A（自分）とB（相手）が戦闘を行う」ことを記録するというI/Fであるべき
// Aは自分自身なので、認証トークンを渡す
// Bは相手なので、パラメータとして指定する
//
// その１ ... １プレイヤーにおける戦闘可能回数と捉えられる
// ・1 ≦ N ≦ 100,000
// ※ 本来的には、相手のプレイヤーも戦闘可能回数は同等のはず
//
// その２ ... 自分ないし、戦闘相手のレベルの最大値と捉えられる
// ・1 ≦ L ≦ 10,000
// ・1 ≦ x_i ≦ 10,000 (1 ≦ i ≦ N)

type handleParam struct {
	initialLevel int
	battleList []int
}

func handle(param handleParam) int {
	currentLevel, battleList := param.initialLevel, param.battleList

	for _, v := range battleList {
		if currentLevel > v {
			currentLevel = currentLevel + (v / 2)
			continue
		}
		if currentLevel == v {
			continue
		}
		if v > currentLevel {
			currentLevel = currentLevel / 2
		}
	}

	return currentLevel
}
