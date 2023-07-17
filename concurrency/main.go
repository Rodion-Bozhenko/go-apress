package main

import (
	"fmt"
)

func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

func main() {
	// fmt.Println("main function started")
	// CalcStoreTotal(Products)
	// fmt.Println("main function complete")

	dispatchChannel := make(chan DispatchNotification, 100)

	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	// go DispatchOrders(sendOnlyChannel)
	//  receiveDispatches(receiveOnlyChannel)
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	receiveDispatches((<-chan DispatchNotification)(dispatchChannel))

	// for {
	// 	if details, open := <-dispatchChannel; open {
	// fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	// 	} else {
	// 		fmt.Println("Channel has been closed")
	// 		break
	// 	}
	// }

	// for details := range dispatchChannel {
	// 	fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	// }
	// fmt.Println("Channel has been closed")

}
