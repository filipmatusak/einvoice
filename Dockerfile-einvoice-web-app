FROM node:12-alpine AS ui_builder

COPY einvoice-web-app/client /client

WORKDIR /client
RUN npm install
RUN npm run build

FROM golang:1.13-alpine AS server_builder

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /server

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v ./einvoice-web-app/server

RUN go build -o /out/server ./einvoice-web-app/server

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=server_builder /out/server /server/server
COPY --from=ui_builder /client/build /client/build


# This container exposes port 8081 to the outside world
EXPOSE 8081

CMD ["/server/server"]