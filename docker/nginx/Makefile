.PHONY: docker clean

docker: Dockerfile default.conf index.html start.sh
# TASK: Build NGINX Docker image
	docker build -t pxlfrontend -f Dockerfile .

docker-run: docker
	docker run -p 80:80 -ti pxlfrontend

clean:
	docker rmi -f pxlfrontend | true
