package android

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

type Callback interface  {
	callJava()
}

func CallJava(c Callback) {
	c.callJava()
}
