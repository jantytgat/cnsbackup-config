build:
	bash scripts/build.sh

clean:
	bash scripts/clean.sh

docker-build:
	docker build -t cnsbackup-config:dev-latest .

run:
	go run main.go

test:
	go test main_test.go