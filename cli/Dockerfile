FROM proto:latest as builder

RUN go build -o counter-cli  ./cli/

FROM alpine:3.10

COPY --from=builder /go/src/github.com/calumball/ori/ .

ENTRYPOINT ["./counter-cli"]
