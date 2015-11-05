package main

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	lmgr := NewLinkMgr()

	lmgr.pushUrl("www.cnn.com")
	fmt.Println(lmgr.getUrl())

	fmt.Println(lmgr)
}
