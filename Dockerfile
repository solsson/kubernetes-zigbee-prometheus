FROM golang:1.13.8-alpine3.11@sha256:1ff752199f17b70e5f4dc2ad7f3e7843c456eb7e1407ed158ed8c237dbf1476a

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .

#RUN go test

RUN sed -i 's/zap.NewDevelopment()/zap.NewProduction()/' main.go

ARG arch=amd64
ARG goarm=6
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$arch GOARM=$goarm \
  go build -ldflags '-w -extldflags "-static"'

FROM gcr.io/distroless/base:nonroot@sha256:2b177fbc9a31b85254d264e1fc9a65accc6636d6f1033631b9b086ee589d1fe2

COPY --from=0 /src/kubernetes-zigbee-prometheus /usr/local/bin/kubernetes-zigbee-prometheus

ENTRYPOINT ["kubernetes-zigbee-prometheus"]
