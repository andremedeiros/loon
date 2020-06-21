catalog_service_files := $(wildcard internal/catalog/service/*/*.json)
catalog_language_files := $(wildcard internal/catalog/language/*/*.json)

loon: catalog/language/bindata.go catalog/service/bindata.go
	go build

catalog/service/bindata.go: $(catalog_service_files)
	go-bindata -pkg service \
		-prefix internal/catalog/service/ \
		-o ./internal/catalog/service/bindata.go \
		$(catalog_service_files)

catalog/language/bindata.go: $(catalog_language_files)
	go-bindata -pkg language \
		-prefix internal/catalog/language/ \
		-o ./internal/catalog/language/bindata.go \
		$(catalog_language_files)

test:
	go test ./... -v
