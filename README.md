## Rest-api for expression calculator

## Installation:
```go get github.com/borichevskiy/restcalc```
## Usage:
1)run throw docker-compose with cli, that examine rest server with some goroutines(worker-pool)
```go
docker-compose up
```
2)run throw docker
```go
docker build -t image_name .
docker run -p 8080:9000 image_name
```
3)simple start server
```go
make run
```
Then go to http://localhost:9000/evaluate/?expr=" " and enter your expression

Swagger supported. To see : 

go to http://localhost:9000/swagger/index.html

## Contributing:

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
