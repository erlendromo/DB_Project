# Makefile for the project

## Deployment commands

composeup:
	docker compose --env-file=.env --verbose up -d --build

composedown:
	docker compose --env-file=.env down

applogs:
	docker compose --env-file=.env logs -f electromart_app

dblogs:
	docker compose --env-file=.env logs -f db



## Development commands

deletedbdata:
	rm -r dbdata

swag:
	swag init -g internal/http/router/route.go && swag fmt

run:
	swag init -g internal/http/router/route.go && swag fmt && docker compose --env-file=.env --verbose up -d --build

stop:
	docker compose --env-file=.env down && rm -r dbdata

# Replace <absolute-path-to-swag-binary-directory> with the absolute path to the directory where the swag binary is located
# Example -> <absolute-path-to-swag-binary-directory> === /Users/username/go/bin/
# Run this if the swag binary is not in the PATH
swagabsolute:
	<absolute-path-to-swag-binary-directory>swag init -g internal/http/router/route.go && <absolute-path-to-swag-binary-directory>swag fmt

# Replace <absolute-path-to-swag-binary-directory> with the absolute path to the directory where the swag binary is located
# Example -> <absolute-path-to-swag-binary-directory> === /Users/username/go/bin/
# Run this if the swag binary is not in the PATH
runabsolute:
	<absolute-path-to-swag-binary-directory>swag init -g internal/http/router/route.go && <absolute-path-to-swag-binary-directory>swag fmt && docker compose --env-file=.env --verbose up -d --build



.PHONY: composeup composedown applogs dblogs deletedbdata swag run stop swagabsolute runabsolute
