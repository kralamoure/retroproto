package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type ChatMessageSuccess struct {
	ChatId  string
	Id      int
	Name    string
	Message string
}

func (m ChatMessageSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatMessageSuccess
}

func (m ChatMessageSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%s|%d|%s|%s", m.ChatId, m.Id, m.Name, m.ChatId), nil
}

func (m *ChatMessageSuccess) Deserialize(extra string) error {
	sli := strings.SplitN(extra, "|", 4)
	if len(sli) != 4 {
		return d1proto.ErrInvalidMsg
	}

	m.ChatId = sli[0]

	id, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	m.Name = sli[2]
	m.Message = sli[3]

	return nil
}
