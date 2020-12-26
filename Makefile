export VERSION 	?= $(shell git show -q --format=%h)
export IMAGE 		?= ghcr.io/coronatorid/core-onator

test:
	go test -race -cover -coverprofile=cover.out $$(go list ./... | grep -Ev "coronator$$|testutil|mocks|testhelper")

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./coronator

build_docker:
	sudo docker build -t $(IMAGE):$(VERSION) -f ./Dockerfile .
