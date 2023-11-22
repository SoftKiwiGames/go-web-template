.PHONY: dev

.PHONY: dev
dev:
	air

.PHONY: gen
gen:
	templ generate

.PHONY: install
install:
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest

.PHONY: changelog
changelog:
	git-chglog -o CHANGELOG.md --next-tag ${TAG}
