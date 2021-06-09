package main

import (
	"fmt"
	"testing"

	"route"
)

func Test_GetRandCode(t *testing.T){

	retString := route.GetRandCode(8)
	if len(retString) != 8{
		fmt.Println("Test_GetRandCode error")
	}else {
		fmt.Println("Test_GetRandCode pass")
	}
}