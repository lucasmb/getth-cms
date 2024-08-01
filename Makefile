# Makefile
.PHONY: build

BINARY_NAME=getth-cms

# build builds the tailwind css sheet, and compiles the binary into a usable thing.
build:
	go mod tidy && \
   	templ generate && \
	go generate && \
	CGO_ENABLED=1 go build -ldflags="-w -s" -o ${BINARY_NAME}

# dev runs the development server where it builds the tailwind css sheet,
# and compiles the project whenever a file is changed.
dev:
	docker compose -f docker-compose.development.yml up

clean:
	go clean

# run tailwind watch through docker for local dev
tailwind:
	docker exec -i getth-cms npx tailwindcss -i ./assets/main.css -o ./assets/tailwind.css --watch