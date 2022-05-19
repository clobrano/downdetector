.PHONY: tests

build:
	go build -o app

docker_build:
	docker build \
		--build-arg TOKEN="${TELEGRAM_TOKEN}" \
		--build-arg CLIENT_ID="${WEB_CHECKER_TELEGRAM_CLIENT_ID}" \
		-t carlo/status-checker .

run:
	docker run --rm -d --name webchecker --net=host carlo/status-checker

stop:
	docker stop webchecker

tests:
	go test -v

clean:
	rm *.log app
