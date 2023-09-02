FROM golang:1.20.2-bullseye AS builder

RUN apt-get update \
    && apt-get install -y  --no-install-recommends \
        git \
        make
    
ENV APP_DIR $GOPATH/src/github.com/projectulterior/2cents-backend
WORKDIR ${APP_DIR}

COPY . ${APP_DIR}

RUN go install github.com/99designs/gqlgen@latest
RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN make build

FROM alpine:latest

ENV PORT=8080
EXPOSE $PORT

COPY --from=builder /go/src/github.com/projectulterior/2cents-backend/bin/daemon /
ENTRYPOINT [ "/daemon" ]
