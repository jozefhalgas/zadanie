.PHONY: docker docker-run clean

backend: main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .

docker: Dockerfile backend
# TASK: Build docker image here

docker-run: docker
	docker run -p 9990:8080 -ti pixelfederation/sample-backend:0.1

clean:
	rm -rf backend
	docker rmi -f pixelfederation/sample-backend:0.1 | true
