package msgsvr

import (
	"fmt"
	"strings"

	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/enum"
	"github.com/kralamoure/d1proto/typ"
)

type GameActions struct {
	ActionType          int
	ActionMovement      GameActionsActionMovement
	ActionLoadGameMap   GameActionsActionLoadGameMap
	ActionChallenge     GameActionsActionChallenge
	ActionChallengeJoin GameActionsActionChallengeJoin
}

type GameActionsActionMovement struct {
	Id          int
	SpriteId    int
	DirAndCells []typ.CommonDirAndCell
}

type GameActionsActionLoadGameMap struct {
	SpriteId  int
	Cinematic int
}

type GameActionsActionChallenge struct {
	ChallengerId int
	ChallengedId int
}

type GameActionsActionChallengeJoin struct {
	ChallengedId int
	ErrorReason  rune
}

func (m GameActions) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameActions
}

func (m GameActions) Serialized() (string, error) {
	sb := strings.Builder{}

	id := ""

	switch m.ActionType {
	case enum.GameActionType.Movement:
		id = fmt.Sprintf("%d", m.ActionMovement.Id)

		sb.WriteString(fmt.Sprintf(";%d;", m.ActionMovement.SpriteId))

		for _, v := range m.ActionMovement.DirAndCells {
			n1 := v.DirId & 7
			n2 := (v.CellId & 4032) >> 6
			n3 := v.CellId & 63

			ch1, err := d1proto.Encode64(n1)
			if err != nil {
				return "", err
			}
			ch2, err := d1proto.Encode64(n2)
			if err != nil {
				return "", err
			}
			ch3, err := d1proto.Encode64(n3)
			if err != nil {
				return "", err
			}

			sb.WriteString(fmt.Sprintf("%s%s%s", string(ch1), string(ch2), string(ch3)))
		}
	case enum.GameActionType.LoadGameMap:
		cinematic := ""
		if m.ActionLoadGameMap.Cinematic != 0 {
			cinematic = fmt.Sprintf("%d", m.ActionLoadGameMap.Cinematic)
		}

		sb.WriteString(fmt.Sprintf(";%d;%s", m.ActionLoadGameMap.SpriteId, cinematic))
	case enum.GameActionType.Challenge:
		sb.WriteString(fmt.Sprintf(";%d;%d", m.ActionChallenge.ChallengerId, m.ActionChallenge.ChallengedId))
	case enum.GameActionType.ChallengeJoin:
		sb.WriteString(fmt.Sprintf(";%d;%s", m.ActionChallengeJoin.ChallengedId, string(m.ActionChallengeJoin.ErrorReason)))
	}

	return fmt.Sprintf("%s;%d%s", id, m.ActionType, sb.String()), nil
}

func (m *GameActions) Deserialize(extra string) error {
	return nil
}
