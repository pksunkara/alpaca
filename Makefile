VERSION = 0.2.0

GO_RICE = rice
GO_FMT = gofmt -w
GO_XC = goxc

GOXC_FILE = .goxc.local.json

DEPS = \
	github.com/GeertJohan/go.rice/rice \
	github.com/robertkrimen/terst \
	github.com/jessevdk/go-flags \
	bitbucket.org/pkg/inflect

all:deps templates

templates:clean
	$(GO_RICE) --import-path github.com/pksunkara/alpaca/alpaca embed

compile:templates goxc

goxc:
	$(shell echo '{\n "ArtifactsDest": "build",\n "ConfigVersion": "0.9",' > $(GOXC_FILE))
	$(shell echo ' "PackageVersion": "$(VERSION)",\n "TaskSettings": {' >> $(GOXC_FILE))
	$(shell echo '  "bintray": {\n   "apikey": "",\n   "package": "alpaca",' >> $(GOXC_FILE))
	$(shell echo '   "repository": "utils",\n   "subject": "pksunkara"' >> $(GOXC_FILE))
	$(shell echo '  }\n }\n}' >> $(GOXC_FILE))
	$(GO_XC)

bintray:
	$(GO_XC) bintray

test:
	go test github.com/pksunkara/alpaca/alpaca

test-cover:
	go test -coverprofile=coverage.out github.com/pksunkara/alpaca/alpaca
	go tool cover -html=coverage.out

install:
	go install -a github.com/pksunkara/alpaca

deps:
	go get -u $(DEPS)

clean:
	$(GO_RICE) --import-path github.com/pksunkara/alpaca/alpaca clean
