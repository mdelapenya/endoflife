.PHONY: build
build:
	goreleaser --snapshot --skip-publish --rm-dist
