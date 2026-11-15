FROM golang:1.27-alpine

WORKDIR /app

# for hot reloading - compile on code change
RUN go install github.com/air-verse/air@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

CMD ["air"]