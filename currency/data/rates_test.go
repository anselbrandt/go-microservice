package data

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestNewRates(t *testing.T) {
	tr, err := NewRates(hclog.Default())

	if err != nil {
		t.Fatal(err)
	}

	for key, element := range tr.rates {
		fmt.Println(key, "=>", element)
	}
	// fmt.Printf("Rates %#v", tr.rates)
}
