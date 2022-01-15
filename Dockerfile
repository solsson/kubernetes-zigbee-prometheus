FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.17.6-bullseye

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .

#RUN go test

RUN sed -i 's/zap.NewDevelopment()/zap.NewProduction()/' main.go

ARG TARGETOS TARGETARCH
# TODO what about GOARM?
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH \
  CGO_ENABLED=0 \
  go build -ldflags '-w -extldflags "-static"'

FROM --platform=${TARGETPLATFORM:-linux/amd64} gcr.io/distroless/static-debian11:nonroot

COPY --from=0 /src/kubernetes-zigbee-prometheus /usr/local/bin/kubernetes-zigbee-prometheus

ENTRYPOINT ["kubernetes-zigbee-prometheus"]
