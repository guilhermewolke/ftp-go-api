FROM golang:latest as builder

WORKDIR /app

COPY . .

CMD GOOS=linux go -ldflags="-w -s" build -o server . 

ENTRYPOINT [ "./server" ] 


