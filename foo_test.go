package foo

 import (
     "testing"
     "github.com/MatsuGitHub/test"
 )

func TestSum(t *testing.T) {
    actual := foo.Sum(10, 20)
    expected := 30
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestSum2(t *testing.T) {
    actual := foo.Sum(1, 2)
    expected := 3
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}