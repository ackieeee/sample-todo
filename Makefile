DOCKER_EXEC=docker-compose exec

run-appserver:
	${DOCKER_EXEC} appserver go run ./main.go
rebuild:
	docker-compose down && docker-compose up -d --build
test:
	go test -v `go list ./... | grep -v models`
