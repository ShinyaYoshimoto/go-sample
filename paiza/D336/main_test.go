package main

import (
	"testing"
)

func TestD336(t *testing.T) {
	t.Run("ケース1: 前に並んでいる人数: 3人, 後ろに並んでいる人数: 5人", func(t *testing.T) {
		total, err := handle(3, 5)
		if err != nil {
			t.Errorf("err is not nil")
		}
		if total != 9 {
			t.Errorf("total is not 9")
		}
	})

	t.Run("ケース2: 前に並んでいる人数: 1人, 後ろに並んでいる人数: 1人", func(t *testing.T) {
		total, err := handle(1, 1)
		if err != nil {
			t.Errorf("err is not nil")
		}
		if total != 3 {
			t.Errorf("total is not 3")
		}
	})
}
