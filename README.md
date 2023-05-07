# bots-fw-telegram-models

DB models for [Telegram Bots Framework](https://github.com/bots-go-framework/bots-fw-telegram)

## Dependants:

- [`github.com/bots-go-framework/bots-fw`](https://github.com/bots-go-framework/bots-fw)
- [`github.com/bots-go-framework/bots-fw-telegram`](https://github.com/bots-go-framework/bots-fw-telegram)

## Interfaces

- [`TgChatData`](botsfwtgmodels/chat.go) - interface for Telegram chat data
- [`TgChatInstanceData`](botsfwtgmodels/chat_instance.go) - interface for Telegram chat instance data

## Structs

- [`TgBotUserData`](botsfwtgmodels/bot_user.go) - struct for Telegram user data - implements `botsfw.BotUser`
- [`TgChatBase`](botsfwtgmodels/chat_base.go) - base struct for Telegram chat data.
  Can be used directly or embedded in other structs.
- [`TgChatInstanceBaseData`](botsfwtgmodels/chat_instance.go) - base struct for Telegram chat instance data.
  Can be used directly or embedded in other structs.
