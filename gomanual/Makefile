# Makefile

.EXPORT_ALL_VARIABLES:

build: 
	@echo "......................................................................................"
	CGO_ENABLED=0 go build --trimpath -ldflags="-s -w"

docker:
	@echo "--------------------------------------------------------------------------------------"
	docker build -f Dockerfile -t jeffotoni/gobootcampmanual .
	docker login
	docker push jeffotoni/gobootcampmanual:latest
