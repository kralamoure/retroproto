package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/dofus/dofustyp"

	"github.com/kralamoure/retroproto"
)

type ChatMessageSuccess struct {
	ChatChannel dofustyp.ChatChannel
	Id          int
	PrivateTo   bool
	Name        string
	Message     string
	Params      string // TODO
}

func NewChatMessageSuccess(extra string) (ChatMessageSuccess, error) {
	var m ChatMessageSuccess

	if err := m.Deserialize(extra); err != nil {
		return ChatMessageSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatMessageSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ChatMessageSuccess
}

func (m ChatMessageSuccess) MessageName() string {
	return "ChatMessageSuccess"
}

func (m ChatMessageSuccess) Serialized() (string, error) {
	chatChannel := string(m.ChatChannel)
	switch m.ChatChannel {
	case dofustyp.ChatChannelPublic:
		chatChannel = ""
	case dofustyp.ChatChannelPrivate:
		if m.PrivateTo {
			chatChannel = "T"
		} else {
			chatChannel = "F"
		}
	}

	return fmt.Sprintf("%s|%d|%s|%s|%s", chatChannel, m.Id, m.Name, m.Message, m.Params), nil
}

func (m *ChatMessageSuccess) Deserialize(extra string) error {
	sli := strings.SplitN(extra, "|", 5)
	if len(sli) != 5 {
		return retroproto.ErrInvalidMsg
	}

	switch sli[0] {
	case "":
		m.ChatChannel = dofustyp.ChatChannelPublic
	case "F":
		m.ChatChannel = dofustyp.ChatChannelPrivate
	case "T":
		m.ChatChannel = dofustyp.ChatChannelPrivate
		m.PrivateTo = true
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
