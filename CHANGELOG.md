# Changelog

## [2.6.1](https://github.com/drumandbytes/paddock-api/compare/v2.6.0...v2.6.1) (2026-10-10)


### Bug Fixes

* accept status text as an OpenF1 session result position ([#60](https://github.com/drumandbytes/paddock-api/issues/60)) ([ffcc2ce](https://github.com/drumandbytes/paddock-api/commit/ffcc2cebd83357036ebfd4fdadf27c254e33e6b7))
* **deps:** bump golang.org/x/net to v0.60.0 ([#62](https://github.com/drumandbytes/paddock-api/issues/62)) ([01593d8](https://github.com/drumandbytes/paddock-api/commit/01593d82b70c2c141053b93ba354c006ec2472b2))

## [2.6.0](https://github.com/drumandbytes/paddock-api/compare/v2.5.0...v2.6.0) (2026-09-26)


### Features

* add /healthz liveness endpoint ([#51](https://github.com/drumandbytes/paddock-api/issues/51)) ([fd15f89](https://github.com/drumandbytes/paddock-api/commit/fd15f89e1f07b322aa5f67ce1f1a91f324f8c4d2))

## [2.5.0](https://github.com/drumandbytes/paddock-api/compare/v2.4.0...v2.5.0) (2026-09-25)


### Features

* latest session falls back to the last complete session during OpenF1 lockout ([#49](https://github.com/drumandbytes/paddock-api/issues/49)) ([f813827](https://github.com/drumandbytes/paddock-api/commit/f813827f1582d4a530922c7918336abdfd2dd142))

## [2.4.0](https://github.com/drumandbytes/paddock-api/compare/v2.3.0...v2.4.0) (2026-09-24)


### Features

* durable cache so finished sessions survive restarts ([#44](https://github.com/drumandbytes/paddock-api/issues/44)) ([6f6d0a3](https://github.com/drumandbytes/paddock-api/commit/6f6d0a362d59e38e1d026689dff993946e4eec73))

## [2.3.0](https://github.com/drumandbytes/paddock-api/compare/v2.2.2...v2.3.0) (2026-09-24)


### Features

* driver wins, short constructor names; clearer wins in standings tiles ([#41](https://github.com/drumandbytes/paddock-api/issues/41)) ([1c9f701](https://github.com/drumandbytes/paddock-api/commit/1c9f701121425aee8dd870d267ce460688b6066a))

## [2.2.2](https://github.com/drumandbytes/paddock-api/compare/v2.2.1...v2.2.2) (2026-09-24)


### Bug Fixes

* **widgets:** readable HARD in light mode, explicit tyre stints ([#39](https://github.com/drumandbytes/paddock-api/issues/39)) ([7e59a26](https://github.com/drumandbytes/paddock-api/commit/7e59a2658ca0fc4a1823064877419fcc2778a613))

## [2.2.1](https://github.com/drumandbytes/paddock-api/compare/v2.2.0...v2.2.1) (2026-09-24)


### Bug Fixes

* consistent team names - McLaren, and OpenF1's names matched to the rest ([#37](https://github.com/drumandbytes/paddock-api/issues/37)) ([42b2f21](https://github.com/drumandbytes/paddock-api/commit/42b2f2163f809e08334ef7abeeb001c73c947631))

## [2.2.0](https://github.com/drumandbytes/paddock-api/compare/v2.1.2...v2.2.0) (2026-09-24)


### Features

* latest-session endpoint/widget, hand Next Race over when the race ends ([#33](https://github.com/drumandbytes/paddock-api/issues/33)) ([451fd52](https://github.com/drumandbytes/paddock-api/commit/451fd52dde3734ad609995cb8581b269106aa691))

## [2.1.2](https://github.com/drumandbytes/paddock-api/compare/v2.1.1...v2.1.2) (2026-09-24)


### Bug Fixes

* keep tyre usage through OpenF1 lockouts and after the race starts ([#31](https://github.com/drumandbytes/paddock-api/issues/31)) ([0c61e2f](https://github.com/drumandbytes/paddock-api/commit/0c61e2f982a09215ba4ff7b263ded536173a2bc3))

## [2.1.1](https://github.com/drumandbytes/paddock-api/compare/v2.1.0...v2.1.1) (2026-09-23)


### Bug Fixes

* port the better-built tyre-usage widget, fix a dead docstring ref ([#27](https://github.com/drumandbytes/paddock-api/issues/27)) ([c06a43c](https://github.com/drumandbytes/paddock-api/commit/c06a43cdb0ed84d53b0eff2812c6c2dc2946f616))

## [2.1.0](https://github.com/drumandbytes/paddock-api/compare/v2.0.0...v2.1.0) (2026-09-22)


### Features

* expose last_race url/country, wire unused fields into widgets ([#25](https://github.com/drumandbytes/paddock-api/issues/25)) ([aef3965](https://github.com/drumandbytes/paddock-api/commit/aef396543f8c985b1b7fe0fea9b78dd9db988de8))

## [2.0.0](https://github.com/drumandbytes/paddock-api/compare/v1.2.2...v2.0.0) (2026-09-22)


### ⚠ BREAKING CHANGES

* replace Python implementation with Go ([#22](https://github.com/drumandbytes/paddock-api/issues/22))

### Features

* replace Python implementation with Go ([#22](https://github.com/drumandbytes/paddock-api/issues/22)) ([2c60623](https://github.com/drumandbytes/paddock-api/commit/2c606230128c10e15840602a3bf2d40e5ee8fb87))

## [1.2.2](https://github.com/drumandbytes/paddock-api/compare/v1.2.1...v1.2.2) (2026-09-21)


### Performance Improvements

* **api:** fix cache reads, move blocking Ergast/FastF1 calls to threadpool ([#20](https://github.com/drumandbytes/paddock-api/issues/20)) ([7a07257](https://github.com/drumandbytes/paddock-api/commit/7a07257dd24106f783e5429a71e0d7e9d5e8de69))

## [1.2.1](https://github.com/drumandbytes/paddock-api/compare/v1.2.0...v1.2.1) (2026-09-21)


### Bug Fixes

* track map stroke width and outline visibility ([#15](https://github.com/drumandbytes/paddock-api/issues/15)) ([34e5d90](https://github.com/drumandbytes/paddock-api/commit/34e5d9012bc499639128b9ee87a08dd408f9ee19))

## [1.2.0](https://github.com/drumandbytes/paddock-api/compare/v1.1.0...v1.2.0) (2026-09-20)


### Features

* add tyre usage per session for the current race weekend ([#7](https://github.com/drumandbytes/paddock-api/issues/7)) ([9fcf5f5](https://github.com/drumandbytes/paddock-api/commit/9fcf5f5b2a5b27963be646d7b234e7cb4d4daea9))
* drop f1api.dev, use fastf1/Ergast for everything ([c6adc62](https://github.com/drumandbytes/paddock-api/commit/c6adc628c1aa45835239e4ef45cfdfba01bd0fde))


### Bug Fixes

* remove dead Lap Record/Length rows from the Next Race widgets ([#13](https://github.com/drumandbytes/paddock-api/issues/13)) ([99f76c6](https://github.com/drumandbytes/paddock-api/commit/99f76c6758934e03f32e4bd13726c2a3e434fc56))
* stop retrying after a rate limit, don't burn through the whole calendar ([c6adc62](https://github.com/drumandbytes/paddock-api/commit/c6adc628c1aa45835239e4ef45cfdfba01bd0fde))


### Performance Improvements

* select tyre_usage's needed columns before the groupby ([#14](https://github.com/drumandbytes/paddock-api/issues/14)) ([69d81b0](https://github.com/drumandbytes/paddock-api/commit/69d81b0171549cba5bf144eb2870d815f52dbece))
