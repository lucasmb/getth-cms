
This project is a dockerized TECHGO Stack (Tailwind, Echo, Caddy, HTMX, Go)

the docker image provides NODE Lts and GO 1.22

## how to use it

You'll need docker and docker compose installed on your host machine

````
make dev
````
This command will install all dependencies and run the server


to watch tailwind changes

````
make tailwind
````

SITE 
https://localhost:2015/