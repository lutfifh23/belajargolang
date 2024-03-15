package main

func main() {
	// empty interface
	// var randomValues interface{}

	// _ = randomValues

	// randomValues = "Jalan Sudirman"

	// randomValues = 20

	// randomValues = true

	// randomValues = []string{"Aurellia", "Jen"}

	// Empty interface (Type assertion)
	// var v interface{}

	// v = 20

	// if value, ok := v.(int); ok == true {
	// 	v = value * 9
	// }

	// Empty interface (Map & Slice with empty interface)
	rs := []interface{}{1, "Aurellia", true, 2, "Lulu", true}

	rm := map[string]interface{}{
		"Name":   "Aurellia",
		"Status": true,
		"Age":    21,
	}

	_, _ = rs, rm
}
