BASE_PACKAGE_NAME=calcsvc
SERVER_DIR_NAME=calc
SERVER_BIN_NAME=calc
SERVER_BIN_PATH=cmd/$(SERVER_DIR_NAME)/$(SERVER_BIN_NAME)

DEV_DATASTORE_PORT=8081

GOTPL_BIN_PATH=$(GOPATH)/bin/gotpl
APP_YAML_TEMPLATE_PATH=app.yaml.tmpl
APP_YAML_PATH=app.yaml

.PHONY: dev
dev: $(SERVER_BIN_PATH) dev_datastore_check
	DATASTORE_EMULATOR_HOST=localhost:$(DEV_DATASTORE_PORT) \
	DATASTORE_PROJECT_ID=dummy-gcp-proejct \
		$(SERVER_BIN_PATH) --host development

${SERVER_BIN_PATH}:
	$(MAKE) build-server

.PHONY: build-server
build-server: cmd/$(SERVER_DIR_NAME)
	cd cmd/$(SERVER_DIR_NAME) && \
	go build

.PHONY: rebuild_server_dir
rebuild_server_dir: rm_server_dir_files cmd/$(SERVER_DIR_NAME)/main.go

cmd/$(SERVER_DIR_NAME)/main.go:
	$(MAKE) goa_example add_port_reading_from_env

.PHONY: goa_example
goa_example:
	goa example $(BASE_PACKAGE_NAME)/design

GOA_PORT_REPL=s/'u.Host += ":80"'/'port := os.Getenv("PORT"); if port == "" { port = "80" }; u.Host += (":" + port)'/g

.PHONY: add_port_reading_from_env
add_port_reading_from_env:
	sed -i '' ${GOA_PORT_REPL} cmd/$(SERVER_DIR_NAME)/main.go && \
	go fmt cmd/$(SERVER_DIR_NAME)/main.go

.PHONY: rm_server_dir_files
rm_server_dir_files:
	rm -f cmd/$(SERVER_DIR_NAME)/*.go


.PHONY: dev_datastore_start
dev_datastore_start:
	gcloud beta emulators datastore start --host-port localhost:$(DEV_DATASTORE_PORT)

.PHONY: dev_datastore_check
dev_datastore_check:
	(ps -ef | grep -v grep | grep cloud_datastore_emulator) || \
	(echo "No datastore_emulator found. Run \`make dev_datastore_start\` in another terminal" && exit 1)

.PHONY: deploy
deploy: require_GCP_PROJECT $(APP_YAML_PATH)
	gcloud --project=$(GCP_PROJECT) app deploy

$(GOTPL_BIN_PATH):
	go get github.com/tsg/gotpl

$(APP_YAML_PATH): $(GOTPL_BIN_PATH) require_GCP_PROJECT
	echo "GCP_PROJECT: $(GCP_PROJECT)" | $(GOTPL_BIN_PATH) $(APP_YAML_TEMPLATE_PATH) > $(APP_YAML_PATH)

.PHONY: require_GCP_PROJECT
require_GCP_PROJECT:
ifeq "$(GCP_PROJECT)" ""
	@echo "GCP_PROJECT must be given" && exit 1
endif
