export GO111MODULE=on

.PHONY: all generate serve deploy

all:

generate:
	go generate

serve:
	go run --tags dev .

deploy:
	gcloud builds submit --tag gcr.io/insta-tmthrgd-dev/server
	gcloud beta run deploy --image gcr.io/insta-tmthrgd-dev/server