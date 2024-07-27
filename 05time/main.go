package main

import (
	"fmt"
	"time"
)

func main() {
	presenttime := time.Now()
	fmt.Println(presenttime)
	fmt.Println(presenttime.Format("02-01-2006 15:04:05 Monday"))

	createdate := time.Date(2024, time.March, 7, 12, 0, 0, 0, time.UTC)
	fmt.Println(createdate)
	fmt.Println(createdate.Format("02-01-2006 Monday"))

}

//in terminal
// GOOS="darwin" go build {build exe file which can be run in mac}
