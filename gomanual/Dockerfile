FROM golang:1.17 as builder
WORKDIR /go/src/gobootcampmanual
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build --trimpath -ldflags="-s -w -X main.version=docker" -o gobootcampmanual main.go
RUN cp gobootcampmanual /go/bin/gobootcampmanual

FROM alpine:latest as builder2
RUN apk add --no-cache upx
COPY --from=builder /go/bin/gobootcampmanual /go/bin/gobootcampmanual
WORKDIR /go/bin
RUN upx gobootcampmanual
RUN apk del --no-cache upx
RUN apk del --no-cache tzdata

FROM scratch
COPY --from=builder2 /go/bin/gobootcampmanual /
EXPOSE 8080
ENTRYPOINT ["/gobootcampmanual"]
