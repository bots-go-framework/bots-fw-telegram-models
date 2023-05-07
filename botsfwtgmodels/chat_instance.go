package botsfwtgmodels

// TgChatInstanceData describes chat instance TgChatData interface
type TgChatInstanceData interface {
	Validate() error
	GetTgChatID() int64
	GetPreferredLanguage() string
	SetPreferredLanguage(v string)
}

var _ TgChatInstanceData = (*TgChatInstanceBaseData)(nil)

// TgChatInstanceBaseData is base struct for storing data related to a Telegram chat instance
type TgChatInstanceBaseData struct {
	TgChatID          int64  `dalgo:",noindex" datastore:",noindex" firestore:",noindex"`
	PreferredLanguage string `dalgo:",noindex" datastore:",noindex" firestore:",noindex"`
}

func (v *TgChatInstanceBaseData) Validate() error {
	return nil
}

// GetTgChatID returns Telegram chat ID
func (v *TgChatInstanceBaseData) GetTgChatID() int64 {
	return v.TgChatID
}

// GetPreferredLanguage returns preferred language for the chat
func (v *TgChatInstanceBaseData) GetPreferredLanguage() string {
	return v.PreferredLanguage
}

// SetPreferredLanguage sets preferred language for the chat
func (v *TgChatInstanceBaseData) SetPreferredLanguage(code string) {
	v.PreferredLanguage = code
}
