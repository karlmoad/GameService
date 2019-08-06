.PHONY: all ref build service dist docker clean gcr
.DEFAULT_GOAL := build
TAG ?= latest
PROJECT_ID ?= unspecified

DIST_PATH := build

all: clean build

ref:
	go get -u

build: clean service.go
	go build -o $(DIST_PATH)/GameService ./service.go

dist: service.go
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o $(DIST_PATH)/GameService ./service.go

docker: clean dist
	docker build -t game_service:$(TAG) .

gcr: docker
	docker tag game_service:$(TAG) us.gcr.io/$(PROJECT_ID)/game_service:$(TAG)
	docker push us.gcr.io/$(PROJECT_ID)/game_service:$(TAG)

clean:
	rm -f $(DIST_PATH)/GameService