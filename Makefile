CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:root@localhost:5432/auth?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://postgres:1111@localhost:5432/auth?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/auth?sslmode=disable -verbose down


migrate_force:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/auth -verbose force 1

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table


build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

swag-init:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1