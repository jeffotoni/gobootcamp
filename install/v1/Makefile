# Makefile
.EXPORT_ALL_VARIABLES:	

GOPRIVATE=github.com/jeffotoni/gobootcamp

install:
	@echo "########## Compilando nosso gobootcamp ... "
	@cp -r ../../gomanual/manual .
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gobootcamp main.go
	upx gobootcamp
	@echo "buid completo..."
	@rm -rf manual
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"
