package botsfwtgmodels

import (
	"github.com/bots-go-framework/bots-fw-store/botsfwmodels"
	"github.com/strongo/app/user"
)

type TgBotUser interface {
	botsfwmodels.BotUserData
	TgBotUserBaseData() *TgBotUserBaseData
}

var _ TgBotUser = (*TgBotUserBaseData)(nil)

// TgBotUserBaseData is Telegram user DB TgChatData (without ID)
type TgBotUserBaseData struct {
	botsfwmodels.BotUserBaseData
	//TgChatID int64
}

func (entity *TgBotUserBaseData) TgBotUserBaseData() *TgBotUserBaseData {
	return entity
}

//var _ user.AccountData = (*TgBotUserBaseData)(nil)

//// TgUser is Telegram user DB record (with ID)
//type TgUser struct {
//	record.WithID[int64]
//	Data *TgBotUserBaseData
//}

// GetEmail returns empty string
func (*TgBotUserBaseData) GetEmail() string {
	return ""
}

// Name returns full display name cmbined from (first+last, nick) name
func (entity *TgBotUserBaseData) Name() string {
	if entity.FirstName == "" && entity.LastName == "" {
		return "@" + entity.UserName
	}
	name := entity.FirstName
	if name != "" {
		name += " " + entity.LastName
	} else {
		name = entity.LastName
	}
	if entity.UserName == "" {
		return name
	}
	return "@" + entity.UserName + " - " + name
}

// GetNames return user names
func (entity *TgBotUserBaseData) GetNames() user.Names {
	return user.Names{
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		NickName:  entity.UserName,
	}
}

// IsEmailConfirmed returns false
func (entity *TgBotUserBaseData) IsEmailConfirmed() bool {
	return false
}

//// Load is for datastore
//func (entity *TgBotUserBaseData) Load(ps []datastore.Property) error {
//	return datastore.LoadStruct(entity, ps)
//}
//
//// Save is for datastore
//func (entity *TgBotUserBaseData) Save() (properties []datastore.Property, err error) {
//	if properties, err = datastore.SaveStruct(entity); err != nil {
//		return properties, err
//	}
//
//	//if properties, err = gaedb.CleanProperties(properties, map[string]gaedb.IsOkToRemove{
//	//	"AccessGranted": gaedb.IsFalse,
//	//	"FirstName":     gaedb.IsEmptyString,
//	//	"LastName":      gaedb.IsEmptyString,
//	//	"UserName":      gaedb.IsEmptyString,
//	//}); err != nil {
//	//	return
//	//}
//
//	return
//}
