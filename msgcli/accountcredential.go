package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type AccountCredential struct {
	Username     string
	Hash         string
	CryptoMethod int
}

func (m AccountCredential) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountCredential
}

func (m AccountCredential) Serialized() (string, error) {
	return fmt.Sprintf("%s\n#%d%s", m.Username, m.CryptoMethod, m.Hash), nil
}

func (m *AccountCredential) Deserialize(extra string) error {
	sli := strings.Split(extra, "\n")

	if len(sli) < 2 {
		return d1proto.ErrInvalidMsg
	}

	m.Username = sli[0]

	if len(sli[1]) < 3 {
		return d1proto.ErrInvalidMsg
	}

	if sli[1][0] != '#' {
		return d1proto.ErrInvalidMsg
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
