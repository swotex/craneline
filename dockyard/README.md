# DB and API

make

docker compose up -d db

make migrate-up

make sqlc-generate

make docker-up