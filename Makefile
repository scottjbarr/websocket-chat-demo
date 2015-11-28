run: build
	heroku local

build:
	go install ./...

deps:
	GO15VENDOREXPERIMENT=0 godep save -r ./...
