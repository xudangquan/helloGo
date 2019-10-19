package testModel

import (
	"errors"
	"fmt"
	"helloGo/model"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	a := model.Acan{"aa", "bb"}
	a.SetV()
	fmt.Println(a)
	SSetV(&a)
	fmt.Println(a)
	fmt.Println(time.Now().Format("20060102"))
	fmt.Println(time.Now().Format("060102"))
}
func SSetV(acan *model.Acan) {
	acan.Name = "SSE"
}

func TestError(t *testing.T) {
	err := errors.New("abc")
	if err.Error() == "abc" {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
