package main

import (
	"fmt"
	"strings"

	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/msgcli"
)

func main() {
	packet := "AAXxRamboPLxX|1|0|3635424|855309|16053493"

	id, _ := d1proto.MsgCliIdByPkt(packet)          // id: "AA"
	extra := strings.TrimPrefix(packet, string(id)) // extra: "XxRamboPLxX|1|0|3635424|855309|16053493"

	switch id {
	case d1proto.AccountAddCharacter:
		message := &msgcli.AccountAddCharacter{}
		message.Deserialize(extra)

		fmt.Printf("%+v", message) //
	}
}
