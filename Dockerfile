# Build
FROM golang:1.22.0-alpine as build-env
ENV APP_NAME skopsgo
ENV CMD_PATH cmd/skopsgo/skopsgo.go
COPY . $GOPATH/src/$APP_NAME
COPY .env $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# Run
FROM alpine:3.14
ENV APP_NAME skopsgo
COPY --from=build-env /$APP_NAME .
COPY .env .
EXPOSE 42069
CMD ./$APP_NAME
