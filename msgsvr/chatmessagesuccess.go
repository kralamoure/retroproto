package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/dofus/dofustyp"

	"github.com/kralamoure/d1proto"
)

type ChatMessageSuccess struct {
	ChatId  dofustyp.ChatChannel
	Id      int
	Name    string
	Message string
}

func (m ChatMessageSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatMessageSuccess
}

func (m ChatMessageSuccess) Serialized() (string, error) {
	chatId := string(m.ChatId)
	switch m.ChatId {
	case dofustyp.ChatChannelPublic:
		chatId = ""
	case dofustyp.ChatChannelPrivate:
		chatId = "T"
	}

	return fmt.Sprintf("%s|%d|%s|%s", chatId, m.Id, m.Name, m.Message), nil
}

func (m *ChatMessageSuccess) Deserialize(extra string) error {
	sli := strings.SplitN(extra, "|", 4)
	if len(sli) != 4 {
		return d1proto.ErrInvalidMsg
	}

	switch sli[0] {
	case "":
		m.ChatId = dofustyp.ChatChannelPublic
	case "T":
		m.ChatId = dofustyp.ChatChannelPrivate
	default:
		for _, v := range sli[0] {
			m.ChatId = dofustyp.ChatChannel(v)
			break
		}
	}

	id, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	m.Name = sli[2]
	m.Message = sli[3]

	return nil
}
