package android

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

func (c CallbackImpl) CallJava() {
	fmt.Println("CallJava")
}

func TestCallJava(t *testing.T)  {
	CallJava(CallbackImpl{})
}