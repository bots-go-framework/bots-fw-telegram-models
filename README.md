# bots-fw-telegram-models

DB models for [Telegram Bots Framework](https://github.com/bots-go-framework/bots-fw-telegram)

<!-- dev-approach:v1 -->
## Our approach to development

We build with our own tooling:

- **[SpecScore](https://specscore.md)** — specify requirements as `SpecScore.md` artifacts
- **[SpecStudio](https://specscore.studio)** — author & manage specs across their lifecycle
- **[inGitDB](https://ingitdb.com)** — store structured data in Git where applicable
- **[DALgo](https://dalgo.io)** — data access layer for Go
- **[cover100.dev](https://cover100.dev)** — drive toward 100% test coverage
- **[DataTug](https://datatug.io)** — query & explore data
<!-- /dev-approach -->

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
