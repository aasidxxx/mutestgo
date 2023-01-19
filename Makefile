.PHONY: build
run:
	go run ./cmd/bot/main.go
#test:
#	echo testttt
build:
	go build -v ./cmd/bot/

#runpath:
#	go run ./cmd/api/ -path configs/api.toml
#
#migrateup:
#	migrate -path migrations -database "mysql://root:rootroot@tcp(localhost:3306)/goapi" up
#
#migratedown:
	#migrate -path migrations -database "mysql://root:rootroot@tcp(localhost:3306)/goapi" down