.PHONY: tests

build:
	go build -o app

docker_build:
	docker build -t carlo/status-checker .

run:
	docker run -d --name webchecker --net=host carlo/status-checker

stop:
	docker stop webchecker

remove:
	docker rm webchecker

tests:
	go test -v

clean:
	rm *.log app
