package foo

 import (
     "testing"
 )

func TestSum(t *testing.T) {
    actual := Sum(10, 20)
    expected := 30
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestSum2(t *testing.T) {
    actual := Sum(1, 2)
    expected := 3
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}