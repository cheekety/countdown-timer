package main

import (
	"fmt"
	"time"
)

func countdown(seconds int){
	for seconds>0 {
		seconds --
		fmt.Println(seconds)
		time.Sleep(1 * time.Second)
	}
}

func main(){
	var seconds int
	fmt.Println("Seconds to countdown: ")
	fmt.Scanf("%d", &seconds)
	
	startHour, startMin, startSec := time.Now().Clock()

	countdown(seconds)

	endHour, endMin, endSec := time.Now().Clock()

	fmt.Println("\nCountdown Complete!")

	fmt.Println("\nStart time: ", startHour, startMin, startSec)
	fmt.Println("\nEnd time: ", endHour, endMin, endSec)
}
