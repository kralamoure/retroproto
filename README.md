# retroproto

`retroproto` is a library that implements the network protocol of Dofus Retro, for the communication between client and
server. It defines the id and structure of the protocol messages, and it provides functionality for their serialization
and deserialization.

## Requirements:

- [Git](https://git-scm.com/)
- [Go](https://golang.org/)

## Installation

```sh
go get github.com/kralamoure/retroproto
```

## Examples

### Deserialize a client message

```go
package main

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/msgcli"
)

func main() {
	packet := "AAXxRamboPLxX|1|0|3635424|855309|16053493"

	id, _ := retroproto.MsgCliIdByPkt(packet)       // "AA"
	extra := strings.TrimPrefix(packet, string(id)) // "XxRamboPLxX|1|0|3635424|855309|16053493"

	switch id {
	case retroproto.AccountAddCharacter:
		message := &msgcli.AccountAddCharacter{}
		message.Deserialize(extra)

		fmt.Printf("%+v", *message) // {Name:XxRamboPLxX Class:1 Sex:0 Color1:3778e0 Color2:0d0d0d Color3:f4f4f5}
	}
}
```

### Deserialize a server message

```go
package main

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/msgsvr"
)

func main() {
	packet := "HCfdbijergrfklvdnsdfgviojsidesokpm"

	id, _ := retroproto.MsgSvrIdByPkt(packet)       // "HC"
	extra := strings.TrimPrefix(packet, string(id)) // "fdbijergrfklvdnsdfgviojsidesokpm"

	switch id {
	case retroproto.AksHelloConnect:
		message := &msgsvr.AksHelloConnect{}
		message.Deserialize(extra)

		fmt.Printf("%+v", *message) // {Salt:fdbijergrfklvdnsdfgviojsidesokpm}
	}
}
```

### Serialize a client message

```go
package main

import (
	"fmt"

	"github.com/kralamoure/retroproto/msgcli"
)

func main() {
	message := msgcli.AccountAddCharacter{
		Name:   "XxRamboPLxX",
		Class:  1,
		Sex:    0,
		Color1: "3778e0",
		Color2: "0d0d0d",
		Color3: "f4f4f5",
	}

	extra, _ := message.Serialized()

	// Note: packets for msgcli.AccountVersion and msgcli.AccountCredential should not include their protocol ID 
	packet := fmt.Sprintf("%s%s", message.ProtocolId(), extra)

	fmt.Print(packet) // "AAXxRamboPLxX|1|0|3635424|855309|16053493"
}
```

### Serialize a server message

```go
package main

import (
	"fmt"

	"github.com/kralamoure/retroproto/msgsvr"
)

func main() {
	message := msgsvr.AksHelloConnect{Salt: "fdbijergrfklvdnsdfgviojsidesokpm"}

	extra, _ := message.Serialized()

	packet := fmt.Sprintf("%s%s", message.ProtocolId(), extra)

	fmt.Print(packet) // "HCfdbijergrfklvdnsdfgviojsidesokpm"
}
```
