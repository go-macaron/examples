run:
	go build -o sqlx && ./sqlx

deps:
	go get -v gopkg.in/macaron.v1
	go get -v github.com/mattn/go-sqlite3
	go get -v github.com/jmoiron/sqlx

clean:
	go clean -i ./...
	rm songs.db