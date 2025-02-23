## LeapKit Template

This is the LeapKit template for building web applications with Go, HTMX and Tailwind CSS.

### Setup tools

```sh
go tool tailo download
go mod download
```

### Building the application
The first step is to build the TailwindCSS assets:

```sh
go tool tailo -i internal/system/assets/tailwind.css -o internal/system/assets/application.css
```

```sh
go build -tags osusergo,netgo -buildvcs=false -o bin/app ./cmd/app
```

### Running the application

To run the application in development mode execute:

```sh
go tool dev
```

And open `http://localhost:3000` in your browser.

### Generating a migration

```sh
> go tool db generate_migration [name]
```
