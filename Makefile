catalog_files := $(wildcard internal/catalog/*/*.json)

loon: catalog/bindata.go
	go build

catalog/bindata.go: $(catalog_files)
	go-bindata -pkg catalog -prefix internal/catalog/ -o ./internal/catalog/bindata.go $(catalog_files)

test:
	go test ./... -v
