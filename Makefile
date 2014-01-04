TEMPLATES_DIR = templates
TEMPLATES_FILE = $(TEMPLATES_DIR)/templates.go

GO_BINDATA = go-bindata
GO_FMT = gofmt -w

func_name_from_file = $(subst .,zz,$(subst /,z,$(1:$(TEMPLATES_DIR)/%=%)))
file_name_from_func = $(addprefix $(TEMPLATES_DIR)/,$(subst z,/,$(subst zz,.,$(1))))
remove_tmpl_dir = $(1:$(TEMPLATES_DIR)/%=%)

SOURCES = $(filter-out %.go,$(shell find $(TEMPLATES_DIR)/* -type f))
OBJECTS = $(addsuffix .go, $(foreach src, $(SOURCES), $(call func_name_from_file,$(src))))

all:deps templates

templates:clean ${TEMPLATES_FILE} ${OBJECTS}
	$(shell echo '}' >> $(TEMPLATES_FILE))
	$(GO_FMT) $(TEMPLATES_DIR)/*.go

%.go:
	${GO_BINDATA} -pkg="templates" -out="$(TEMPLATES_DIR)/$@" -func="$(*)" $(call file_name_from_func,$*)
	$(shell echo '\t"$(call remove_tmpl_dir,$(call file_name_from_func,$*))" : $(*:($TEMPLATES_DIR)/%=%),' >> $(TEMPLATES_FILE))

${TEMPLATES_FILE}:
	$(shell echo "package templates\nvar Data = map[string] func() []byte{" > $(TEMPLATES_FILE))

test:
	go test github.com/pksunkara/alpaca/alpaca

install:
	go install github.com/pksunkara/alpaca

deps:
	go get github.com/jteeuwen/go-bindata
	go get github.com/robertkrimen/terst
	go get github.com/jessevdk/go-flags
	go get bitbucket.org/pkg/inflect

clean:
	rm -f ${TEMPLATES_DIR}/*.go
