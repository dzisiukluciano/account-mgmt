FROM golang:1.13 as builder

WORKDIR /account-mgmt-api
COPY . .

RUN go get -d -v ./...
RUN go build main.go

FROM golang:1.13

WORKDIR /code
COPY --from=builder /account-mgmt-api .
ENTRYPOINT ./main
EXPOSE 3000