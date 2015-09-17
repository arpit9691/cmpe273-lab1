package Sleep

import "testing"
import "fmt"
import "time"

func Test_Sleep(t *testing.T) {

	t1 := time.Now().Second()
	fmt.Println(t1)

	Sleep(10)

	t2 := time.Now().Second()
	fmt.Println(t2)

	if t2-t1 == 10 {

		t.Log("Pass")
		fmt.Println("Done")
	} else {

		t.Error("Fail")

	}

}
