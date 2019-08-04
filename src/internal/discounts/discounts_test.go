package discounts

import (
	"testing"

	"github.com/brucebales/discounts-api/src/internal/dto"
)

func TestSwitchDiscount(t *testing.T) {
	switches := []dto.Item{
		{
			ProductID: "B103",
			Quantity:  6,
			UnitPrice: 12.95,
			Total:     77.7,
		},
	}

	expAmount := 12.95

	expReason := "switchDiscount"

	amount, reason := SwitchDiscount(switches)

	if expAmount != amount {
		t.Errorf("Expected amount: %v\n Actual amount: %v\n", expAmount, amount)
	}
	if expReason != reason {
		t.Errorf("Expected reason: %s\n Actual reason: %s\n", expReason, reason)
	}

}

func TestToolDiscount(t *testing.T) {
	tools := []dto.Item{
		{
			ProductID: "A101",
			Quantity:  3,
			UnitPrice: 10,
			Total:     30,
		},
		{
			ProductID: "A102",
			Quantity:  3,
			UnitPrice: 49.50,
			Total:     148.5,
		},
	}

	expAmount := 2.0

	expReason := "FirstToolDiscount"

	amount, reason := ToolDiscount(tools)
	if expAmount != amount {
		t.Errorf("Expected amount: %v\n Actual amount: %v\n", expAmount, amount)
	}
	if expReason != reason {
		t.Errorf("Expected reason: %s\n Actual reason: %s\n", expReason, reason)
	}

}
