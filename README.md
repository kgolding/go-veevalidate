# go-veevalidate

Create Veevalidate rules in go.

Example:

````
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kgolding/go-veevalidate"
)

func main() {
	v := veevalidate.New().
		Required().
		MaxValue(100).
		MaxValue(0xffff)

	fmt.Println(v.String())

	myStruct := struct {
		Vee veevalidate.VBuilder
	}{
		Vee: veevalidate.New().Required().IP_or_FQDN(),
	}

	b, _ := json.Marshal(myStruct)

	fmt.Println(string(b))

}


````

https://play.golang.org/p/a7ZI7T9Pw7O