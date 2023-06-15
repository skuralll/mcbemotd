# mcbemotd

You can get MOTD of the MinecraftBE server.

## Getting started

Execute the following command.

```shell
go get github.com/skuralll/mcbemotd@latest
```

## Example

```go
import (
	"fmt"
	"github.com/skuralll/mcbemotd"
)

func main() {
	motd, err := mcbemotd.GetServerInfo("localhost:19132")
	if err != nil {
		// error handling
	} else {
		fmt.Println(motd.Edition)    // Edition
		fmt.Println(motd.Motd1)      // MOTD line 1
		fmt.Println(motd.Protocol)   // Protocol version
		fmt.Println(motd.Version)    // Server version
		fmt.Println(motd.Players)    // Player Count
		fmt.Println(motd.PlayersMax) // Max Player Count
		fmt.Println(motd.Uid)        // Server unique id
		fmt.Println(motd.Motd2)      // MOTD line 2
		fmt.Println(motd.ModeStr)    // Game mode (string)
		fmt.Println(motd.ModeNum)    // Game mode (numeric)
		fmt.Println(motd.Portv4)     // Srever port (v4)
		fmt.Println(motd.Portv6)     // Srever port (v6)
	}
}
/*
Output:
  MCPE
  MOTD Line 1
  589
  1.20.0
  0
  20
  -8033921687892980065
  MOTD Line 2
  Survival
  1
  19132
  19133
*/
```
