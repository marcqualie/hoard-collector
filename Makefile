default: docker-build

build:
	`<.env xargs -I % echo "export %"`; \
	go build -o hoard-collector *.go

run:
	`<.env xargs -I % echo "export %"`; \
	go run *.go

docker-build:
	docker build -t tutum.co/marcqualie/hoard-collector .

docker-run: docker-build
	docker run --env-file .env -e "PORT=80" -h "hoard-collector-development" -a STDOUT -a STDERR -p "80:80" -i --rm --name hoard-collector tutum.co/marcqualie/hoard-collector

docker-bash: docker-build
	docker run --env-file .env -e "PORT=80" -h "hoard-collector-development" -i -t --rm --name hoard-collector-bash tutum.co/marcqualie/hoard-collector bash

docker-kill:
	docker kill hoard-collector
