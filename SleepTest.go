package Sleep

import "time"

//import "fmt"

func Sleep(a int64) {

	<-time.After(time.Second * time.Duration(a))

}
