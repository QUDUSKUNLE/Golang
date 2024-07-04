package main

import (
	"fmt"
	"log"

	order "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/order"
	shipping "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/shipping"
	product "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/shipping/product"
)

func main() {
	wallet := order.NewOrderFacade("abc", 10)

	fmt.Println()
	if err := wallet.AddMoneyToWallet("abc", 10, 10000); err != nil {
		log.Fatalf("Error: %s\n\n", err.Error())
	}

	fmt.Println()
	if err := wallet.DeductMoneyFromWallet("abc", 10, 10); err != nil {
		log.Fatalf("Error: %s\n\n", err.Error())
	}

	fmt.Println()
	ship := shipping.NewShipping("abc", product.Ammunition)

	fmt.Println()
	if err := ship.ScheduleShipping("abc", "41 Jibowu Estate Road", product.Ammunition); err != nil {
		log.Fatalf("Error: %s\n\n", err.Error())
	}
}
