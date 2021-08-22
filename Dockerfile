FROM golang:1.16 as builder

RUN apk update \
  && apk add --no-cache git curl make gcc g++ \
  && go get github.com/oxequa/realize

WORKDIR /go/src

COPY ./go.mod ./go.sum ./

ENV GO111MODULE=on

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./app ./main.go

FROM alpine

RUN apk add tzdata
ENV TZ=Asia/Tokyo

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app /go/src/app

COPY .env /go/src

ENV DOTENV_PATH=/go/src/.env
ENV GIN_MODE=release

CMD ["/go/src/app"]
