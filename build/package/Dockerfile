FROM golang:1.16.0-alpine3.13
LABEL maintainer="Eric Hartline <eric.hartline@northwestern.edu>"
LABEL maintainer="Siddhartha Basu <siddhartha-basu@northwestern.edu>"
ENV GOPROXY https://proxy.golang.org
RUN mkdir -p /graphql-server
WORKDIR /graphql-server
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD cmd cmd
ADD internal internal
RUN CGO_ENABLED=0 go build -o app cmd/graphql-server/main.go

FROM gcr.io/distroless/static
COPY --from=0 /graphql-server/app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]