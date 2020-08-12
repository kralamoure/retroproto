package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/dofus/dofustyp"

	"github.com/kralamoure/d1proto"
)

type ChatMessageSuccess struct {
	ChatChannel dofustyp.ChatChannel
	Id          int
	Name        string
	Message     string
	Params      string // TODO
}

func (m ChatMessageSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatMessageSuccess
}

func (m ChatMessageSuccess) Serialized() (string, error) {
	chatChannel := string(m.ChatChannel)
	switch m.ChatChannel {
	case dofustyp.ChatChannelPublic:
		chatChannel = ""
	case dofustyp.ChatChannelPrivate:
		chatChannel = "T"
	}

	return fmt.Sprintf("%s|%d|%s|%s|%s", chatChannel, m.Id, m.Name, m.Message, m.Params), nil
}

func (m *ChatMessageSuccess) Deserialize(extra string) error {
	sli := strings.SplitN(extra, "|", 5)
	if len(sli) != 5 {
		return d1proto.ErrInvalidMsg
	}

	switch sli[0] {
	case "":
		m.ChatChannel = dofustyp.ChatChannelPublic
	case "T":
		m.ChatChannel = dofustyp.ChatChannelPrivate
	default:
		for _, v := range sli[0] {
			m.ChatChannel = dofustyp.ChatChannel(v)
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
	m.Params = sli[4]

	return nil
}
