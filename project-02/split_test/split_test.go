package main

import (
	"reflect"
	"testing"
)

// 通过 go test -v执行所有的单元测试
func TestSplit(t *testing.T) {
	result := split("abcd", "b")
	want := []string{"a", "cd"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("want %#v, got %#v", want, result)
	}
}

func Test2Split(t *testing.T) {
	result := split("abcd", "bc")
	want := []string{"a1", "cd"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("want %#v, got %#v", want, result)
	}
}
