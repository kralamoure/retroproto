package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type DialogCreate struct {
	NPCId int
}

func (m DialogCreate) MessageId() retroproto.MsgCliId {
	return retroproto.DialogCreate
}

func (m DialogCreate) MessageName() string {
	return "DialogCreate"
}

func (m DialogCreate) Serialized() (string, error) {
	return fmt.Sprint(m.NPCId), nil
}

func (m *DialogCreate) Deserialize(extra string) error {
	npcId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.NPCId = int(npcId)
	return nil
}
