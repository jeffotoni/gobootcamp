#################################################
# Go + Strach + Muiltistage
#################################################
FROM golang:1.17.0 AS builder
WORKDIR /go/src/main
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X main.ambiente=docker" -o zerohero

FROM alpine:latest as builder2
RUN apk add --no-cache upx
RUN apk add --no-cache tzdata
COPY --from=builder /go/src/main /go/src/main
WORKDIR /go/src/main
RUN upx zerohero
ENV TZ="America/Sao_Paulo"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

FROM scratch
# Copy our static executable.
COPY --from=builder2 /go/src/main /
# Run the hello binary.
ENTRYPOINT ["/zerohero"]
