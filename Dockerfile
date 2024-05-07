FROM golang:1.22.3-alpine as builder

ARG TARGETPLATFORM
ARG REVISON

WORKDIR /workspace

# copy modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache modules
RUN go mod download

# copy source code
COPY cmd/ cmd/
COPY pkg/ pkg/

# build
RUN CGO_ENABLED=0 go build \
    -a -o bpdispatcher ./cmd/bpdispatcher

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

USER nobody

COPY --from=builder --chown=nobody:nobody /workspace/bpdispatcher .

ENTRYPOINT ["./bpdispatcher"]