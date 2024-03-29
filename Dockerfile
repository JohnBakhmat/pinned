FROM golang:1.18-alpine
WORKDIR /usr/src/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./bin/pinned .
EXPOSE 8080
CMD [ "./bin/pinned" ]
