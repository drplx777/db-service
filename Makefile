migrations_diff:
	~/.local/share/go/bin/gormite --tool=goose --config=gormite.yaml --dsn="postgresql://postgres:postgres@127.0.0.1:5432/tasker?timezone=UTC&sslmode=disable"