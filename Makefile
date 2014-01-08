VERSION = 0.2.0

TEMPLATES_DIR = templates
TEMPLATES_FILE = $(TEMPLATES_DIR)/templates.go

GO_BINDATA = go-bindata
GO_FMT = gofmt -w
GO_XC = goxc

GOXC_FILE = .goxc.local.json

func_name_from_file = $(subst .,zz,$(subst /,z,$(1:$(TEMPLATES_DIR)/%=%)))
file_name_from_func = $(addprefix $(TEMPLATES_DIR)/,$(subst z,/,$(subst zz,.,$(1))))
remove_tmpl_dir = $(1:$(TEMPLATES_DIR)/%=%)

SOURCES = $(filter-out %.go,$(shell find $(TEMPLATES_DIR)/* -type f))
OBJECTS = $(addsuffix .go, $(foreach src, $(SOURCES), $(call func_name_from_file,$(src))))

DEPS = \
	github.com/jteeuwen/go-bindata \
	github.com/robertkrimen/terst \
	github.com/jessevdk/go-flags \
	bitbucket.org/pkg/inflect

all:deps templates

templates:clean ${TEMPLATES_FILE} ${OBJECTS}
	$(shell echo '}' >> $(TEMPLATES_FILE))
	$(GO_FMT) $(TEMPLATES_DIR)/*.go

%.go:
	${GO_BINDATA} -pkg="templates" -out="$(TEMPLATES_DIR)/$@" -func="$(*)" $(call file_name_from_func,$*)
	$(shell echo '\t"$(call remove_tmpl_dir,$(call file_name_from_func,$*))" : $(*:($TEMPLATES_DIR)/%=%),' >> $(TEMPLATES_FILE))

${TEMPLATES_FILE}:
	$(shell echo "package templates\nvar Data = map[string] func() []byte{" > $(TEMPLATES_FILE))

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
	go install github.com/pksunkara/alpaca

deps:
	go get -u $(DEPS)

clean:
	rm -f ${TEMPLATES_DIR}/*.go
