package botsfwtgmodels

// TgChatData must be implemented by a data struct that stores chat data related to specific app/bot.
type TgChatData interface {

	// BaseChatData returns a base struct that should be included in all structs that implement TgChatData.
	BaseTgChatData() *TgChatBase

	GetTgChatInstanceID() string

	SetTgChatInstanceID(v string)

	GetPreferredLanguage() string

	SetPreferredLanguage(code string)
}
