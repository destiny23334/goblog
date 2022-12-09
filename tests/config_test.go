package tests

import (
	"fmt"
	"goblog/utils"
	"testing"
)

func TestConfig(t *testing.T) {
	c := utils.InitConfig()
	fmt.Println(c)
}
