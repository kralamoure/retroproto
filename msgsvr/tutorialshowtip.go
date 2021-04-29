package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type TutorialShowTip struct {
	Id int
}

func (m TutorialShowTip) ProtocolId() d1proto.MsgSvrId {
	return d1proto.TutorialShowTip
}

func (m TutorialShowTip) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *TutorialShowTip) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1proto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
