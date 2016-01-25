package golang

import (
	"testing"
	"fmt"
)

func TestHello(t *testing.T)  {
	str := Hello("linguofeng")
	t.Log(str)
}

type CallbackImpl struct {

}

func (c CallbackImpl) CallByGo(str string) {
	fmt.Println(str)
}

func TestCallJava(t *testing.T)  {
	SendStr("test", CallbackImpl{})
}