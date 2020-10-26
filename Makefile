NETWORK_NAME := phenopod-network

docker-compose-up:
	@$(MAKE) verify-network
	DOCKER_BUILDKIT=1 && docker-compose up --remove-orphans --build

open-redis-cli:
	@$(MAKE) verify-network
	docker run -it --rm --network $(NETWORK_NAME) redis:6.0 redis-cli -h redis

open-mysql-cli:
	@$(MAKE) verify-network
	docker run -it --rm --network $(NETWORK_NAME) mysql:8.0 mysql -h mysql -u root -pbirds

verify-network:
	@docker network create -d bridge $(NETWORK_NAME) || true

#
# DEPRECATED: Not being used anymore with the latest docker setup
#

# purge-data: 
# 	curl -X DELETE 'http://localhost:9200/_all'
# 	rabbitmqadmin purge queue name=create_thumbnail
# 	rabbitmqadmin purge queue name=create_thumbnail_dead_letter
# 	rabbitmqadmin purge queue name=import_podcast
# 	rabbitmqadmin purge queue name=refresh_podcast
# 	rabbitmqadmin purge queue name=sync_playback
# 	mc rm -r --force --dangerous minio/chartable-charts
# 	mc rm -r --force --dangerous minio/phenopod-charts
# 	mc rm -r --force --dangerous minio/thumbnails
# 	mysql -u root -pbirds phenopod < /Users/varmamsp/Documents/code/cello/config/db_schema/01\ -\ feed.sql
# 	mysql -u root -pbirds phenopod < /Users/varmamsp/Documents/code/cello/config/db_schema/02\ -\ user.sql
# 	mysql -u root -pbirds phenopod < /Users/varmamsp/Documents/code/cello/config/db_schema/03\ -\ playlist.sql

# start-services:
# 	nginx
# 	elasticsearch -d
# 	redis-server --daemonize yes
# 	rabbitmq-server -d

# start-minio:
# 	minio server ~/minio-data

# create-image-directories:
# 	mkdir /var/www && mkdir /var/www/img

# nginx-reload:
# 	cp ./config/proxy.conf /usr/local/etc/nginx/nginx.conf && nginx -s reload
