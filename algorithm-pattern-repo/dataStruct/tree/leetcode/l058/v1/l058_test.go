package v1

import (
	"fmt"
	"testing"
)

func TestL058(t *testing.T) {
	var res bool
	calender := Constructor()
	res = calender.Book(10, 20)
	if res {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	res = calender.Book(15, 25)
	if res {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	res = calender.Book(20, 30)
	if res {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
