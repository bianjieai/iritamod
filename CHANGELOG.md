<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [Unreleased]

## [v1.2.0] - 2021-12-31

### Application

- (modules/perm) [#33]  Add EVM contract permission management

## [v1.1.1] - 2021-12-07
### Improvements
- (modules/identity) [#32] add `data` field, and the field length limit is only related to the block and transaction size limit.

## [v1.1.0] - 2021-10-27
### Bug Fixes
- [#30] Bump Cosmos-SDK to v0.44.2

## [v1.0.0] - 2021-04-13

### Features

- Add modules `perm`, `params`, `node`, `slashing`, `identity`, `upgrade`.

<!-- Release links -->
[v1.2.0]: https://github.com/bianjieai/iritamod/releases/tag/v1.2.0
[v1.1.1]: https://github.com/bianjieai/iritamod/releases/tag/v1.1.1
[v1.1.0]: https://github.com/bianjieai/iritamod/releases/tag/v1.1.0
[v1.0.0]: https://github.com/bianjieai/iritamod/releases/tag/v1.0.0

<!-- Pull request links -->
[#33]: https://github.com/bianjieai/iritamod/pull/33
[#32]: https://github.com/bianjieai/iritamod/pull/32
[#30]: https://github.com/bianjieai/iritamod/pull/30
