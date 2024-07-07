package trie

import (
	"fmt"
	"testing"
)

func TestAdd1(t *testing.T) {
	testString := string("aboba")
	trie := Trie{}
	fmt.Printf("trie is ... %v\n", trie)
	trie.Add(testString)
	if trie.Consists(testString) != true {
		t.Errorf("%s adding failed", testString)
	}

}


func  TestAdd2(t *testing.T) {
	testString := string("flower")
	trie := Trie{}
	trie.Add(testString)
	if trie.Consists(testString) != true {
		t.Errorf("%s adding failed", testString)
	}

}

func TestAdd3(t *testing.T) {
	testString := string("flower")
	testString2 := string("flowers")
	trie := Trie{}
	trie.Add(testString)
	trie.Add(testString2)
	if !trie.Consists(testString) || !trie.Consists(testString2) {
		t.Errorf("%s and %s adding failed", testString, testString2)
	}
}



