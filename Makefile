run-hal:
	go run ./cmd/hal/main.go

run-anton:
	sudo go run ./cmd/anton/main.go

run-ui-server:
	cd ./ui-server && npm run dev

create-image-directories:
	mkdir /var/www && mkdir /var/www/img

nginx-reload:
	cp ./configs/nginx/dev.conf /usr/local/etc/nginx/nginx.conf && nginx -s reload
