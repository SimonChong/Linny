

rm insights.db

go generate -x ./...

go run main.go -serve
