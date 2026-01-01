build_ui:
	cd ui && yarn build

build: build_ui
	go build -o bin/app main.go
build_vercel: build_ui
	go build -o bin/app api/index.go
start: build
	bin/app