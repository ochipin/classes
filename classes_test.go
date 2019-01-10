package classes

import (
	"fmt"
	"testing"
)

func Test__NEW_CLASSES(t *testing.T) {
	class := New("test")

	strs, err := class.Classlist()
	if err != nil {
		t.Fatal(err)
	}

	if len(strs) != 1 {
		t.Fatal("Classes: Error")
	}

	if strs[0] != "App" {
		t.Fatal("Classes: Error")
	}

	fmt.Println(strs)

	class = New("test2")
	if _, err := New("test2").Classlist(); err == nil {
		t.Fatal("Classes: Error")
	}
}

func Test__NEW_CLASSES_NO_EXPORTED(t *testing.T) {
	class := New("test")
	class.ExportOnly = false

	strs, err := class.Classlist()
	if err != nil {
		t.Fatal(err)
	}

	if len(strs) != 2 {
		t.Fatal("Classes: Error")
	}

	var check = map[string]bool{
		"base": true,
		"App":  true,
	}
	for _, v := range strs {
		if _, ok := check[v]; !ok {
			t.Fatal("Classes: Error")
		}
	}
}
