package botsfwtgmodels

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConstants(t *testing.T) {
	t.Run("BotUserCollection", func(t *testing.T) {
		assert.False(t, strings.TrimSpace(BotUserCollection) == "", "BotUserCollection is empty")
	})
	t.Run("TgChatCollection", func(t *testing.T) {
		assert.False(t, strings.TrimSpace(TgChatCollection) == "", "TgChatCollection is empty")
	})
}
