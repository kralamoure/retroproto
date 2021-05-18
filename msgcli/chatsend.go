package msgcli

import (
	"fmt"
	"html"
	"strings"

	"github.com/kralamoure/dofus/dofustyp"

	"github.com/kralamoure/retroproto"
)

type ChatSend struct {
	ChatChannel     dofustyp.ChatChannel
	PrivateReceiver string
	Message         string
	Params          string // TODO
}

func (m ChatSend) ProtocolId() retroproto.MsgCliId {
	return retroproto.ChatSend
}

func (m ChatSend) Serialized() (string, error) {
	dest := string(m.ChatChannel)
	if m.ChatChannel == dofustyp.ChatChannelPrivate {
		dest = m.PrivateReceiver
	}

	return fmt.Sprintf("%s|%s|%s", dest, m.Message, m.Params), nil
}

func (m *ChatSend) Deserialize(extra string) error {
	sli := strings.SplitN(extra, "|", 3)
	if len(sli) != 3 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] == "" {
		return retroproto.ErrInvalidMsg
	}

	var r rune
	for _, v := range sli[0] {
		r = v
		break
	}
	chatChannel := dofustyp.ChatChannel(r)

	switch chatChannel {
	case '¤':
		chatChannel = dofustyp.ChatChannelNewbies
	case dofustyp.ChatChannelPrivate:
		return retroproto.ErrInvalidMsg
	}

	if len(sli[0]) >= 2 {
		chatChannel = dofustyp.ChatChannelPrivate
		m.PrivateReceiver = html.EscapeString(sli[0])
	}

	_, ok := dofustyp.ChatChannels[chatChannel]
	if !ok {
		return retroproto.ErrInvalidMsg
	}

	m.ChatChannel = chatChannel

	message := html.EscapeString(sli[1])
	message = strings.TrimSpace(message)
	m.Message = message

	m.Params = html.EscapeString(sli[2])

	return nil
}
