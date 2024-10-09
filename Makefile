# ------------------ Golang commands ------------------
.PHONY: install swagger test build run
install:
	@go mod download
	@go install github.com/swaggo/swag/cmd/swag@latest
swagger:
	@swag init --generalInfo ./cmd/main.go --output ./docs --parseDependency ./docs
test:
	@go test -v ./tests/...
build:
	@cd cmd/ ;\
	go build -o ../.build/main ;\
	cd ../
run:
	@./.build/main


# ------------------ Docker commands ------------------
.PHONY: docker-vault docker-redis docker-postgres docker-api docker-nginx
docker-vault:
	@docker-compose --env-file app.env up --build --no-deps -d vault
docker-redis:
	@docker-compose --env-file app.env up --build --no-deps -d redis
docker-postgres:
	@docker-compose --env-file app.env up --build --no-deps -d postgres
docker-api:
	@docker-compose --env-file app.env up --build --no-deps -d api
docker-nginx:
	@docker-compose --env-file app.env up --build --no-deps -d nginx

.PHONY: docker-ghcr-login
docker-ghcr-login:
	@echo "" ;\
	echo "Login - GitHub Docker Registry" ;\
	read -p "Enter your Github username: " gUsername ;\
	read -p "Enter your Github personal access token: " gPass ;\
	echo "" ;\
	echo $$gPass | docker login ghcr.io -u $$gUsername --password-stdin ;\

.PHONY: docker-ghcr-push-all docker-ghcr-push-specific
docker-ghcr-push-specific:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="4kpros" ;\
	gRepo="api" ;\
	read -p "Enter your package name(vault, redis, memcached, postgres, api): " gPackage; gTag=$${gPackage:-"api"} ;\
	read -p "Enter your tag(default is 0.0.1): " gTag; gTag=$${gTag:-"0.0.1"} ;\
	docker tag snip-backend-$$gPackage ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag ;\
	echo "" ;\
	echo "Pushing $$gPackage - GitHub Docker Registry" ;\
	docker push ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag;
docker-ghcr-push-all:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="emenec-finance" ;\
	gRepo="snip-backend" ;\
	read -p "Enter your tag(default is 0.0.1): " gTag; gTag=$${gTag:-"0.0.1"} ;\
	docker tag snip-backend-vault ghcr.io/$$gCorp/$$gRepo/vault:$$gTag ;\
	docker tag snip-backend-postgres ghcr.io/$$gCorp/$$gRepo/postgres:$$gTag ;\
	docker tag snip-backend-api ghcr.io/$$gCorp/$$gRepo/api:$$gTag ;\
	echo "" ;\
	echo "Pushing all - GitHub Docker Registry" ;\
	docker push ghcr.io/$$gCorp/$$gRepo/vault:$$gTag ;\
	docker push ghcr.io/$$gCorp/$$gRepo/redis:$$gTag ;\
	docker push ghcr.io/$$gCorp/$$gRepo/memcached:$$gTag ;\
	docker push ghcr.io/$$gCorp/$$gRepo/postgres:$$gTag ;\
	docker push ghcr.io/$$gCorp/$$gRepo/api:$$gTag ;\

.PHONY: docker-ghcr-pull-all docker-ghcr-pull-specific
docker-ghcr-pull-specific:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="emenec-finance" ;\
	gRepo="snip-backend" ;\
	read -p "Enter your package name(vault, redis, memcached, postgres, api): " gPackage; gTag=$${gPackage:-"api"} ;\
	read -p "Enter your tag(default is 0.0.1): " gTag; gTag=$${gTag:-"0.0.1"} ;\
	echo "" ;\
	echo "Puslling $$gPackage - GitHub Docker Registry" ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag
docker-ghcr-pull-all:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="emenec-finance" ;\
	gRepo="snip-backend" ;\
	read -p "Enter your tag(default is 0.0.1): " gTag; gTag=$${gTag:-"0.0.1"} ;\
	echo "" ;\
	echo "Puslling all - GitHub Docker Registry" ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/vault:$$gTag ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/redis:$$gTag ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/memcached:$$gTag ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/postgres:$$gTag ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/api:$$gTag

