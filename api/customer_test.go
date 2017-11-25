package api

import (
	"github.com/jadengis/modserver/duid"
	"testing"
)

func TestMakeNewCustomer(t *testing.T) {
	SetIdGenerator(duid.NewGenerator(0))
	var customer = NewCustomer("Acme")

	if customer.Id() == 0 {
		t.Errorf("Customer id not sensible: id == %d", int(customer.Id()))
	}

	if customer.Name() != "Acme" {
		t.Errorf("Customer name not correct: %s instead of Acme", string(customer.Name()))
	}

	if customer.TableName() != CustomerTableName {
		t.Errorf("Customer table name not set correctly: set as %s", customer.TableName())
	}

	customer.SetId(123)
	if customer.Id() != 123 {
		t.Errorf("Customer id not correct: id == %d", int(customer.Id()))
	}

	customer.SetName("Big Pharma")
	if customer.Name() != "Big Pharma" {
		t.Errorf("Customer name not correct: %s instead of Big Pharma", string(customer.Name()))
	}

}
