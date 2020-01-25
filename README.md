# ALTV Go API

[altv-Homepage](https://altv.mp/)

##### This depends on Hazards CAPI
[altv-capi (Gitlab)](https://gitlab.com/7Hazard/altv-capi)


#### Example Resource

```golang
package main

import (
	"fmt"

	"github.com/immali/go-altv"
)

func main() {
	r := gocapi.NewResource()
	r.OnPlayerConnect(func(p *gocapi.Player) {
		fmt.Println(fmt.Sprintf("Player (%s) with ID %d has connected", p.Name, p.ID))
		p.Spawn(gocapi.Vector3{
			X: -425.517,
			Y: 1123.620,
			Z: 325.8544,
		})
	})

	ch := make(chan int)
	<-ch
}

```
