package util

import "testing"

func TestFeed(t *testing.T) {

	data := Data{
		"xuqingfeng",
		"gotpl",
	}
	fs, err := Feed(data)
	if err != nil {
		t.Errorf("E! %v", err)
	}
	t.Logf("I! %v", fs)
}
