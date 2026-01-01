build_ui:
	cd ui && yarn build
build: build_ui
	go build -o bin/app main.go
start: build build_ui
	rm -rf bin/public
	mv ui/dist bin/public
	cp .env bin
	cd bin && ./app

start_vercel: build_ui
	rm -rf public
	cp -r ui/dist public
	vercel dev -d