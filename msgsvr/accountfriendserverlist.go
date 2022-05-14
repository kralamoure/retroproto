package msgsvr

import (
	"strings"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type AccountFriendServerList struct {
	ServersCharacters []typ.AccountServersListServerCharacters
}

func (m AccountFriendServerList) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountFriendServerList
}

func (m AccountFriendServerList) Serialized() (string, error) {
	if len(m.ServersCharacters) == 0 {
		return "null", nil
	}

	serverCharactersSli := make([]string, len(m.ServersCharacters))
	for i, v := range m.ServersCharacters {
		serverCharacterStr, err := v.Serialized()
		if err != nil {
			return "", err
		}
		serverCharactersSli[i] = serverCharacterStr
	}

	return strings.Join(serverCharactersSli, ";"), nil
}

func (m *AccountFriendServerList) Deserialize(extra string) error {
	if extra == "" {
		return retroproto.ErrInvalidMsg
	}

	if extra == "null" {
		return nil
	}

	characters := strings.Split(extra, ";")

	m.ServersCharacters = make([]typ.AccountServersListServerCharacters, len(characters))
	for i, v := range characters {
		character := &typ.AccountServersListServerCharacters{}
		err := character.Deserialize(v)
		if err != nil {
			return err
		}
		m.ServersCharacters[i] = *character
	}

	return nil
}
