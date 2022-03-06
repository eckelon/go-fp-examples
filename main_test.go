package main

import (
	"reflect"
	"testing"
)

func TestUsingCompose(t *testing.T) {
	// desiredResult := []int{1, 2, 3}
	result, _ := usingCompose(2)
	if result != 8 {
		t.Error("usingCompose(2) should return 8")
	}
}

func TestUsingTraverse(t *testing.T) {
	list := []int{1, 2, 3}
	desiredResult := []int{3, 5, 7}
	result, _ := usingTraverse(list)
	if !reflect.DeepEqual(result, desiredResult) {
		t.Error("usingTraverse(", list, ") should return", desiredResult)
	}
}
