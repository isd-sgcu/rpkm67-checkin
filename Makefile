docker:
	docker-compose up

server:
	go run cmd/main.go

watch: 
	air

mock-gen:
	mockgen -source ./internal/checkin/checkin.service.go -destination ./mocks/checkin/checkin.service.go
	mockgen -source ./internal/checkin/checkin.repository.go -destination ./mocks/checkin/checkin.repository.go

test:
	go vet ./...
	go test  -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

proto:
	go get github.com/isd-sgcu/rpkm67-go-proto@latest

model:
	go get github.com/isd-sgcu/rpkm67-model@latest
