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
