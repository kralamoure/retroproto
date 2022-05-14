package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type DialogCreateSuccess struct {
	NPCId int
}

func NewDialogCreateSuccess(extra string) (DialogCreateSuccess, error) {
	var m DialogCreateSuccess

	if err := m.Deserialize(extra); err != nil {
		return DialogCreateSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DialogCreateSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogCreateSuccess
}

func (m DialogCreateSuccess) MessageName() string {
	return "DialogCreateSuccess"
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
