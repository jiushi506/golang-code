package game

import (
	"fmt"
	"sync"
)

var gameWait *sync.WaitGroup

func ConnSocket(wait *sync.WaitGroup) {
	gameWait = wait
	go readMessage()
}

func readMessage() {
	fmt.Println("readMessage:")
	for {
		OK := true
		if OK {
			gameWait.Done() //调用 sync.WaitGroup.Add(-1)
			break  //之前少加这一句，一直报negative counter,要跳出循环
		}
	}
}