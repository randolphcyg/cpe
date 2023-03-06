# cpe

Used to resolve the URL of CPE2.3 or CPE2.3 into a golang structure

## Usage

```shell
go get "github.com/randolphcyg/cpe"
```

```go
package main

import (
	"fmt"

	"github.com/randolphcyg/cpe"
)

func main() {
	cpeString := "cpe:/a:teamspeak:teamspeak2:2.0.23.19:tes /t:test2/"
	c, err := cpe.ParseCPE(cpeString)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)

	fmt.Println("Part:", c.Part)
	fmt.Println("Vendor:", c.Vendor)
	fmt.Println("Product:", c.Product)
	fmt.Println("Version:", c.Version)
	fmt.Println("Update:", c.Update)
	fmt.Println("Edition:", c.Edition)
	fmt.Println("Language:", c.Language)
}
```