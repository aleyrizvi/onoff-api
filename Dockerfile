FROM golang:1.18 as base

RUN go install github.com/cosmtrek/air@latest
WORKDIR /api

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o api cmd/*.go

FROM gcr.io/distroless/base
COPY --from=base /api/app /

EXPOSE 9000

ENTRYPOINT ["./api"]