
package main

import (
	"testing"
)

func TestC148(t *testing.T) {
	t.Run("ケース1: 入力1: 5, 入力2: 10, 入力3: 5, 11, 20, 8, 7", func(t *testing.T) {
		total := handle(5, 10, []int{5, 11, 20, 8, 7})
		if total != 11 {
			t.Errorf("total is not 10")
		}
	})

	t.Run("ケース2: 入力1: 3, 入力2: 9, 入力3: 10, 4, 4", func(t *testing.T) {
		total := handle(3, 9, []int{10, 4, 4})
		if total != 4 {
			t.Errorf("total is not 9")
		}
	})
}

