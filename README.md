# Go Assignment
Url : https://golang.yigitaltunay.com [Google Cloud Run]

## How To Run
```sh
cd github.com/yigitaltunay/go-assignment
go run .
```
## EndPoint List

```sh
GET https://golang.yigitaltunay.com/currency/EUR
```
```sh
GET https://golang.yigitaltunay.com/currency/USD
```
```sh
GET https://golang.yigitaltunay.com/currency/GBP
```

## Providers List

```sh
cd github.com/yigitaltunay/go-assignment/config
nano config.json
providers
```
## Test


```sh
cd github.com/yigitaltunay/go-assignment
go test .\test\ -v
```






## Docker

```sh
docker build -t ygtbackend .
docker run -p 8080:8080 -d  ygtbackend
```


## Tech

- [Gin] - Gin Web Framework
- [Viper] - Go configuration with fangs
- [Google-Cloud] - Cloud Deploy


## TODO

- Swagger