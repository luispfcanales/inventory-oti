run: build
	@./bin/main
build:
	@go build -o ./bin/main .
start: build
	@./service-daemon.sh
	@echo "[ Execute next commands to load service ]"
	@echo "systemctl start goweb.service"
	@echo "systemctl stop goweb.service"
	@echo "systemctl enable goweb.service"
