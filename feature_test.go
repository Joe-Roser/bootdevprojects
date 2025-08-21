package main

import (
	"testing"
)

func building_test(t *testing.T) {
	conf := config{
		next_offset: 0,
	}
	building(&conf, "")
	t.Error()
}
