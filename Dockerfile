FROM golang:1.21.6-alpine AS build
WORKDIR /build
COPY ./urlshortner .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o urlshortner .

FROM golang:1.21.6-alpine
WORKDIR /app
COPY --from=build /build/urlshortner /app/urlshortner
COPY ./configs /app/configs
EXPOSE 9080
ENTRYPOINT [ "./urlshortner", "--configPath=./configs","--configName=config.json"]
