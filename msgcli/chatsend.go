package msgcli

import (
	"fmt"
	"strings"

	"github.com/kralamoure/dofus/dofustyp"

	"github.com/kralamoure/d1proto"
)

type ChatSend struct {
	ChatChannel     dofustyp.ChatChannel
	PrivateReceiver string
	Message         string
}

func (m ChatSend) ProtocolId() d1proto.MsgCliId {
	return d1proto.ChatSend
}

func (m ChatSend) Serialized() (string, error) {
	dest := string(m.ChatChannel)
	if m.ChatChannel == dofustyp.ChatChannelPrivate {
		dest = m.PrivateReceiver
	}

	return fmt.Sprintf("%s|%s", dest, m.Message), nil
}

func (m *ChatSend) Deserialize(extra string) error {
	sli := strings.SplitN(extra, "|", 2)
	if len(sli) != 2 {
		return d1proto.ErrInvalidMsg
	}

	if sli[0] == "" {
		return d1proto.ErrInvalidMsg
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
	}

	_, ok := dofustyp.ChatChannels[chatChannel]
	if !ok {
		chatChannel = dofustyp.ChatChannelPrivate
		m.PrivateReceiver = sli[0]
	}

	m.ChatChannel = chatChannel
	m.Message = sli[1]

	return nil
}
