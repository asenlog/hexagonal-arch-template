FROM golang:1.20 as builder

RUN cd ..
WORKDIR build
COPY . ./

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download -x

RUN chmod +x ./scripts/*.sh

RUN CGO_ENABLED=0 go build -installsuffix cgo  -ldflags '-extldflags "-static"' -o api ./cmd/service/main.go

FROM alpine:3.15.1

# Update the index of available packages
RUN apk update

# Upgrade the currently installed packages
RUN apk upgrade

# Install psql client for the wait_and_run.sh script
RUN apk add postgresql-client

COPY --from=builder /go/build/api .
COPY --from=builder /go/build/scripts .

CMD ["./wait_and_run.sh"]