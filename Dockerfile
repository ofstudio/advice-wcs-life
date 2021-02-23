FROM golang:1.16-alpine AS build
WORKDIR /src
COPY [".", "."]
RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /build/server .

FROM scratch
COPY --from=build ["/build/server", "/"]
EXPOSE 8080
CMD ["/server"]
