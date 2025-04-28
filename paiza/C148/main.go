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
	battleCount, err := strconv.Atoi(info[0])
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

	result := handle(battleCount, initialLevel, battleList)
	fmt.Println(result)
}

func handle(battleCount int, initialLevel int, battleList []int) int {
	currentLevel := initialLevel

	for i := 0; i < battleCount; i++ {
		if (currentLevel > battleList[i]) {
			currentLevel = currentLevel + (battleList[i] / 2)
			continue
		}

		if (currentLevel == battleList[i]) {
			continue
		}

		if battleList[i] > currentLevel {
			currentLevel = currentLevel / 2
		}
	}

	return currentLevel
}
