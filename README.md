# level-seven-matcha

`level-seven-matcha` is a simple Dockerized Postgres example to demonstrate database initialization. In the `init` folder, you can find sample database initialization scripts. `level-seven-matcha` connects to PostgreSQL using CRUD operations, powered along with [gorm](https://github.com/jinzhu/gorm) and driver for `database/sql` [pq](https://github.com/lib/pq).

## Initial Setup

Install Go and [Docker](https://hub.docker.com/editions/community/docker-ce-desktop-mac/)

## Usage

Run in console (cached):
```
$ docker-compose -f docker-compose.yml up
```

Run in detached mode:
```
$ docker-compose -f docker-compose.yml up -d
```

Run with new build:
```
$ docker-compose down && docker-compose build --no-cache && docker-compose up
```

Stop and clean up volumes:
```
$ docker-compose down -v
```
