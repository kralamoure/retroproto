package msgsvr

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/enum"
	"github.com/kralamoure/d1proto/typ"
)

type GameMovement struct {
	Sprites []GameMovementSprite
}

type GameMovementSprite struct {
	Transition       bool
	Fight            bool
	Type             int
	Id               int
	CellId           int
	Direction        int
	Creature         GameMovementCreature
	Monster          GameMovementMonster
	MonsterGroup     GameMovementMonsterGroup
	NPC              GameMovementNPC
	OfflineCharacter GameMovementOfflineCharacter
	TaxCollector     GameMovementTaxCollector
	Mutant           GameMovementMutant
	MutantPlayer     GameMovementMutantPlayer
	ParkMount        GameMovementParkMount
	Prism            GameMovementPrism
	Character        GameMovementCharacter
}

type GameMovementCreature struct {
	TemplateId  int
	GFXId       int
	ScaleX      int
	ScaleY      int
	NoFlip      bool
	PowerLevel  int
	Color1      string
	Color2      string
	Color3      string
	Accessories typ.CommonAccessories
	LP          int
	AP          int
	MP          int
	Resistances typ.CommonResistances
	Team        int
}

type GameMovementMonster struct {
	TemplateId  int
	GFXId       int
	ScaleX      int
	ScaleY      int
	NoFlip      bool
	PowerLevel  int
	Color1      string
	Color2      string
	Color3      string
	Accessories typ.CommonAccessories
	LP          int
	AP          int
	MP          int
	Team        int
}

type GameMovementMonsterGroup struct {
	NoFlip     bool
	BonusValue int
	Monsters   []GameMovementMonsterGroupMonster
}

type GameMovementMonsterGroupMonster struct {
	TemplateId  int
	GFXId       int
	ScaleX      int
	ScaleY      int
	Level       int
	Color1      string
	Color2      string
	Color3      string
	Accessories typ.CommonAccessories
}

type GameMovementNPC struct {
	TemplateId    int
	GFXId         int
	Sex           int
	ScaleX        int
	ScaleY        int
	Color1        string
	Color2        string
	Color3        string
	Accessories   typ.CommonAccessories
	ExtraClipId   int
	CustomArtwork int
}

type GameMovementOfflineCharacter struct {
	Name        string
	GFXId       int
	ScaleX      int
	ScaleY      int
	Color1      string
	Color2      string
	Color3      string
	Accessories typ.CommonAccessories
	GuildName   string
	GuildEmblem typ.CommonGuildEmblem
	OfflineType int
}

type GameMovementTaxCollector struct {
	Name        string
	GFXId       int
	ScaleX      int
	ScaleY      int
	Level       int
	GuildName   string
	GuildEmblem typ.CommonGuildEmblem
	LP          int
	AP          int
	MP          int
	Resistances typ.CommonResistances
	Team        int
}

type GameMovementMutant struct {
	TemplateId   int
	GFXId        int
	Sex          int
	ScaleX       int
	ScaleY       int
	PowerLevel   int
	Accessories  typ.CommonAccessories
	Emote        int
	EmoteTimer   time.Duration
	Restrictions typ.CommonRestrictions
	LP           int
	AP           int
	MP           int
	Team         int
}

type GameMovementMutantPlayer struct {
	TemplateId   int
	PlayerName   string
	GFXId        int
	Sex          int
	ScaleX       int
	ScaleY       int
	PowerLevel   int
	Accessories  typ.CommonAccessories
	Emote        int
	EmoteTimer   time.Duration
	Restrictions typ.CommonRestrictions
	LP           int
	AP           int
	MP           int
	Team         int
}

type GameMovementParkMount struct {
	Name      string
	OwnerName string
	GFXId     int
	ScaleX    int
	ScaleY    int
	ModelId   int // it seems in this message ModelId is equal to ScaleX || ScaleY, when ScaleX == ScaleY
	Level     int
}

type GameMovementPrism struct {
	TemplateId     int
	GFXId          int
	ScaleX         int
	ScaleY         int
	Level          int
	AlignmentIndex int
	AlignmentValue int
}

type GameMovementCharacter struct {
	Name                      string
	Title                     typ.CommonTitle
	AllowGhostMode            bool
	GFXId                     int
	Sex                       int
	ScaleX                    int
	ScaleY                    int
	Level                     int
	Color1                    string
	Color2                    string
	Color3                    string
	Accessories               typ.CommonAccessories
	AlignmentId               int
	AlignmentLevel            int
	Grade                     int
	AlignmentFallenAngelDemon bool
	Aura                      int
	Emote                     int
	EmoteTimer                time.Duration
	GuildName                 string
	GuildEmblem               typ.CommonGuildEmblem
	Restrictions              AccountRestrictions
	MountModelId              int
	MountCustomColor1         string
	MountCustomColor2         string
	MountCustomColor3         string
	LP                        int
	AP                        int
	MP                        int
	Resistances               typ.CommonResistances
	Team                      int
	LinkedSprites             GameMovementLinkedSprites
}

type GameMovementLinkedSprites struct {
	Shape   string
	Sprites []GameMovementLinkedSpritesSprite
}

type GameMovementLinkedSpritesSprite struct {
	GFXId  int
	ScaleX int
	ScaleY int
}

type GameMovementGFXData struct {
	Shape string
	GFXs  []string
}

func (m GameMovement) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameMovement
}

func (m GameMovement) Serialized() (string, error) {
	spritesSli := make([]string, len(m.Sprites))

	for i, s := range m.Sprites {
		sb := &strings.Builder{}

		movementType := "+"
		if s.Transition {
			movementType = "~"
		}

		switch s.Type {
		case enum.GameMovementSpriteType.Creature:
			t := s.Creature

			noFlip := ""
			if t.NoFlip {
				noFlip = "*"
			}

			accessories, err := t.Accessories.Serialized()
			if err != nil {
				return "", err
			}

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%d;%d;%d^%s%s;%d;%s;%s;%s;%s",
				movementType, s.CellId, s.Direction, 0, s.Id, t.TemplateId, s.Type, t.GFXId, scale, noFlip,
				t.PowerLevel, t.Color1, t.Color2, t.Color3, accessories,
			))

			if s.Fight {
				sb.WriteString(fmt.Sprintf(";%d;%d;%d", t.LP, t.AP, t.MP))

				// It seems Dofus doesn't actually include resistances for monsters (and maybe neither for creatures),
				// as monsters resistances are already included in lang files.
				resistances, err := t.Resistances.Serialized()
				if err != nil {
					return "", err
				}
				sb.WriteString(fmt.Sprintf(";%s", resistances))

				sb.WriteString(fmt.Sprintf(";%d", t.Team))
			}
		case enum.GameMovementSpriteType.Monster:
			t := s.Monster

			noFlip := ""
			if t.NoFlip {
				noFlip = "*"
			}

			accessories, err := t.Accessories.Serialized()
			if err != nil {
				return "", err
			}

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%d;%d;%d^%s%s;%d;%s;%s;%s;%s",
				movementType, s.CellId, s.Direction, 0, s.Id, t.TemplateId, s.Type, t.GFXId, scale, noFlip,
				t.PowerLevel, t.Color1, t.Color2, t.Color3, accessories,
			))

			if s.Fight {
				sb.WriteString(fmt.Sprintf(";%d;%d;%d", t.LP, t.AP, t.MP))

				// It seems Dofus doesn't actually include resistances for monsters (and maybe neither for creatures),
				// as monsters resistances are already included in lang files.
				/*resistances, err := t.Resistances.Serialized()
				if err != nil {
					return "", err
				}
				sb.WriteString(fmt.Sprintf(";%s", resistances))*/

				sb.WriteString(fmt.Sprintf(";%d", t.Team))
			}
		case enum.GameMovementSpriteType.MonsterGroup:
			t := s.MonsterGroup

			noFlip := ""
			if t.NoFlip {
				noFlip = "*"
			}

			/*sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%d;%d;%d^%dx%d%s;%d;%s;%s;%s;%s",
				movementType, s.CellId, s.Direction, t.BonusValue, s.Id, templateIds, s.Type,
			))*/

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d",
				movementType, s.CellId, s.Direction, t.BonusValue, s.Id,
			))

			templatesId := make([]string, len(s.MonsterGroup.Monsters))
			gfxIdsAndScales := make([]string, len(s.MonsterGroup.Monsters))
			levels := make([]string, len(s.MonsterGroup.Monsters))
			colors := make([]string, len(s.MonsterGroup.Monsters))
			accessories := make([]string, len(s.MonsterGroup.Monsters))
			colorsAndAccessories := make([]string, len(s.MonsterGroup.Monsters))
			for i, monster := range s.MonsterGroup.Monsters {
				templatesId[i] = fmt.Sprintf("%d", monster.TemplateId)

				var scale string
				if monster.ScaleX == monster.ScaleY {
					scale = fmt.Sprint(monster.ScaleX)
				} else {
					scale = fmt.Sprintf("%dx%d", monster.ScaleX, monster.ScaleY)
				}

				gfxIdsAndScales[i] = fmt.Sprintf("%d^%s", monster.GFXId, scale)
				levels[i] = fmt.Sprintf("%d", monster.Level)

				colors[i] = fmt.Sprintf("%s,%s,%s", monster.Color1, monster.Color2, monster.Color3)

				if tmp, err := monster.Accessories.Serialized(); err != nil {
					return "", err
				} else {
					accessories[i] = tmp
				}

				colorsAndAccessories[i] = fmt.Sprintf("%s;%s", colors[i], accessories[i])
			}

			sb.WriteString(fmt.Sprintf(";%s", strings.Join(templatesId, ",")))

			sb.WriteString(fmt.Sprintf(";%d", s.Type))

			sb.WriteString(fmt.Sprintf(";%s", strings.Join(gfxIdsAndScales, ",")))

			sb.WriteString(fmt.Sprintf("%s", noFlip))

			sb.WriteString(fmt.Sprintf(";%s", strings.Join(levels, ",")))

			sb.WriteString(fmt.Sprintf(";%s", strings.Join(colorsAndAccessories, ";")))
		case enum.GameMovementSpriteType.NPC:
			t := s.NPC

			accessories, err := t.Accessories.Serialized()
			if err != nil {
				return "", err
			}

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%d;%d;%d^%s;%d;%s;%s;%s;%s;%d;%d",
				movementType, s.CellId, s.Direction, 0, s.Id, t.TemplateId, s.Type, t.GFXId, scale, t.Sex,
				t.Color1, t.Color2, t.Color3, accessories, t.ExtraClipId, t.CustomArtwork,
			))
		case enum.GameMovementSpriteType.OfflineCharacter:
			t := s.OfflineCharacter

			accessories, err := t.Accessories.Serialized()
			if err != nil {
				return "", err
			}

			guildEmblem, err := t.GuildEmblem.Serialized()
			if err != nil {
				return "", err
			}

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%s;%d;%d^%s;%s;%s;%s;%s;%s;%s;%d",
				movementType, s.CellId, s.Direction, 0, s.Id, t.Name, s.Type, t.GFXId, scale,
				t.Color1, t.Color2, t.Color3, accessories, t.GuildName, guildEmblem, t.OfflineType,
			))
		case enum.GameMovementSpriteType.TaxCollector:
			t := s.TaxCollector

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%s;%d;%d^%s;%d",
				movementType, s.CellId, s.Direction, 0, s.Id, t.Name, s.Type, t.GFXId, scale, t.Level,
			))

			if s.Fight {
				resistances, err := t.Resistances.Serialized()
				if err != nil {
					return "", err
				}

				sb.WriteString(fmt.Sprintf(";%d;%d;%d;%s;%d",
					t.LP, t.AP, t.MP, resistances, t.Team,
				))
			} else {
				guildEmblem, err := t.GuildEmblem.Serialized()
				if err != nil {
					return "", err
				}

				sb.WriteString(fmt.Sprintf(";%s;%s",
					t.GuildName, guildEmblem,
				))
			}
		case enum.GameMovementSpriteType.Mutant:
			t := s.Mutant

			accessories, err := t.Accessories.Serialized()
			if err != nil {
				return "", err
			}

			restrictions, err := t.Restrictions.Serialized()
			if err != nil {
				return "", err
			}

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%d;%d;%d^%s;%d;%d;%s",
				movementType, s.CellId, s.Direction, 0, s.Id, t.TemplateId, s.Type, t.GFXId, scale, t.Sex,
				t.PowerLevel, accessories,
			))

			if s.Fight {
				sb.WriteString(fmt.Sprintf(";%d;%d;%d;;;;;;;;%d",
					t.LP, t.AP, t.MP, t.Team,
				))
			} else {
				sb.WriteString(fmt.Sprintf(";%d;%d;%s",
					t.Emote, t.EmoteTimer.Milliseconds(), restrictions,
				))
			}
		case enum.GameMovementSpriteType.MutantPlayer:
			t := s.MutantPlayer

			accessories, err := t.Accessories.Serialized()
			if err != nil {
				return "", err
			}

			restrictions, err := t.Restrictions.Serialized()
			if err != nil {
				return "", err
			}

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%d~%s;%d;%d^%s;%d;%d;%s",
				movementType, s.CellId, s.Direction, 0, s.Id, t.TemplateId, t.PlayerName, s.Type,
				t.GFXId, scale, t.Sex, t.PowerLevel, accessories,
			))

			if s.Fight {
				sb.WriteString(fmt.Sprintf(";%d;%d;%d;;;;;;;;%d",
					t.LP, t.AP, t.MP, t.Team,
				))
			} else {
				sb.WriteString(fmt.Sprintf(";%d;%d;%s",
					t.Emote, t.EmoteTimer.Milliseconds(), restrictions,
				))
			}
		case enum.GameMovementSpriteType.ParkMount:
			t := s.ParkMount

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%s;%d;%d^%s;%s;%d;%d",
				movementType, s.CellId, s.Direction, 0, s.Id, t.Name, s.Type, t.GFXId, scale, t.OwnerName,
				t.Level, t.ModelId,
			))
		case enum.GameMovementSpriteType.Prism:
			t := s.Prism

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%d;%d;%d^%s;%d;%d;%d",
				movementType, s.CellId, s.Direction, 0, s.Id, t.TemplateId, s.Type, t.GFXId, scale,
				t.Level, t.AlignmentIndex, t.AlignmentValue,
			))
		default:
			t := s.Character

			allowGhostMode := ""
			if !t.AllowGhostMode {
				allowGhostMode = "*"
			}

			accessories, err := t.Accessories.Serialized()
			if err != nil {
				return "", err
			}

			title := ""
			if tmp, err := t.Title.Serialized(); err != nil {
				return "", err
			} else {
				if tmp != "" {
					title = "," + tmp
				}
			}

			gfxIdsAndScales := make([]string, 1+len(s.Character.LinkedSprites.Sprites))

			var scale string
			if t.ScaleX == t.ScaleY {
				scale = fmt.Sprint(t.ScaleX)
			} else {
				scale = fmt.Sprintf("%dx%d", t.ScaleX, t.ScaleY)
			}
			gfxIdsAndScales[0] = fmt.Sprintf("%d^%s", t.GFXId, scale)

			for i, ls := range s.Character.LinkedSprites.Sprites {
				var scale string
				if ls.ScaleX == ls.ScaleY {
					scale = fmt.Sprint(ls.ScaleX)
				} else {
					scale = fmt.Sprintf("%dx%d", ls.ScaleX, ls.ScaleY)
				}
				gfxIdsAndScales[i+1] = fmt.Sprintf("%d^%s", ls.GFXId, scale)
			}

			var gfxIdsAndScalesStr string
			if s.Character.LinkedSprites.Shape == "line" {
				gfxIdsAndScalesStr = strings.Join(gfxIdsAndScales, ":")
			} else {
				gfxIdsAndScalesStr = strings.Join(gfxIdsAndScales, ",")
			}

			sb.WriteString(fmt.Sprintf("%s%d;%d;%d;%d;%s;%d%s;%s%s;%d",
				movementType, s.CellId, s.Direction, 0, s.Id, t.Name, s.Type, title, allowGhostMode, gfxIdsAndScalesStr,
				t.Sex,
			))

			if s.Fight {
				sb.WriteString(fmt.Sprintf(";%d", t.Level))
			}

			alignmentFallenAngelDemon := ""
			if t.AlignmentFallenAngelDemon {
				alignmentFallenAngelDemon = ",1"
			}

			sb.WriteString(fmt.Sprintf(";%d,%d,%d,%d%s;%s;%s;%s;%s",
				t.AlignmentId, t.AlignmentLevel, t.Grade, s.Id+t.Level, alignmentFallenAngelDemon,
				t.Color1, t.Color2, t.Color3, accessories,
			))

			if s.Fight {
				resistances, err := t.Resistances.Serialized()
				if err != nil {
					return "", err
				}

				sb.WriteString(fmt.Sprintf(";%d;%d;%d;%s;%d", t.LP, t.AP, t.MP, resistances, t.Team))
			} else {
				guildEmblem, err := t.GuildEmblem.Serialized()
				if err != nil {
					return "", err
				}

				restrictions, err := t.Restrictions.Serialized()
				if err != nil {
					return "", err
				}

				emote := ""
				emoteTimer := ""
				if t.Emote != 0 {
					emote = fmt.Sprintf("%d", t.Emote)
					emoteTimer = fmt.Sprintf("%d", t.EmoteTimer)
				}

				sb.WriteString(fmt.Sprintf(";%d;%s;%s;%s;%s;%s", t.Aura, emote, emoteTimer,
					t.GuildName, guildEmblem, restrictions,
				))
			}

			mount := ""
			if t.MountModelId != 0 {
				mountSB := &strings.Builder{}

				mountSB.WriteString(fmt.Sprintf("%d", t.MountModelId))

				if !(t.MountCustomColor1 == "" || t.MountCustomColor2 == "" || t.MountCustomColor3 == "") {
					mountSB.WriteString(fmt.Sprintf(",%s,%s,%s",
						t.MountCustomColor1, t.MountCustomColor2, t.MountCustomColor3,
					))
				}

				mount = mountSB.String()
			}
			sb.WriteString(fmt.Sprintf(";%s", mount))
		}

		spritesSli[i] = sb.String()
	}

	return fmt.Sprintf("|%s", strings.Join(spritesSli, "|")), nil
}

func (m *GameMovement) Deserialize(extra string) error {
	spritesData := strings.Split(strings.TrimPrefix(extra, "|"), "|")

	for _, spriteData := range spritesData {
		if len(spriteData) < 2 {
			continue
		}

		sprite := &GameMovementSprite{}

		if spriteData[0] == '~' {
			sprite.Transition = true
		}

		sli := strings.Split(spriteData[1:], ";")

		cellId, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		sprite.CellId = int(cellId)

		direction, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		sprite.Direction = int(direction)

		bonusValue, err := strconv.ParseInt(sli[2], 10, 32) // to use later
		if err != nil {
			return err
		}

		id, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		sprite.Id = int(id)

		spriteTypeAndTitleSli := strings.Split(sli[5], ",")

		spriteType, err := strconv.ParseInt(spriteTypeAndTitleSli[0], 10, 32)
		if err != nil {
			return err
		}
		sprite.Type = int(spriteType)

		var titleId int       // to use later
		var titleParam string // to use later
		if len(spriteTypeAndTitleSli) > 1 {
			titleSli := strings.Split(spriteTypeAndTitleSli[1], "*")

			titleIdN, err := strconv.ParseInt(titleSli[0], 10, 32) // to use later
			if err != nil {
				return err
			}
			titleId = int(titleIdN)

			if len(titleSli) > 1 {
				titleParam = titleSli[1]
			}
		}

		var noFlip bool // to use later
		if strings.HasSuffix(sli[6], "*") {
			sli[6] = strings.TrimSuffix(sli[6], "*")
			noFlip = true
		}

		allowGhostMode := true // to use later
		if strings.HasPrefix(sli[6], "*") {
			sli[6] = strings.TrimPrefix(sli[6], "*")
			allowGhostMode = false
		}

		switch sprite.Type {
		case enum.GameMovementSpriteType.Creature:
			t := &GameMovementCreature{}

			t.NoFlip = noFlip

			templateId, err := strconv.ParseInt(sli[4], 10, 32)
			if err != nil {
				return err
			}
			t.TemplateId = int(templateId)

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			powerLevel, err := strconv.ParseInt(sli[7], 10, 32)
			if err != nil {
				return err
			}
			t.PowerLevel = int(powerLevel)

			t.Color1 = sli[8]
			t.Color2 = sli[9]
			t.Color3 = sli[10]

			if err := t.Accessories.Deserialize(sli[11]); err != nil {
				return err
			}

			if len(sli) >= 15 {
				sprite.Fight = true

				if sli[12] != "" {
					lp, err := strconv.ParseInt(sli[12], 10, 32)
					if err != nil {
						return err
					}
					t.LP = int(lp)
				}
				if sli[13] != "" {
					ap, err := strconv.ParseInt(sli[13], 10, 32)
					if err != nil {
						return err
					}
					t.AP = int(ap)
				}
				if sli[14] != "" {
					mp, err := strconv.ParseInt(sli[14], 10, 32)
					if err != nil {
						return err
					}
					t.MP = int(mp)
				}

				if len(sli) >= 23 {
					if err := t.Resistances.Deserialize(strings.Join(sli[15:22], ";")); err != nil {
						return err
					}

					if sli[22] != "" {
						team, err := strconv.ParseInt(sli[22], 10, 32)
						if err != nil {
							return err
						}
						t.Team = int(team)
					}
				} else {
					if sli[15] != "" {
						team, err := strconv.ParseInt(sli[15], 10, 32)
						if err != nil {
							return err
						}
						t.Team = int(team)
					}
				}
			}

			sprite.Creature = *t
		case enum.GameMovementSpriteType.Monster:
			t := &GameMovementMonster{}

			t.NoFlip = noFlip

			templateId, err := strconv.ParseInt(sli[4], 10, 32)
			if err != nil {
				return err
			}
			t.TemplateId = int(templateId)

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			powerLevel, err := strconv.ParseInt(sli[7], 10, 32)
			if err != nil {
				return err
			}
			t.PowerLevel = int(powerLevel)

			t.Color1 = sli[8]
			t.Color2 = sli[9]
			t.Color3 = sli[10]

			if err := t.Accessories.Deserialize(sli[11]); err != nil {
				return err
			}

			if len(sli) >= 15 {
				sprite.Fight = true

				if sli[12] != "" {
					lp, err := strconv.ParseInt(sli[12], 10, 32)
					if err != nil {
						return err
					}
					t.LP = int(lp)
				}
				if sli[13] != "" {
					ap, err := strconv.ParseInt(sli[13], 10, 32)
					if err != nil {
						return err
					}
					t.AP = int(ap)
				}
				if sli[14] != "" {
					mp, err := strconv.ParseInt(sli[14], 10, 32)
					if err != nil {
						return err
					}
					t.MP = int(mp)
				}

				if len(sli) >= 23 {
					// It seems Dofus doesn't actually include resistances for monsters (and maybe neither for creatures),
					// as monsters resistances are already included in lang files.
					/*if err := t.Resistances.Deserialize(strings.Join(sli[15:22], ";")); err != nil {
						return err
					}
					if sli[22] != "" {
						team, err := strconv.ParseInt(sli[22], 10, 32)
						if err != nil {
							return err
						}
						t.Team = int(team)
					}*/
				} else {
					if sli[15] != "" {
						team, err := strconv.ParseInt(sli[15], 10, 32)
						if err != nil {
							return err
						}
						t.Team = int(team)
					}
				}
			}

			sprite.Monster = *t
		case enum.GameMovementSpriteType.MonsterGroup:
			t := &GameMovementMonsterGroup{}

			t.NoFlip = noFlip
			t.BonusValue = int(bonusValue)

			templateIds := strings.Split(sli[4], ",")
			gfxIdsAndScales := strings.Split(sli[6], ",")
			levels := strings.Split(sli[7], ",")

			t.Monsters = make([]GameMovementMonsterGroupMonster, len(templateIds))
			for i := 0; i < len(templateIds); i++ {
				monster := &GameMovementMonsterGroupMonster{}

				templateId, err := strconv.ParseInt(templateIds[i], 10, 32)
				if err != nil {
					return err
				}
				monster.TemplateId = int(templateId)

				level, err := strconv.ParseInt(levels[i], 10, 32)
				if err != nil {
					return err
				}
				monster.Level = int(level)

				gfxId, scaleX, scaleY, err := splitGFXForScale(gfxIdsAndScales[i])
				if err != nil {
					return err
				}
				monster.GFXId = gfxId
				monster.ScaleX = scaleX
				monster.ScaleY = scaleY

				colors := strings.Split(sli[8+2*i], ",")

				monster.Color1 = colors[0]
				monster.Color2 = colors[1]
				monster.Color3 = colors[2]

				accessories := sli[9+2*i]
				if err := monster.Accessories.Deserialize(accessories); err != nil {
					return err
				}

				t.Monsters[i] = *monster
			}

			sprite.MonsterGroup = *t
		case enum.GameMovementSpriteType.NPC:
			t := &GameMovementNPC{}

			templateId, err := strconv.ParseInt(sli[4], 10, 32)
			if err != nil {
				return err
			}
			t.TemplateId = int(templateId)

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			sex, err := strconv.ParseInt(sli[7], 10, 32)
			if err != nil {
				return err
			}
			t.Sex = int(sex)

			t.Color1 = sli[8]
			t.Color2 = sli[9]
			t.Color3 = sli[10]

			if err := t.Accessories.Deserialize(sli[11]); err != nil {
				return err
			}

			if sli[12] != "" {
				extraClipIdN, err := strconv.ParseInt(sli[12], 10, 32)
				if err != nil {
					return err
				}
				t.ExtraClipId = int(extraClipIdN)
			}

			customArtworkId, err := strconv.ParseInt(sli[13], 10, 32)
			if err != nil {
				return err
			}
			t.CustomArtwork = int(customArtworkId)

			sprite.NPC = *t
		case enum.GameMovementSpriteType.OfflineCharacter:
			t := &GameMovementOfflineCharacter{}

			t.Name = sli[4]

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			t.Color1 = sli[7]
			t.Color2 = sli[8]
			t.Color3 = sli[9]

			if err := t.Accessories.Deserialize(sli[10]); err != nil {
				return err
			}

			t.GuildName = sli[11]

			if sli[12] != "" {
				if err := t.GuildEmblem.Deserialize(sli[12]); err != nil {
					return err
				}
			}

			offlineType, err := strconv.ParseInt(sli[13], 10, 32)
			if err != nil {
				return err
			}
			t.OfflineType = int(offlineType)

			sprite.OfflineCharacter = *t
		case enum.GameMovementSpriteType.TaxCollector:
			t := &GameMovementTaxCollector{}

			t.Name = sli[4]

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			level, err := strconv.ParseInt(sli[7], 10, 32)
			if err != nil {
				return err
			}
			t.Level = int(level)

			if len(sli) >= 19 {
				sprite.Fight = true

				if sli[8] != "" {
					lp, err := strconv.ParseInt(sli[8], 10, 32)
					if err != nil {
						return err
					}
					t.LP = int(lp)
				}
				if sli[9] != "" {
					ap, err := strconv.ParseInt(sli[9], 10, 32)
					if err != nil {
						return err
					}
					t.AP = int(ap)
				}
				if sli[10] != "" {
					mp, err := strconv.ParseInt(sli[10], 10, 32)
					if err != nil {
						return err
					}
					t.MP = int(mp)
				}

				if err := t.Resistances.Deserialize(strings.Join(sli[11:18], ";")); err != nil {
					return err
				}

				if sli[18] != "" {
					team, err := strconv.ParseInt(sli[18], 10, 32)
					if err != nil {
						return err
					}
					t.Team = int(team)
				}
			} else {
				t.GuildName = sli[8]

				if sli[9] != "" {
					if err := t.GuildEmblem.Deserialize(sli[9]); err != nil {
						return err
					}
				}
			}

			sprite.TaxCollector = *t
		case enum.GameMovementSpriteType.Mutant:
			t := &GameMovementMutant{}

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			sex, err := strconv.ParseInt(sli[7], 10, 32)
			if err != nil {
				return err
			}
			t.Sex = int(sex)

			powerLevel, err := strconv.ParseInt(sli[8], 10, 32)
			if err != nil {
				return err
			}
			t.PowerLevel = int(powerLevel)

			if err := t.Accessories.Deserialize(sli[9]); err != nil {
				return err
			}

			if len(sli) >= 21 {
				sprite.Fight = true

				if sli[10] != "" {
					lp, err := strconv.ParseInt(sli[10], 10, 32)
					if err != nil {
						return err
					}
					t.LP = int(lp)
				}
				if sli[11] != "" {
					ap, err := strconv.ParseInt(sli[11], 10, 32)
					if err != nil {
						return err
					}
					t.AP = int(ap)
				}
				if sli[12] != "" {
					mp, err := strconv.ParseInt(sli[12], 10, 32)
					if err != nil {
						return err
					}
					t.MP = int(mp)
				}

				if sli[20] != "" {
					team, err := strconv.ParseInt(sli[20], 10, 32)
					if err != nil {
						return err
					}
					t.Team = int(team)
				}
			} else {
				if sli[10] != "" {
					emote, err := strconv.ParseInt(sli[10], 10, 32)
					if err != nil {
						return err
					}
					t.Emote = int(emote)
				}

				if sli[11] != "" {
					emoteTimer, err := time.ParseDuration(fmt.Sprintf("%sms", sli[11]))
					if err != nil {
						return err
					}
					t.EmoteTimer = emoteTimer
				}

				if err := t.Restrictions.Deserialize(sli[12]); err != nil {
					return err
				}
			}

			templateId, err := strconv.ParseInt(sli[4], 10, 32)
			if err != nil {
				return err
			}
			t.TemplateId = int(templateId)

			sprite.Mutant = *t
		case enum.GameMovementSpriteType.MutantPlayer:
			t := &GameMovementMutantPlayer{}

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			sex, err := strconv.ParseInt(sli[7], 10, 32)
			if err != nil {
				return err
			}
			t.Sex = int(sex)

			powerLevel, err := strconv.ParseInt(sli[8], 10, 32)
			if err != nil {
				return err
			}
			t.PowerLevel = int(powerLevel)

			if err := t.Accessories.Deserialize(sli[9]); err != nil {
				return err
			}

			if len(sli) >= 21 {
				sprite.Fight = true

				if sli[10] != "" {
					lp, err := strconv.ParseInt(sli[10], 10, 32)
					if err != nil {
						return err
					}
					t.LP = int(lp)
				}
				if sli[11] != "" {
					ap, err := strconv.ParseInt(sli[11], 10, 32)
					if err != nil {
						return err
					}
					t.AP = int(ap)
				}
				if sli[12] != "" {
					mp, err := strconv.ParseInt(sli[12], 10, 32)
					if err != nil {
						return err
					}
					t.MP = int(mp)
				}

				if sli[20] != "" {
					team, err := strconv.ParseInt(sli[20], 10, 32)
					if err != nil {
						return err
					}
					t.Team = int(team)
				}
			} else {
				if sli[10] != "" {
					emote, err := strconv.ParseInt(sli[10], 10, 32)
					if err != nil {
						return err
					}
					t.Emote = int(emote)
				}

				if sli[11] != "" {
					emoteTimer, err := time.ParseDuration(fmt.Sprintf("%sms", sli[11]))
					if err != nil {
						return err
					}
					t.EmoteTimer = emoteTimer
				}

				if err := t.Restrictions.Deserialize(sli[12]); err != nil {
					return err
				}
			}

			templateIdAndNameSli := strings.Split(sli[4], "~")

			if len(templateIdAndNameSli) < 2 {
				return d1proto.ErrInvalidMsg
			}

			if templateIdAndNameSli[0] != "" {
				templateId, err := strconv.ParseInt(templateIdAndNameSli[0], 10, 32)
				if err != nil {
					return err
				}
				t.TemplateId = int(templateId)
			}

			t.PlayerName = templateIdAndNameSli[1]

			sprite.MutantPlayer = *t
		case enum.GameMovementSpriteType.ParkMount:
			t := &GameMovementParkMount{}

			t.Name = sli[4]
			t.OwnerName = sli[7]

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			level, err := strconv.ParseInt(sli[8], 10, 32)
			if err != nil {
				return err
			}
			t.Level = int(level)

			modelId, err := strconv.ParseInt(sli[9], 10, 32)
			if err != nil {
				return err
			}
			t.ModelId = int(modelId)

			sprite.ParkMount = *t
		case enum.GameMovementSpriteType.Prism:
			t := &GameMovementPrism{}

			templateId, err := strconv.ParseInt(sli[4], 10, 32)
			if err != nil {
				return err
			}
			t.TemplateId = int(templateId)

			gfxId, scaleX, scaleY, err := splitGFXForScale(sli[6])
			if err != nil {
				return err
			}
			t.GFXId = gfxId
			t.ScaleX = scaleX
			t.ScaleY = scaleY

			level, err := strconv.ParseInt(sli[7], 10, 32)
			if err != nil {
				return err
			}
			t.Level = int(level)

			alignmentValue, err := strconv.ParseInt(sli[8], 10, 32)
			if err != nil {
				return err
			}
			t.AlignmentValue = int(alignmentValue)

			alignmentIndex, err := strconv.ParseInt(sli[9], 10, 32)
			if err != nil {
				return err
			}
			t.AlignmentIndex = int(alignmentIndex)

			sprite.Prism = *t
		default:
			t := &GameMovementCharacter{}

			t.Name = sli[4]

			sex, err := strconv.ParseInt(sli[7], 10, 32)
			if err != nil {
				return err
			}
			t.Sex = int(sex)

			var alignment string

			if len(sli) >= 26 {
				// fight
				sprite.Fight = true

				level, err := strconv.ParseInt(sli[8], 10, 32)
				if err != nil {
					return err
				}
				t.Level = int(level)

				alignment = sli[9]

				t.Color1 = sli[10]
				t.Color2 = sli[11]
				t.Color3 = sli[12]

				if err := t.Accessories.Deserialize(sli[13]); err != nil {
					return err
				}

				if sli[14] != "" {
					lp, err := strconv.ParseInt(sli[14], 10, 32)
					if err != nil {
						return err
					}
					t.LP = int(lp)
				}
				if sli[15] != "" {
					ap, err := strconv.ParseInt(sli[15], 10, 32)
					if err != nil {
						return err
					}
					t.AP = int(ap)
				}
				if sli[16] != "" {
					mp, err := strconv.ParseInt(sli[16], 10, 32)
					if err != nil {
						return err
					}
					t.MP = int(mp)
				}

				if err := t.Resistances.Deserialize(strings.Join(sli[17:24], ";")); err != nil {
					return err
				}

				if sli[24] != "" {
					team, err := strconv.ParseInt(sli[24], 10, 32)
					if err != nil {
						return err
					}
					t.Team = int(team)
				}

				if sli[25] != "" {
					if strings.Contains(sli[25], ",") {

						mountSli := strings.Split(sli[25], ",")

						mountModelId, err := strconv.ParseInt(mountSli[0], 10, 32)
						if err != nil {
							return err
						}
						t.MountModelId = int(mountModelId)

						t.MountCustomColor1 = mountSli[1]
						t.MountCustomColor2 = mountSli[2]
						t.MountCustomColor3 = mountSli[3]
					} else {
						mountModelId, err := strconv.ParseInt(sli[25], 10, 32)
						if err != nil {
							return err
						}
						t.MountModelId = int(mountModelId)
					}
				}
			} else {
				// not fight
				alignment = sli[8]

				t.Color1 = sli[9]
				t.Color2 = sli[10]
				t.Color3 = sli[11]

				if err := t.Accessories.Deserialize(sli[12]); err != nil {
					return err
				}

				if sli[13] != "" {
					aura, err := strconv.ParseInt(sli[13], 10, 32)
					if err != nil {
						return err
					}
					t.Aura = int(aura)
				}

				if sli[14] != "" {
					emote, err := strconv.ParseInt(sli[14], 10, 32)
					if err != nil {
						return err
					}
					t.Emote = int(emote)
				}

				if sli[15] != "" {
					emoteTimer, err := time.ParseDuration(fmt.Sprintf("%sms", sli[15]))
					if err != nil {
						return err
					}
					t.EmoteTimer = emoteTimer
				}

				t.GuildName = sli[16]

				if sli[17] != "" {
					if err := t.GuildEmblem.Deserialize(sli[17]); err != nil {
						return err
					}
				}

				if sli[18] != "" {
					if err := t.Restrictions.Deserialize(sli[18]); err != nil {
						return err
					}
				}

				if sli[19] != "" {
					if strings.Contains(sli[19], ",") {

						mountSli := strings.Split(sli[19], ",")

						mountModelId, err := strconv.ParseInt(mountSli[0], 10, 32)
						if err != nil {
							return err
						}
						t.MountModelId = int(mountModelId)

						t.MountCustomColor1 = mountSli[1]
						t.MountCustomColor2 = mountSli[2]
						t.MountCustomColor3 = mountSli[3]
					} else {
						mountModelId, err := strconv.ParseInt(sli[19], 10, 32)
						if err != nil {
							return err
						}
						t.MountModelId = int(mountModelId)
					}
				}
			}

			alignmentSli := strings.Split(alignment, ",")

			alignmentId, err := strconv.ParseInt(alignmentSli[0], 10, 32)
			if err != nil {
				return err
			}
			t.AlignmentId = int(alignmentId)

			alignmentLevel, err := strconv.ParseInt(alignmentSli[1], 10, 32)
			if err != nil {
				return err
			}
			t.AlignmentLevel = int(alignmentLevel)

			grade, err := strconv.ParseInt(alignmentSli[2], 10, 32)
			if err != nil {
				return err
			}
			t.Grade = int(grade)

			if len(alignmentSli) > 3 && alignmentSli[3] != "" {
				idPlusLevel, err := strconv.ParseInt(alignmentSli[3], 10, 32)
				if err != nil {
					return err
				}
				t.Level = int(idPlusLevel) - sprite.Id
			}

			if len(alignmentSli) > 4 && alignmentSli[4] != "" {
				alignmentFallenAngelDemon, err := strconv.ParseBool(alignmentSli[4])
				if err != nil {
					return err
				}
				t.AlignmentFallenAngelDemon = alignmentFallenAngelDemon
			}

			gfxData := splitGFXData(sli[6])
			t.LinkedSprites.Shape = gfxData.Shape

			for i, gfx := range gfxData.GFXs {
				gfxId, scaleX, scaleY, err := splitGFXForScale(gfx)
				if err != nil {
					return err
				}

				if i == 0 {
					t.GFXId = gfxId
					t.ScaleX = scaleX
					t.ScaleY = scaleY
				} else {
					t.LinkedSprites.Sprites = append(t.LinkedSprites.Sprites, GameMovementLinkedSpritesSprite{
						GFXId:  gfxId,
						ScaleX: scaleX,
						ScaleY: scaleY,
					})
				}
			}

			t.Title.Id = titleId
			t.Title.Param = titleParam

			t.AllowGhostMode = allowGhostMode

			sprite.Character = *t
		}

		m.Sprites = append(m.Sprites, *sprite)
	}

	return nil
}

func splitGFXForScale(s string) (gfxId, scaleX, scaleY int, e error) {
	scaleX = 100
	scaleY = 100

	sli := strings.Split(s, "^")

	if len(sli) == 2 {
		gfxIdN, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			e = err
			return
		}
		gfxId = int(gfxIdN)

		if strings.Contains(sli[1], "x") {
			scales := strings.Split(sli[1], "x")
			if len(scales) == 2 {
				scaleXN, err := strconv.ParseInt(scales[0], 10, 32)
				if err != nil {
					e = err
					return
				}
				scaleX = int(scaleXN)

				scaleYN, err := strconv.ParseInt(scales[1], 10, 32)
				if err != nil {
					e = err
					return
				}
				scaleY = int(scaleYN)
			}
		} else {
			scale, err := strconv.ParseInt(sli[1], 10, 32)
			if err != nil {
				e = err
				return
			}
			scaleX = int(scale)
			scaleY = int(scale)
		}
	} else {
		gfxIdN, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			e = err
			return
		}
		gfxId = int(gfxIdN)
	}

	return
}

func splitGFXData(extra string) GameMovementGFXData {
	gfxData := &GameMovementGFXData{}

	if strings.Contains(extra, ",") {
		gfxData.Shape = "circle"
		gfxData.GFXs = strings.Split(extra, ",")
		return *gfxData
	} else if strings.Contains(extra, ":") {
		gfxData.Shape = "line"
		gfxData.GFXs = strings.Split(extra, ":")
	} else {
		gfxData.Shape = "none"
		gfxData.GFXs = []string{extra}
	}

	return *gfxData
}
