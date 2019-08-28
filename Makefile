run-hal:
	go run ./cmd/hal/main.go

run-anton:
	go run ./cmd/anton/main.go

run-ui-server: 
	cd ./ui-server && npm run dev
