set -o allexport
source .env
set +o allexport

go run ./cmd/server/

