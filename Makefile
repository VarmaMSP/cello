run-cello:
	go run ./cmd/cello/main.go

run-ui-server:
	cd ./ui-server && npm run dev

purge-data: 
	curl -X DELETE 'http://localhost:9200/_all'
	mysql -u root -pbirds phenopod < ./config/mysql/schema/v0.0.0.sql
	rabbitmqadmin purge queue name=create_thumbnail
	rabbitmqadmin purge queue name=import_podcast
	rabbitmqadmin purge queue name=refresh_podcast

start-services:
	nginx
	elasticsearch -d
	redis-server --daemonize yes
	rabbitmq-server -d

create-image-directories:
	mkdir /var/www && mkdir /var/www/img

nginx-reload:
	cp ./config/nginx/dev.conf /usr/local/etc/nginx/nginx.conf && nginx -s reload
