# Makefile

.EXPORT_ALL_VARIABLES:

build: 
	@echo "......................................................................................"
	CGO_ENABLED=0 go build --trimpath -ldflags="-s -w"

docker:
	@echo "--------------------------------------------------------------------------------------"
	docker build -f Dockerfile -t jeffotoni/zeroherostqueue:latest .
	docker run --rm --name zeroherostqueue -it -p 8080:8080 jeffotoni/zeroherostqueue:latest