FROM golang:1.22.3-bookworm

WORKDIR /services/application

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/zzwx/fresh@latest

COPY . .

RUN chmod +x ./start.sh

EXPOSE 8080

CMD ["./start.sh"]