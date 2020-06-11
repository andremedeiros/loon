catalog_files := $(wildcard catalog/nix/*/*.json)

loon: catalog/bindata.go
	go build

catalog/bindata.go: $(catalog_files)
	go-bindata -pkg catalog -prefix catalog/ -o ./catalog/bindata.go $(catalog_files)

test:
	go test ./... -v
