package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type DialogBeginning struct {
	NPCId int
}

func (m DialogBeginning) ProtocolId() retroproto.MsgCliId {
	return retroproto.DialogBeginning
}

func (m DialogBeginning) Serialized() (string, error) {
	return fmt.Sprint(m.NPCId), nil
}

func (m *DialogBeginning) Deserialize(extra string) error {
	npcId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.NPCId = int(npcId)
	return nil
}
