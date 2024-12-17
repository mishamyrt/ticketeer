# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][],
and this project adheres to [Semantic Versioning][].


## [v0.1.2](https://github.com/mishamyrt/ticketeer/releases/tag/v0.1.2) - 2024-12-18
### Bug Fixes
- check runner script size before read
- avoid unnecessary error throwing
- don't panic

### Features
- improve hook detection, add tandem usage link

### Refactoring
- clean up


## [v0.1.1](https://github.com/mishamyrt/ticketeer/releases/tag/v0.1.1) - 2024-12-17
### Bug Fixes
- correctly unmap
- use correct default templates
- exit with correct code

### Features
- add npm packaging
- add branch ignore handling
- check if message already contains the ticket id


## [v0.1.0](https://github.com/mishamyrt/ticketeer/releases/tag/v0.1.0) - 2024-12-15
### Bug Fixes
- skip if task id is not found

### CI
- add qa workflow

### Features
- add multiple branch format support
- add basic hook installation

### Refactoring
- simplify regexp parsing
- clean up

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
[Unreleased]: https://github.com/mishamyrt/ticketeer/compare/v0.1.2...HEAD
[v0.1.2]: https://github.com/mishamyrt/ticketeer/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/mishamyrt/ticketeer/compare/v0.1.0...v0.1.1
