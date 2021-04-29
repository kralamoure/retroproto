package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type DialogBeginning struct {
	NPCId int
}

func (m DialogBeginning) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.DialogBeginning
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
