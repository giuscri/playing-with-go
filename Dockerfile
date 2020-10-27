FROM golang as build

WORKDIR /go/src/app
COPY . .

RUN go test && \
    go build

FROM scratch
COPY --from=build /go/src/app/playing /usr/local/bin/playing

CMD ["/usr/local/bin/playing"]
