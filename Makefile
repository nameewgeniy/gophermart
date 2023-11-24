
up:
	@docker-compose -f build/docker-compose.yaml up -d

build.app:
	@go build -o cmd/gophermart/gophermart cmd/gophermart/*.go

db.migration.make:
	@goose --dir internal/storage/pg/migrations create $(name) sql

test.get:
	@wget -O gophermarttest  https://github.com/Yandex-Practicum/go-autotests/releases/download/v0.10.2/gophermarttest
	@chmod +x gophermarttest

test.run: build.app
	@gophermarttest \
		-test.v -test.run=^TestGophermart$ \
		-gophermart-binary-path=cmd/gophermart/gophermart \
		-gophermart-host=localhost \
		-gophermart-port=8090 \
		-gophermart-database-uri="postgres://user:password@localhost:5452/db?sslmode=disable" \
		-accrual-binary-path=cmd/accrual/accrual_linux_amd64 \
		-accrual-host=localhost \
		-accrual-port=8081 \
		-accrual-database-uri="postgres://user:password@localhost:5452/db?sslmode=disable"