/* Hands-on exercise #1
Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this
hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a
package.
solution: https://play.golang.org/p/8BK6PXj3aG */

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func main() {
	u1 := user{
		Name: "James",
		Age:  32,
	}
	u2 := user{
		Name: "Monneypenny",
		Age:  27,
	}
	u3 := user{
		Name: "M",
		Age:  54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

// func Marshal(v interface{}) ([]byte, error)
