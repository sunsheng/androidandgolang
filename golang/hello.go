package golang

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

type Callback interface  {
	CallByGo(string)
}

func SendStr(str string, c Callback) {
	str2 := "Go " + str
	c.CallByGo(str2)
}
