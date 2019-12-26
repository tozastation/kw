#statik:
#    statik -src=./templates
test:
	docker build -t kw -f Dockerfile.test .
	docker run -v /var/run/docker.sock:/var/run/docker.sock -it kw bash