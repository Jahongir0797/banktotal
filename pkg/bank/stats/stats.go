package stats

import (
	"github.com/corecudr/bankavg/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	sm := 0
	cnt := 0
	for _, payment := range payments {
		cnt++
		sm += int(payment.Amount)
	}
	return types.Money(sm/cnt)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sm := 0
	for _, payment := range payments {
		if (payment.Category == category) {
			sm += int(payment.Amount)
		} 
	}
	return types.Money(sm)
}