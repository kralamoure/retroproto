package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type DialogCreate struct {
	NPCId int
}

func (m DialogCreate) ProtocolId() d1proto.MsgCliId {
	return d1proto.DialogCreate
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
