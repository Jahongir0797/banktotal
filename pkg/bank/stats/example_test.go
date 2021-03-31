package stats

import (
	"fmt"

	"github.com/corecudr/bankavg/pkg/bank/types")


func ExampleTotalInCategory() {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 10_000,
			Category: "Standart",
		},
		{
			ID: 2,
			Amount: 10_000,
			Category: "Normal",
		},
		{
			ID: 3,
			Amount: 30_000,
			Category: "Standart",
		},
	}
	answer := TotalInCategory(payments, "Standart")
	fmt.Println(answer)
	// Output: 40000
}