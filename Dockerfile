FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download || echo "No dependencies to download."

COPY . .

EXPOSE 80 

CMD ["go", "run", "main.go"]
