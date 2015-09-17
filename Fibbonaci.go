package main
package Fibbonaci

import "fmt"

func main() {
	var num int64
	fmt.Println("Enter an Number: ")
	fmt.Scanf("%f", &num)
	fmt.Println(fibo(num))
}

func fibo(num int64) int64 {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibo(num-1) + fibo(num-2)
	}

}
