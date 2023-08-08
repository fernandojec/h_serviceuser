FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o app cmd/api/main.go

FROM progrium/busybox
RUN  opkg-install ca-certificates

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 443

CMD [ "./app" ]