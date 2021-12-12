module zerohero/fiber

go 1.17

require github.com/gofiber/fiber/v2 v2.23.0

require (
	interno/handler v0.0.0
	interno/middleware v0.0.0
)

require (
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gofiber/fiber v1.14.6 // indirect
	github.com/gofiber/utils v0.0.10 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.31.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	golang.org/x/text v0.3.6 // indirect
	interno/mongo v0.0.0 // indirect
)

replace (
	interno/handler => ./interno/handler
	interno/middleware => ./interno/middleware
	interno/mongo => ./interno/mongo
)
