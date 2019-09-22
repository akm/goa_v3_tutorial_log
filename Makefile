BASE_PACKAGE_NAME=calcsvc
SERVER_DIR_NAME=calc

.PHONY: build-server
build-server: cmd/$(SERVER_DIR_NAME)
	cd cmd/$(SERVER_DIR_NAME) && \
	go build

.PHONY: rebuild_server_dir
rebuild_server_dir: rm_server_dir_files cmd/$(SERVER_DIR_NAME)/main.go

cmd/$(SERVER_DIR_NAME)/main.go: goa_example add_port_reading_from_env

.PHONY: goa_example
goa_example:
	goa example $(BASE_PACKAGE_NAME)/design

GOA_PORT_REPL=s/'u.Host += ":80"'/'port := os.Getenv("PORT"); if port == "" { port = "80" }; u.Host += (":" + port)'/g

.PHONY: add_port_reading_from_env
add_port_reading_from_env	:
	sed -i '' ${GOA_PORT_REPL} cmd/$(SERVER_DIR_NAME)/main.go && \
	go fmt cmd/$(SERVER_DIR_NAME)/main.go

.PHONY: rm_server_dir_files
rm_server_dir_files:
	rm -f cmd/$(SERVER_DIR_NAME)/*.go
