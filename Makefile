#build:
#    statik -src = "./templates"
build:
	docker-compose build
test:
	docker-compose up -d
	sleep 5
	docker-compose exec dind sh ./init.sh
	docker-compose exec dind go test ./...
clean:
	docker-compose down -v
develop:
	docker-compose up -d
	docker-compose exec dind ash
