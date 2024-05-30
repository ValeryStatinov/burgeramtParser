package parser

import "testing"

func TestRegexpMatchers(t *testing.T) {
	t.Run("regexp date", func(t *testing.T) {
		s := "20.06.2024 - An diesem Tag einen Termin buchen"
		d := regexpDate(s)

		if d != "20.06.2024" {
			t.Fatal("failed to parse string")
		}
	})
}
