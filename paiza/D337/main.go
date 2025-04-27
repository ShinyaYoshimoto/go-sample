package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	// "slices"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	degree1, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
	}
	sc.Scan()
	degree2, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
	}
	sc.Scan()
	degree3, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
	}

	total, err := handle(struct {
		degree1 int
		degree2 int
		degree3 int
	}{degree1, degree2, degree3})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(total)
	}
}

func handle(param struct {
	degree1 int
	degree2 int
	degree3 int
}) (int, error) {
	degree1, degree2, degree3 := param.degree1, param.degree2, param.degree3

	if !validate(degree1, degree2, degree3) {
		return 0, errors.New("invalid input")
	}

	degrees := []int{degree1, degree2, degree3}

	return max(degrees), nil
}

const (
	minDegree = 0
	maxDegree = 40
)

func validate(degree1 int, degree2 int, degree3 int) bool {
	return degree1 >= minDegree && degree1 <= maxDegree &&
		degree2 >= minDegree && degree2 <= maxDegree &&
		degree3 >= minDegree && degree3 <= maxDegree
}

// MEMO: 25.04.27時点
// paizaの動作環境はgo1.21.0未満なので、slices.Maxが使えない
func max(degrees []int) int {
	max := degrees[0]
	for _, degree := range degrees {
		if degree > max {
			max = degree
		}
	}
	return max
}
