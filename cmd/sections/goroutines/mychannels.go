package goroutines

import (
	"fmt"
	"time"
)

// / Using Channels as WaitGroups
func proccessTx_Channels_as_WG(txNum int, done chan<- bool) {
	fmt.Println("$ - [Processing] Tx:", txNum)
	time.Sleep(2 * time.Second)
	fmt.Println("* - [Completed] Tx:", txNum)
	done <- true // send a signal that the Tx is completed
}

// Using Channels to pass Data
func proccessTx_Channels_Data(txNum int, done chan<- int) {
	fmt.Println("$ - [Processing] Tx:", txNum)
	time.Sleep(2 * time.Second)
	fmt.Println("* - [Completed] Tx:", txNum)
	done <- txNum // send a signal that the Tx is completed
}

func Main2() {
	/// This example does the same thing as WaitGroups
	// Creating a Channel
	// complete := make(chan bool)

	// for i := 1; i < 5; i++ {
	// 	go proccessTx_Channels_as_WG(i, complete)
	// }

	// for i := 1; i < 5; i++ {
	// 	<-complete // wait for a signal from each goroutine
	// }

	// fmt.Println("|| All Txs completed! ||")
	///////////////////////////////////////////////////////////////////
	/// Using Channels to transfer Data
	totalTransactions := 5
	processed := make(chan int, totalTransactions) // create a channel to communicate completion

	for i := 1; i <= totalTransactions; i++ {
		go proccessTx_Channels_Data(i, processed)
	}

	// This is super weird syntax.
	// Here we basically say:
	// 1. Infinte loop
	for txNum := range processed {
		fmt.Printf("Received completion signal fro Tx %d\n", txNum)
		if txNum == totalTransactions {
			close(processed) // Close the channel when the last TX is completed.
		}
	}

}
