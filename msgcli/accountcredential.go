package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type AccountCredential struct {
	Username     string
	Hash         string
	CryptoMethod int
}

func (m AccountCredential) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountCredential
}

func (m AccountCredential) Serialized() (string, error) {
	return fmt.Sprintf("%s\n#%d%s", m.Username, m.CryptoMethod, m.Hash), nil
}

func (m *AccountCredential) Deserialize(extra string) error {
	sli := strings.Split(extra, "\n")

	if len(sli) < 2 {
		return retroproto.ErrInvalidMsg
	}

	m.Username = sli[0]

	if len(sli[1]) < 3 {
		return retroproto.ErrInvalidMsg
	}

	if sli[1][0] != '#' {
		return retroproto.ErrInvalidMsg
	}

	cryptoMethod := sli[1][1]
	cryptoMethodN, err := strconv.ParseInt(string(cryptoMethod), 10, 32)
	if err != nil {
		return err
	}
	m.CryptoMethod = int(cryptoMethodN)

	m.Hash = sli[1][2:]

	return nil
}
