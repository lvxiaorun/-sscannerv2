package tool

import (
	"fmt"
	"testing"
)

func TestFisherYates(t *testing.T) {
	ss := []string{"1", "2", "3", "4", "5"}
	fmt.Println(FisherYates(ss))
}
