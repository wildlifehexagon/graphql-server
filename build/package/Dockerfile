FROM golang:1.11.10-alpine3.10
LABEL maintainer="Eric Hartline <eric.hartline@northwestern.edu>"
LABEL maintainer="Siddhartha Basu <siddhartha-basu@northwestern.edu>"
ENV GOPROXY https://proxy.golang.org
RUN apk add --no-cache git build-base
RUN mkdir -p /graphql-server
WORKDIR /graphql-server
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD cmd cmd
ADD internal internal
RUN go build -o app cmd/graphql-server/main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=0 /graphql-server/app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]