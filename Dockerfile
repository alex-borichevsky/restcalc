FROM golang:latest
LABEL maintainer="<https://github.com/borichevskiy/expr_rest-api>"
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main
COPY go.mod .
COPY go.sum .
COPY ./ ./
CMD ["/app/main"]
