package botsfwtgmodels

import (
	"github.com/bots-go-framework/bots-fw-store/botsfwmodels"
	"github.com/strongo/strongoapp/person"
)

type TgPlatformUser interface {
	botsfwmodels.PlatformUserData
	TgPlatformUserBaseDbo() *TgPlatformUserBaseDbo
}

var _ TgPlatformUser = (*TgPlatformUserBaseDbo)(nil)

// TgPlatformUserBaseDbo is Telegram user DB TgChatData (without ID)
type TgPlatformUserBaseDbo struct {
	botsfwmodels.PlatformUserBaseDbo
	//TgChatID int64
}

func (entity *TgPlatformUserBaseDbo) TgPlatformUserBaseDbo() *TgPlatformUserBaseDbo {
	return entity
}

//var _ user.AccountData = (*TgPlatformUserBaseDbo)(nil)

//// TgUser is Telegram user DB record (with ID)
//type TgUser struct {
//	record.WithID[int64]
//	Data *TgPlatformUserBaseDbo
//}

// GetEmail returns empty string
func (*TgPlatformUserBaseDbo) GetEmail() string {
	return ""
}

// Name returns full display name combined from (first+last, nick) name
func (entity *TgPlatformUserBaseDbo) Name() string {
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
func (entity *TgPlatformUserBaseDbo) GetNames() person.NameFields {
	return person.NameFields{
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		UserName:  entity.UserName,
	}
}

// IsEmailConfirmed returns false
func (entity *TgPlatformUserBaseDbo) IsEmailConfirmed() bool {
	return false
}

//// Load is for datastore
//func (entity *TgPlatformUserBaseDbo) Load(ps []datastore.Property) error {
//	return datastore.LoadStruct(entity, ps)
//}
//
//// Save is for datastore
//func (entity *TgPlatformUserBaseDbo) Save() (properties []datastore.Property, err error) {
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
