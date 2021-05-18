package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type TutorialShowTip struct {
	Id int
}

func (m TutorialShowTip) ProtocolId() retroproto.MsgSvrId {
	return retroproto.TutorialShowTip
}

func (m TutorialShowTip) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *TutorialShowTip) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
