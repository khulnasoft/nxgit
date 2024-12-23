DIST := dist
IMPORT := go.khulnasoft.com/nxgit

GO ?= go
SED_INPLACE := sed -i

export PATH := $($(GO) env GOPATH)/bin:$(PATH)

ifeq ($(OS), Windows_NT)
	EXECUTABLE := nxgit.exe
else
	EXECUTABLE := nxgit
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		SED_INPLACE := sed -i ''
	endif
endif

BINDATA := modules/{options,public,templates}/bindata.go
GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*" ! -path "*/bindata.go")
GOFMT ?= gofmt -s

GOFLAGS := -i -v
EXTRA_GOFLAGS ?=

ifneq ($(DRONE_TAG),)
	VERSION ?= $(subst v,,$(DRONE_TAG))
	NXGIT_VERSION := $(VERSION)
else
	ifneq ($(DRONE_BRANCH),)
		VERSION ?= $(subst release/v,,$(DRONE_BRANCH))
	else
		VERSION ?= master
	endif
	NXGIT_VERSION := $(shell git describe --tags --always | sed 's/-/+/' | sed 's/^v//')
endif

LDFLAGS := -X "main.Version=$(NXGIT_VERSION)" -X "main.Tags=$(TAGS)"

PACKAGES ?= $(filter-out go.khulnasoft.com/nxgit/integrations/migration-test,$(filter-out go.khulnasoft.com/nxgit/integrations,$(shell $(GO) list ./... | grep -v /vendor/)))
SOURCES ?= $(shell find . -name "*.go" -type f)

TAGS ?=

TMPDIR := $(shell mktemp -d 2>/dev/null || mktemp -d -t 'nxgit-temp')

SWAGGER_SPEC := templates/swagger/v1_json.tmpl
SWAGGER_SPEC_S_TMPL := s|"basePath":\s*"/api/v1"|"basePath": "{{AppSubUrl}}/api/v1"|g
SWAGGER_SPEC_S_JSON := s|"basePath":\s*"{{AppSubUrl}}/api/v1"|"basePath": "/api/v1"|g

TEST_MYSQL_HOST ?= mysql:3306
TEST_MYSQL_DBNAME ?= testnxgit
TEST_MYSQL_USERNAME ?= root
TEST_MYSQL_PASSWORD ?=
TEST_MYSQL8_HOST ?= mysql8:3306
TEST_MYSQL8_DBNAME ?= testnxgit
TEST_MYSQL8_USERNAME ?= root
TEST_MYSQL8_PASSWORD ?=
TEST_PGSQL_HOST ?= pgsql:5432
TEST_PGSQL_DBNAME ?= testnxgit
TEST_PGSQL_USERNAME ?= postgres
TEST_PGSQL_PASSWORD ?= postgres
TEST_MSSQL_HOST ?= mssql:1433
TEST_MSSQL_DBNAME ?= nxgit
TEST_MSSQL_USERNAME ?= sa
TEST_MSSQL_PASSWORD ?= MwantsaSecurePassword1

ifeq ($(OS), Windows_NT)
	EXECUTABLE := nxgit.exe
else
	EXECUTABLE := nxgit
endif

# $(call strip-suffix,filename)
strip-suffix = $(firstword $(subst ., ,$(1)))

.PHONY: all
all: build

include docker/Makefile

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -rf $(EXECUTABLE) $(DIST) $(BINDATA) \
		integrations*.test \
		integrations/nxgit-integration-pgsql/ integrations/nxgit-integration-mysql/ integrations/nxgit-integration-mysql8/ integrations/nxgit-integration-sqlite/ \
		integrations/nxgit-integration-mssql/ integrations/indexers-mysql/ integrations/indexers-mysql8/ integrations/indexers-pgsql integrations/indexers-sqlite \
		integrations/indexers-mssql integrations/mysql.ini integrations/mysql8.ini integrations/pgsql.ini integrations/mssql.ini

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: vet
vet:
	$(GO) vet $(PACKAGES)

.PHONY: generate
generate:
	@hash go-bindata > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/jteeuwen/go-bindata/go-bindata; \
	fi
	$(GO) generate $(PACKAGES)

.PHONY: generate-swagger
generate-swagger:
	@hash swagger > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/go-swagger/go-swagger/cmd/swagger; \
	fi
	swagger generate spec -o './$(SWAGGER_SPEC)'
	$(SED_INPLACE) '$(SWAGGER_SPEC_S_TMPL)' './$(SWAGGER_SPEC)'

.PHONY: swagger-check
swagger-check: generate-swagger
	@diff=$$(git diff '$(SWAGGER_SPEC)'); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make generate-swagger' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: swagger-validate
swagger-validate:
	@hash swagger > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/go-swagger/go-swagger/cmd/swagger; \
	fi
	$(SED_INPLACE) '$(SWAGGER_SPEC_S_JSON)' './$(SWAGGER_SPEC)'
	swagger validate './$(SWAGGER_SPEC)'
	$(SED_INPLACE) '$(SWAGGER_SPEC_S_TMPL)' './$(SWAGGER_SPEC)'

.PHONY: errcheck
errcheck:
	@hash errcheck > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/kisielk/errcheck; \
	fi
	errcheck $(PACKAGES)

.PHONY: lint
lint:
	@hash revive > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mgechev/revive; \
	fi
	revive -config .revive.toml -exclude=./vendor/... ./... || exit 1

.PHONY: misspell-check
misspell-check:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -error -i unknwon $(GOFILES)

.PHONY: misspell
misspell:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -w -i unknwon $(GOFILES)

.PHONY: fmt-check
fmt-check:
	# get all go files and run go fmt on them
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: test
test:
	$(GO) test -tags='sqlite sqlite_unlock_notify' $(PACKAGES)

.PHONY: coverage
coverage:
	@hash gocovmerge > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/wadey/gocovmerge; \
	fi
	gocovmerge integration.coverage.out $(shell find . -type f -name "coverage.out") > coverage.all;\

.PHONY: unit-test-coverage
unit-test-coverage:
	$(GO) test -tags='sqlite sqlite_unlock_notify' -cover -coverprofile coverage.out $(PACKAGES) && echo "\n==>\033[32m Ok\033[m\n" || exit 1

.PHONY: vendor
vendor:
	@hash dep > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/golang/dep/cmd/dep; \
	fi
	dep ensure -vendor-only

.PHONY: test-vendor
test-vendor: vendor
	@diff=$$(git diff vendor/); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make vendor' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;
#TODO add dep status -missing when implemented

.PHONY: test-sqlite
test-sqlite: integrations.sqlite.test
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/sqlite.ini ./integrations.sqlite.test

.PHONY: test-sqlite-migration
test-sqlite-migration:  migrations.sqlite.test
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/sqlite.ini ./migrations.sqlite.test

generate-ini:
	sed -e 's|{{TEST_MYSQL_HOST}}|${TEST_MYSQL_HOST}|g' \
		-e 's|{{TEST_MYSQL_DBNAME}}|${TEST_MYSQL_DBNAME}|g' \
		-e 's|{{TEST_MYSQL_USERNAME}}|${TEST_MYSQL_USERNAME}|g' \
		-e 's|{{TEST_MYSQL_PASSWORD}}|${TEST_MYSQL_PASSWORD}|g' \
			integrations/mysql.ini.tmpl > integrations/mysql.ini
	sed -e 's|{{TEST_MYSQL8_HOST}}|${TEST_MYSQL8_HOST}|g' \
		-e 's|{{TEST_MYSQL8_DBNAME}}|${TEST_MYSQL8_DBNAME}|g' \
		-e 's|{{TEST_MYSQL8_USERNAME}}|${TEST_MYSQL8_USERNAME}|g' \
		-e 's|{{TEST_MYSQL8_PASSWORD}}|${TEST_MYSQL8_PASSWORD}|g' \
			integrations/mysql8.ini.tmpl > integrations/mysql8.ini
	sed -e 's|{{TEST_PGSQL_HOST}}|${TEST_PGSQL_HOST}|g' \
		-e 's|{{TEST_PGSQL_DBNAME}}|${TEST_PGSQL_DBNAME}|g' \
		-e 's|{{TEST_PGSQL_USERNAME}}|${TEST_PGSQL_USERNAME}|g' \
		-e 's|{{TEST_PGSQL_PASSWORD}}|${TEST_PGSQL_PASSWORD}|g' \
			integrations/pgsql.ini.tmpl > integrations/pgsql.ini
	sed -e 's|{{TEST_MSSQL_HOST}}|${TEST_MSSQL_HOST}|g' \
		-e 's|{{TEST_MSSQL_DBNAME}}|${TEST_MSSQL_DBNAME}|g' \
		-e 's|{{TEST_MSSQL_USERNAME}}|${TEST_MSSQL_USERNAME}|g' \
		-e 's|{{TEST_MSSQL_PASSWORD}}|${TEST_MSSQL_PASSWORD}|g' \
			integrations/mssql.ini.tmpl > integrations/mssql.ini

.PHONY: test-mysql
test-mysql: integrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mysql.ini ./integrations.test

.PHONY: test-mysql-migration
test-mysql-migration: migrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mysql.ini ./migrations.test

.PHONY: test-mysql8
test-mysql8: integrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mysql8.ini ./integrations.test

.PHONY: test-mysql8-migration
test-mysql8-migration: migrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mysql8.ini ./migrations.test

.PHONY: test-pgsql
test-pgsql: integrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/pgsql.ini ./integrations.test

.PHONY: test-pgsql-migration
test-pgsql-migration: migrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/pgsql.ini ./migrations.test

.PHONY: test-mssql
test-mssql: integrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mssql.ini ./integrations.test

.PHONY: test-mssql-migration
test-mssql-migration: migrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mssql.ini ./migrations.test


.PHONY: bench-sqlite
bench-sqlite: integrations.sqlite.test
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/sqlite.ini ./integrations.sqlite.test -test.cpuprofile=cpu.out -test.run DontRunTests -test.bench .

.PHONY: bench-mysql
bench-mysql: integrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mysql.ini ./integrations.test -test.cpuprofile=cpu.out -test.run DontRunTests -test.bench .

.PHONY: bench-mssql
bench-mssql: integrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mssql.ini ./integrations.test -test.cpuprofile=cpu.out -test.run DontRunTests -test.bench .

.PHONY: bench-pgsql
bench-pgsql: integrations.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/pgsql.ini ./integrations.test -test.cpuprofile=cpu.out -test.run DontRunTests -test.bench .


.PHONY: integration-test-coverage
integration-test-coverage: integrations.cover.test generate-ini
	NXGIT_ROOT=${CURDIR} NXGIT_CONF=integrations/mysql.ini ./integrations.cover.test -test.coverprofile=integration.coverage.out

integrations.test: $(SOURCES)
	$(GO) test -c go.khulnasoft.com/nxgit/integrations -o integrations.test

integrations.sqlite.test: $(SOURCES)
	$(GO) test -c go.khulnasoft.com/nxgit/integrations -o integrations.sqlite.test -tags 'sqlite sqlite_unlock_notify'

integrations.cover.test: $(SOURCES)
	$(GO) test -c go.khulnasoft.com/nxgit/integrations -coverpkg $(shell echo $(PACKAGES) | tr ' ' ',') -o integrations.cover.test

.PHONY: migrations.test
migrations.test: $(SOURCES)
	$(GO) test -c go.khulnasoft.com/nxgit/integrations/migration-test -o migrations.test

.PHONY: migrations.sqlite.test
migrations.sqlite.test: $(SOURCES)
	$(GO) test -c go.khulnasoft.com/nxgit/integrations/migration-test -o migrations.sqlite.test -tags 'sqlite sqlite_unlock_notify'

.PHONY: check
check: test

.PHONY: install
install: $(wildcard *.go)
	$(GO) install -v -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)'

.PHONY: build
build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -o $@

.PHONY: release
release: release-dirs release-windows release-linux release-darwin release-copy release-compress release-check

.PHONY: release-dirs
release-dirs:
	mkdir -p $(DIST)/binaries $(DIST)/release

.PHONY: release-windows
release-windows:
	@hash xgo > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u src.techknowlogick.com/xgo; \
	fi
	xgo -dest $(DIST)/binaries -tags 'netgo $(TAGS)' -ldflags '-linkmode external -extldflags "-static" $(LDFLAGS)' -targets 'windows/*' -out nxgit-$(VERSION) .
ifeq ($(CI),drone)
	mv /build/* $(DIST)/binaries
endif

.PHONY: release-linux
release-linux:
	@hash xgo > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u src.techknowlogick.com/xgo; \
	fi
	xgo -dest $(DIST)/binaries -tags 'netgo $(TAGS)' -ldflags '-linkmode external -extldflags "-static" $(LDFLAGS)' -targets 'linux/*' -out nxgit-$(VERSION) .
ifeq ($(CI),drone)
	mv /build/* $(DIST)/binaries
endif

.PHONY: release-darwin
release-darwin:
	@hash xgo > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u src.techknowlogick.com/xgo; \
	fi
	xgo -dest $(DIST)/binaries -tags 'netgo $(TAGS)' -ldflags '$(LDFLAGS)' -targets 'darwin/*' -out nxgit-$(VERSION) .
ifeq ($(CI),drone)
	mv /build/* $(DIST)/binaries
endif

.PHONY: release-copy
release-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

.PHONY: release-check
release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

.PHONY: release-compress
release-compress:
	@hash gxz > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/ulikunitz/xz/cmd/gxz; \
	fi
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),gxz -k -9 $(notdir $(file));)

.PHONY: javascripts
javascripts: public/js/index.js

.IGNORE: public/js/index.js
public/js/index.js: $(JAVASCRIPTS)
	cat $< >| $@

.PHONY: stylesheets-check
stylesheets-check: generate-stylesheets
	@diff=$$(git diff public/css/*); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make generate-stylesheets' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: generate-stylesheets
generate-stylesheets:
	@hash npx > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "Please install npm version 5.2+"; \
		exit 1; \
	fi;
	$(eval BROWSERS := "> 1%, last 2 firefox versions, last 2 safari versions, ie 11")
	npx lessc --clean-css public/less/index.less public/css/index.css
	$(foreach file, $(filter-out public/less/themes/_base.less, $(wildcard public/less/themes/*)),npx lessc --clean-css public/less/themes/$(notdir $(file)) > public/css/theme-$(notdir $(call strip-suffix,$(file))).css;)
	$(foreach file, $(wildcard public/css/*),npx postcss --use autoprefixer --autoprefixer.browsers $(BROWSERS) -o $(file) $(file);)

.PHONY: swagger-ui
swagger-ui:
	rm -Rf public/vendor/assets/swagger-ui
	git clone --depth=10 -b v3.13.4 --single-branch https://github.com/swagger-api/swagger-ui.git $(TMPDIR)/swagger-ui
	mv $(TMPDIR)/swagger-ui/dist public/vendor/assets/swagger-ui
	rm -Rf $(TMPDIR)/swagger-ui
	$(SED_INPLACE) "s;http://petstore.swagger.io/v2/swagger.json;../../../swagger.v1.json;g" public/vendor/assets/swagger-ui/index.html

.PHONY: update-translations
update-translations:
	mkdir -p ./translations
	cd ./translations && curl -L https://crowdin.com/download/project/nxgit.zip > nxgit.zip && unzip nxgit.zip
	rm ./translations/nxgit.zip
	$(SED_INPLACE) -e 's/="/=/g' -e 's/"$$//g' ./translations/*.ini
	$(SED_INPLACE) -e 's/\\"/"/g' ./translations/*.ini
	mv ./translations/*.ini ./options/locale/
	rmdir ./translations

.PHONY: generate-images
generate-images:
	mkdir -p $(TMPDIR)/images
	inkscape -f $(PWD)/assets/logo.svg -w 880 -h 880 -e $(PWD)/public/img/nxgit-lg.png
	inkscape -f $(PWD)/assets/logo.svg -w 512 -h 512 -e $(PWD)/public/img/nxgit-512.png
	inkscape -f $(PWD)/assets/logo.svg -w 192 -h 192 -e $(PWD)/public/img/nxgit-192.png
	inkscape -f $(PWD)/assets/logo.svg -w 120 -h 120 -jC -i layer1 -e $(TMPDIR)/images/sm-1.png
	inkscape -f $(PWD)/assets/logo.svg -w 120 -h 120 -jC -i layer2 -e $(TMPDIR)/images/sm-2.png
	composite -compose atop $(TMPDIR)/images/sm-2.png $(TMPDIR)/images/sm-1.png $(PWD)/public/img/nxgit-sm.png
	inkscape -f $(PWD)/assets/logo.svg -w 200 -h 200 -e $(PWD)/public/img/avatar_default.png
	inkscape -f $(PWD)/assets/logo.svg -w 180 -h 180 -e $(PWD)/public/img/favicon.png
	inkscape -f $(PWD)/assets/logo.svg -w 128 -h 128 -e $(TMPDIR)/images/128-raw.png
	inkscape -f $(PWD)/assets/logo.svg -w 64 -h 64 -e $(TMPDIR)/images/64-raw.png
	inkscape -f $(PWD)/assets/logo.svg -w 32 -h 32 -jC -i layer1 -e $(TMPDIR)/images/32-1.png
	inkscape -f $(PWD)/assets/logo.svg -w 32 -h 32 -jC -i layer2 -e $(TMPDIR)/images/32-2.png
	composite -compose atop $(TMPDIR)/images/32-2.png $(TMPDIR)/images/32-1.png $(TMPDIR)/images/32-raw.png
	inkscape -f $(PWD)/assets/logo.svg -w 16 -h 16 -jC -i layer1 -e $(TMPDIR)/images/16-raw.png
	zopflipng $(TMPDIR)/images/128-raw.png $(TMPDIR)/images/128.png
	zopflipng $(TMPDIR)/images/64-raw.png $(TMPDIR)/images/64.png
	zopflipng $(TMPDIR)/images/32-raw.png $(TMPDIR)/images/32.png
	zopflipng $(TMPDIR)/images/16-raw.png $(TMPDIR)/images/16.png
	rm -f $(TMPDIR)/images/*-*.png
	convert $(TMPDIR)/images/16.png $(TMPDIR)/images/32.png \
					$(TMPDIR)/images/64.png $(TMPDIR)/images/128.png \
					$(PWD)/public/img/favicon.ico
	rm -rf $(TMPDIR)/images
	
.PHONY: pr
pr:
	$(GO) run contrib/pr/checkout.go $(PR)
