module github.com/leapkit/template

go 1.22.0

toolchain go1.22.1

require (
	github.com/leapkit/core v1.2.15
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/paganotoni/tailo v1.0.3
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gofrs/uuid/v5 v5.0.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	golang.org/x/sys v0.16.0 // indirect
)

replace github.com/leapkit/core => ../core
