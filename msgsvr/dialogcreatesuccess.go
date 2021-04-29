package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type DialogCreateSuccess struct {
	NPCId int
}

func (m DialogCreateSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.DialogCreateSuccess
}

func (m DialogCreateSuccess) Serialized() (string, error) {
	return fmt.Sprint(m.NPCId), nil
}

func (m *DialogCreateSuccess) Deserialize(extra string) error {
	npcId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.NPCId = int(npcId)
	return nil
}
