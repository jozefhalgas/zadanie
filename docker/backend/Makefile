.PHONY: docker docker-run clean

#backend: main.go
#	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .

docker: Dockerfile main.go #backend -> zmenene na main.go, build robim v containeri
# TASK: Build docker image here
	docker build -t pxlbackend -f Dockerfile .

docker-run: docker
	docker run -p 9990:8080 -ti pxlbackend

clean:
	#rm -rf backend -> nepotrebujem to mazat lebo je v  containeri
	docker rmi -f pxlbackend | true
