export GO111MODULE=on

.PHONY: all generate serve deploy docker-serve docker-stop

all:

generate:
	go generate

serve:
	go run --tags dev .

docker-serve:
	docker build -t insta-tmthrgd-dev-server -f Dockerfile .
	docker run -p 8090:8090 -e PORT=8090 -d insta-tmthrgd-dev-server

docker-stop:
	docker stop $(shell docker ps -f ancestor=insta-tmthrgd-dev-server --format "{{.ID}}")

deploy:
	gcloud builds submit --tag gcr.io/insta-tmthrgd-dev/server
	gcloud beta run deploy server --image gcr.io/insta-tmthrgd-dev/server