package Fibbonaci

import "testing"

//import "fmt"

func Test_fibbo(t *testing.T) {

	if fibo(5) != 5 {

		t.Error("Failed")

	} else {

		t.Log("Pass")

	}

}
