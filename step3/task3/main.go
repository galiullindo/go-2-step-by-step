package main

func sendOneTwoThree(channel chan int) {
	for i := 0; i < 3; i++ {
		channel <- i
	}
}

func Send(channel1 chan int, channel2 chan int) {
	go sendOneTwoThree(channel1)
	go sendOneTwoThree(channel2)
}
