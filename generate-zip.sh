GOOS=linux CGO_ENABLED=0 go build cmd/lambda/main.go
zip function.zip main