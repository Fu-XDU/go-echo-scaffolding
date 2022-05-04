VERSION ?= 0.0.1
IMG_NAME ?= echo/golang-server:${VERSION}
DB_NAME ?= echo_server

docker:
	docker build -t ${IMG_NAME} -f Dockerfile .

start:
	IMG_NAME=${IMG_NAME} DB_NAME=${DB_NAME} docker-compose -f docker-compose.yaml up -d

stop:
	docker-compose -f docker-compose.yaml down

restart:
	docker-compose -f docker-compose.yaml restart