// struct demo
package main

import (
	"fmt"
	"os"
)

// Trade is a trade in stocks
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)

	t, err := NewTrade("MSFT", 10, 99.98, true)

	if err != nil {
		fmt.Printf("error: can't create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.nValue())
}

// Value returns the trade value
// pointer to trade method name and type
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol can't be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("Volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >=0 (was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil

}

// Value returns the trade value
func (t *Trade) nValue() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}
