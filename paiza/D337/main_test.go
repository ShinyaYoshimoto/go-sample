
package main

import (
	"testing"
)

func TestD337(t *testing.T) {
	t.Run("ケース1: 温度1: 20, 温度2: 23, 温度3: 18", func(t *testing.T) {
		total, err := handle(struct {
			degree1 int
			degree2 int
			degree3 int
		}{20, 23, 18})
		if err != nil {
			t.Errorf("err is not nil")
		}
		if total != 23 {
			t.Errorf("max degree is not 23")
		}
	})

	t.Run("ケース2: 温度1: 5, 温度2: 10, 温度3: 15", func(t *testing.T) {
		total, err := handle(struct {
			degree1 int
			degree2 int
			degree3 int
		}{5, 10, 15})
		if err != nil {
			t.Errorf("err is not nil")
		}
		if total != 15 {
			t.Errorf("max degree is not 15")
		}
	})
}

