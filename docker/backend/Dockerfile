# TASK:
# Treba implementovat vraci zadania
#
# Aplikacia je by default dostupna na porte 8080
FROM golang
WORKDIR /go/src/pxl_backend
COPY main.go .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .
FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/src/pxl_backend/backend .
CMD [ "./backend" ]
