# zaproxy_fix


![](https://github.com/actions/starter-workflows/blob/master/icons/go.svg)

[![GitHub Actions](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fatrox%2Fsync-dotenv%2Fbadge&label=build&logo=none)](https://actions-badge.atrox.dev/atrox/sync-dotenv/goto)

Grabs the latest geckodriver and drops it into the zaproxy folder. It is designed to address [this](https://github.com/zaproxy/zaproxy/issues/5434) issue.

It only currently works for 64 bit linux systems.

Requires bash, tar, mv, and wget.
