package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type AccountStatsStat struct {
	Base      int
	Equipment int
	Feats     int
	Boost     int
}

func (m AccountStatsStat) Serialized() (string, error) {
	return fmt.Sprintf("%d,%d,%d,%d", m.Base, m.Equipment, m.Feats, m.Boost), nil
}

func (m *AccountStatsStat) Deserialize(extra string) error {
	sli := strings.Split(extra, ",")
	if len(sli) < 4 {
		return d1proto.ErrInvalidMsg
	}

	if sli[0] != "" {
		base, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Base = int(base)
	}

	if sli[1] != "" {
		equipment, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.Equipment = int(equipment)
	}

	if sli[2] != "" {
		feats, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.Feats = int(feats)
	}

	if sli[3] != "" {
		boost, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		m.Boost = int(boost)
	}

	return nil
}
