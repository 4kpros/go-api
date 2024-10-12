# ------------------ Golang commands ------------------
.PHONY: clean install update test build run
clean:
	@go clean -cache
	@go clean -testcache
	@go clean -modcache
install:
	@go mod download
update:
	@go get -u all
test:
	@go test -v ./tests/...
build:
	@cd cmd/ ;\
	go build -o ../.build/main ;\
	cd ../
run:
	@./.build/main

# Third party libraries commands
.PHONY: scan
scan:
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@govulncheck ./...


# ------------------ Docker commands ------------------
.PHONY: docker-redis docker-postgres docker-api
docker-redis:
	@docker-compose --env-file app.env up --build --no-deps -d redis
docker-postgres:
	@docker-compose --env-file app.env up --build --no-deps -d postgres
docker-api:
	@docker-compose --env-file app.env up --build --no-deps -d api

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
	gRepo="go-api" ;\
	read -p "Enter your package name(redis, postgres, api): " gPackage; gTag=$${gPackage:-"api"} ;\
	read -p "Enter your tag(default is 0.0.1): " gTag; gTag=$${gTag:-"0.0.1"} ;\
	docker tag go-api-$$gPackage ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag ;\
	echo "" ;\
	echo "Pushing $$gPackage - GitHub Docker Registry" ;\
	docker push ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag;
docker-ghcr-push-all:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="4kpros" ;\
	gRepo="go-api" ;\
	read -p "Enter your tag(default is 0.0.1): " gTag; gTag=$${gTag:-"0.0.1"} ;\
	docker tag go-api-redis ghcr.io/$$gCorp/$$gRepo/redis:$$gTag ;\
	docker tag go-api-postgres ghcr.io/$$gCorp/$$gRepo/postgres:$$gTag ;\
	docker tag go-api-api ghcr.io/$$gCorp/$$gRepo/api:$$gTag ;\
	echo "" ;\
	echo "Pushing all - GitHub Docker Registry" ;\
	docker push ghcr.io/$$gCorp/$$gRepo/redis:$$gTag ;\
	docker push ghcr.io/$$gCorp/$$gRepo/postgres:$$gTag ;\
	docker push ghcr.io/$$gCorp/$$gRepo/api:$$gTag ;\

.PHONY: docker-ghcr-pull-all docker-ghcr-pull-specific
docker-ghcr-pull-specific:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="4kpros" ;\
	gRepo="go-api" ;\
	read -p "Enter your package name(redis, postgres, api): " gPackage; gTag=$${gPackage:-"api"} ;\
	read -p "Enter your tag(default is 0.0.1): " gTag; gTag=$${gTag:-"0.0.1"} ;\
	echo "" ;\
	echo "Puslling $$gPackage - GitHub Docker Registry" ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag
docker-ghcr-pull-all:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="4kpros" ;\
	gRepo="go-api" ;\
	read -p "Enter your tag(default is 0.0.1): " gTag; gTag=$${gTag:-"0.0.1"} ;\
	echo "" ;\
	echo "Puslling all - GitHub Docker Registry" ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/redis:$$gTag ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/postgres:$$gTag ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/api:$$gTag

