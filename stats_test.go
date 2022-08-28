package main

import (
	"testing"
	"fmt"
	"github.com/montanaflynn/stats"
)

func TestStats(t *testing.T) {
	fmt.Println(stats.NormCdf(1.09,0,1))
}
