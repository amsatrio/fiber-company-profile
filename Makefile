build_ui:
	cd ui && yarn build
build: build_ui
	go build -o bin/app main.go
start: build build_ui
	mv ui/dist bin/public
	cp .env bin
	cd bin && ./app