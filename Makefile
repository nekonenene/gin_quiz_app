.PHONY: init
init:
	dep ensure
	cp default.env .env
