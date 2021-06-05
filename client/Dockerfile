FROM golang:alpine AS development

ENV CGO_ENABLED=0
COPY . /go/src/app/
WORKDIR /go/src/app/
ENV GOPATH=/go/

RUN apk add --no-cache \
    git \
    curl \
    gcc \
    zip \
    && mkdir /build/ \
    # && cp -r docs /build \
    && go get github.com/go-delve/delve/cmd/dlv


# Let's keep it in a separate layer
RUN go build -o /build/app .
ENTRYPOINT [ "dlv", "debug", "--headless", "--log", "--listen=:2345", "--api-version=2"]

#----------------------------#

FROM development AS test

WORKDIR /go/src/app/

ENV EXEC_PATH=/go/src/app/

ENTRYPOINT ["go", "test", "-v", "./..."]

#----------------------------#

FROM alpine:latest AS production

WORKDIR /app/
COPY --from=development /build .
RUN apk --no-cache add \
    curl 

ENTRYPOINT ["./app"]
