# d1proto

`d1proto` is a network protocol library for Dofus 1 client/server communication.

Some examples of what you can do with this library:

- Deserialize client/server packets into structured data
- Serialize structured data into client/server packets

## Requirements:

- [Git](https://git-scm.com/)
- [Go](https://golang.org/)

## Installation

```sh
go get github.com/kralamoure/d1proto
```

## Usage

### Deserializing a client packet into structured data
 
```go
package main

import (
    "fmt"
    "strings"

    "github.com/kralamoure/d1proto"
    "github.com/kralamoure/d1proto/msgcli"
)

func main() {
    packet := "AAXxRamboPLxX|1|0|3635424|855309|16053493"

    id, _ := d1proto.MsgCliIdByPkt(packet)          // "AA"
    extra := strings.TrimPrefix(packet, string(id)) // "XxRamboPLxX|1|0|3635424|855309|16053493"

    switch id {
    case d1proto.AccountAddCharacter:
        message := &msgcli.AccountAddCharacter{}
        message.Deserialize(extra)

        fmt.Printf("%+v", *message) // {Name:XxRamboPLxX Class:1 Sex:0 Color1:3778e0 Color2:0d0d0d Color3:f4f4f5}
    }
}
```

### Deserializing a server packet into structured data
 
```go
package main

import (
    "fmt"
    "strings"

    "github.com/kralamoure/d1proto"
    "github.com/kralamoure/d1proto/msgsvr"
)

func main() {
    packet := "HCfdbijergrfklvdnsdfgviojsidesokpm"

    id, _ := d1proto.MsgSvrIdByPkt(packet)          // "HC"
    extra := strings.TrimPrefix(packet, string(id)) // "fdbijergrfklvdnsdfgviojsidesokpm"

    switch id {
    case d1proto.AksHelloConnect:
        message := &msgsvr.AksHelloConnect{}
        message.Deserialize(extra)

        fmt.Printf("%+v", *message) // {Salt:fdbijergrfklvdnsdfgviojsidesokpm}
    }
}
```

### Serializing structured data into a client packet
 
```go
package main

import (
    "fmt"

    "github.com/kralamoure/d1proto/msgcli"
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

### Serializing structured data into a server packet
 
```go
package main

import (
    "fmt"

    "github.com/kralamoure/d1proto/msgsvr"
)

func main() {
    message := msgsvr.AksHelloConnect{Salt: "fdbijergrfklvdnsdfgviojsidesokpm"}

    extra, _ := message.Serialized()
 
    packet := fmt.Sprintf("%s%s", message.ProtocolId(), extra)

    fmt.Print(packet) // "HCfdbijergrfklvdnsdfgviojsidesokpm"
}
```
