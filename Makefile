run: build
	@./main
build:
	@go clean
	@go build -o main .
bucket:
	s3fs oti-bucket /mnt/s3 -o passwd_file=~/.passwd-s3fs
start: build
	@./service-daemon.sh
	@echo "[ Execute next commands to load service ]"
	@echo "systemctl start goweb.service"
	@echo "systemctl stop goweb.service"
	@echo "systemctl enable goweb.service"
