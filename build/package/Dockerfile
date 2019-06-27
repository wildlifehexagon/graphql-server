FROM golang:1.11.10-alpine3.9
LABEL maintainer="Siddhartha Basu <siddhartha-basu@northwestern.edu>"
LABEL maintainer="Eric Hartline <eric.hartline@northwestern.edu>"
RUN apk add --no-cache git build-base
RUN mkdir -p /graphql-server
WORKDIR /graphql-server
ADD cmd cmd
ADD internal internal
COPY go.mod go.sum ./
RUN go get ./...
RUN go build -o app cmd/graphql-server/main.go

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
COPY --from=0 /graphql-server/app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]