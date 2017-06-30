VERSION := 0.0.4
NAME := mqtt-proxy
DATE := $(shell date +'%Y-%M-%d_%H:%M:%S')
BUILD := $(shell git rev-parse HEAD | cut -c1-8)
LDFLAGS :=-ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -X=main.Date=$(DATE)"
IMAGE := $(NAME)
REGISTRY := registry.dctest.docker-cluster.axwaytest.net/internal
PUBLISH := $(REGISTRY)/$(IMAGE)

.PHONY: docker all certs deps

all: build

build:
	(cd src ; go build -o ../$(NAME) $(LDFLAGS))

dev:
	ls -d src/* | entr -r sh -c "make && ./mqtt-proxy --mqtt-broker-host localhost --mqtt-broker-port 9883"

docker-test:
	docker-compose -f docker-compose.test.yml down
	docker-compose -f docker-compose.test.yml build
	docker-compose -f docker-compose.test.yml run sut  || (docker-compose -f docker-compose.test.yml logs -t | sort -k 3 ; docker-compose -f docker-compose.test.yml down ; exit 1)
	docker-compose -f docker-compose.test.yml down

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
	docker build -t $(IMAGE) .

docker-publish-all: docker-publish docker-publish-version

docker-publish-version:
	docker tag $(IMAGE) $(PUBLISH):$(VERSION)
	docker push $(PUBLISH):$(VERSION)

docker-publish: docker
	docker tag $(IMAGE) $(PUBLISH):latest
	docker push $(PUBLISH):latest

certs: certs-proxy certs-policy

certs-proxy:
	openssl genrsa -out certs/server.key 2048
	openssl req -new -x509 -sha256 -key certs/server.key -out certs/server.pem -days 3650 -subj "/C=FR/ST=Paris/L=La Defense/O=Axway/CN=mqtt-proxy"
	openssl x509 -text -noout -in certs/server.pem
	cp certs/server.pem tests/test/certs/mqtt-proxy.pem

certs-policy:
	openssl genrsa -out tests/policy/certs/server.key 2048
	openssl req -new -x509 -sha256 -key tests/policy/certs/server.key -out tests/policy/certs/server.pem -days 3650 -subj "/C=FR/ST=Paris/L=La Defense/O=Axway/CN=policy"
	openssl x509 -text -noout -in tests/policy/certs/server.pem
	cp tests/policy/certs/server.pem certs/policy.pem
