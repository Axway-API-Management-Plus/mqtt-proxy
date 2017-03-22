VERSION := 0.0.3
NAME := mqtt-proxy
DATE := $(shell date +'%Y-%M-%d_%H:%M:%S')
BUILD := $(shell git rev-parse HEAD | cut -c1-8)
LDFLAGS :=-ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -X=main.Date=$(DATE)"
IMAGE := jdavanne/$(NAME)
.PHONY: docker all

all: build

build:
	(cd src ; go build -o ../$(NAME) $(LDFLAGS))

docker-test:
	docker-compose -f docker-compose.test.yml down
	docker-compose -f docker-compose.test.yml build
	docker-compose -f docker-compose.test.yml run sut

docker-test-logs:
	docker-compose -f docker-compose.test.yml logs

clean:
	rm -f $(NAME) $(NAME).tar.gz

test:
	for dir in $$(find . -name "*_test.go" | grep -v ./vendor | xargs dirname | sort -u -r); do echo "$$dir..."; go test -v $$dir || exit 1 ; done | tee output.txt
	cat output.txt | egrep -- "--- FAIL:|--- SKIP:" || true

test-specific:
	go test -v $$(ls *.go | grep -v "_test.go") $(ARGS)

deps:
	go list -f '{{range .TestImports}}{{.}} {{end}} {{range .Imports}}{{.}} {{end}}' ./... | sed 's/ /\n/g' | grep -e "^[^/_\.][^/]*\.[^/]*/" |sort -u >.deps

deps-install:
	go get -v $$(cat .deps)
	#for dep in $$(cat .deps); do echo "installing '$$dep'... "; go get -v $$dep; done

deps-install-force: deps
	go get -u -v $$(cat .deps)
	#for dep in $$(cat .deps); do echo "installing '$$dep'... "; go get -u -v $$dep; done

docker-run:
	docker-compose up

docker:
	docker build -t $(IMAGE):build .
	docker run --rm $(IMAGE):build tar cz $(NAME) >$(NAME).tar.gz
	docker build -t $(IMAGE) -f Dockerfile.small .
	docker rmi $(IMAGE):build
