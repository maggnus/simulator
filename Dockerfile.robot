FROM golang:1.24.5-alpine AS build

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY api ./api
COPY cmd/simulated_robot ./cmd/simulated_robot
COPY internal/robot ./internal/robot
COPY configs ./configs
COPY ent ./ent

RUN apk add --no-cache build-base sqlite-dev

RUN CGO_ENABLED=1 GOOS=linux go build -o app ./cmd/simulated_robot

ARG TAG=latest
FROM golang:1.24.5-alpine

RUN apk add --no-cache sqlite-libs

WORKDIR /root/

COPY --from=build /go/src/app/app .
COPY --from=build /go/src/app/configs ./configs

CMD ["./app"]
