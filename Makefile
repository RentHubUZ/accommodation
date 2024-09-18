CURRENT_DIR=$(shell pwd)
DB_URL=postgres://postgres:pass@localhost:5432/rent_hub_accommodation?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

run :
	go run cmd/main.go
  
mig-up:
	migrate -path migrations -database ${DB_URL}  -verbose up

mig-down:
	migrate -path migrations -database ${DB_URL}  -verbose down

mig-force:
	migrate -path migrations -database ${DB_URL}  -verbose force 1

mig-file:
	migrate create -ext sql -dir migrations -seq create_tables

test:
	go test -v -cover ./...
tidy:
	go mod tidy