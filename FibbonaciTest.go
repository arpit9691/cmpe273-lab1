package Fibbonaci

//import "fmt"

func fibo(num int64) int64 {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibo(num-1) + fibo(num-2)
	}

}
