# Makefile

.EXPORT_ALL_VARIABLES:

build: 
	@echo "......................................................................................"
	CGO_ENABLED=0 go build --trimpath -ldflags="-s -w"

docker:
	@echo "--------------------------------------------------------------------------------------"
	docker build -f Dockerfile -t jeffotoni/zeroherofiber:latest .
	docker run --rm --name zeroherofiber -it -p 8080:8080 jeffotoni/zeroherofiber:latest