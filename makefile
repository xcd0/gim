all:
	cd cmd/gim && go build && mv gim ../..

release-no-tag:
	goreleaser release --snapshot --clean
