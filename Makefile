dev_api: 
	go run cmd/api/main.go

build_docker_api:
	docker build -t 172.16.80.157/hacktiv8-final-project/h_service_user:v0.0.1 -f deploy/docker/api/Dockerfile . 