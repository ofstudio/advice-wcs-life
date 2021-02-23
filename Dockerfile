FROM golang:1.16-alpine AS build
WORKDIR /src
COPY [".", "."]
RUN go get -d -v
RUN go build -ldflags "-s -w" -o /build/server .

FROM alpine
USER nobody
COPY --from=build ["/build/server", "/"]
EXPOSE 8080
CMD ["/server"]
