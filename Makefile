# Makefile for the project

composeup:
	docker compose --env-file=.env --verbose up -d --build

composedown:
	docker compose --env-file=.env down

applogs:
	docker compose --env-file=.env logs -f electromart_app

dblogs:
	docker compose --env-file=.env logs -f db

deletedbdata:
	rm -r dbdata

swag:
	swag init -g internal/http/router/route.go && swag fmt

run:
	swag init -g internal/http/router/route.go && swag fmt && docker compose --env-file=.env --verbose up -d --build

stop:
	docker compose --env-file=.env down && rm -r dbdata

.PHONY: composeup composedown applogs dblogs deletedbdata swag run stop
