.PHONY: docker clean

docker: Dockerfile default.conf
# TASK: Build NGINX Docker image

clean:
	docker rmi -f pixelfederation/sample-nginx:0.1 | true
