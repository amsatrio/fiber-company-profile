build_ui:
	cd ui && yarn build

build: build_ui
	go build -o bin/app main.go

start: build
	bin/app