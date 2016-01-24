package android

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

type Callback interface  {
	CallJava()
}

func CallJava(c Callback) {
	c.CallJava()
}
