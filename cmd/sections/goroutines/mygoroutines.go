package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// / Without using WaitGroups
func proccessTx(txNum int) {
	fmt.Println("Processing Tx:", txNum)
	time.Sleep(2 * time.Second)
	fmt.Println("Completed Tx:", txNum)
}

// / Using WaitGroups
func proccessTx_WG(txNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("$ - [Processing] Tx:", txNum)
	time.Sleep(2 * time.Second)
	fmt.Println("* - [Completed] Tx:", txNum)
}

func Main1() {
	var wg sync.WaitGroup
	var fakeArr [15]uint

	// wg.Add(len(fakeArr))
	for i := 1; i <= len(fakeArr); i++ {
		wg.Add(1)
		// go proccessTx(i)
		go proccessTx_WG(i, &wg)
	}

	// time.Sleep(3 * time.Second) // / Without using WaitGroups
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("|| All Txs completed! ||")
}
