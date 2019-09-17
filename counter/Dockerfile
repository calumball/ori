FROM proto:latest as builder 
        
WORKDIR /go/src/github.com/calumball/ori/counter
COPY . .

RUN go build -o counter .

FROM alpine:3.10

COPY --from=builder /go/src/github.com/calumball/ori/counter/ .

ENV COUNTER_PORT=8888
EXPOSE 8888

CMD ["./counter"]
