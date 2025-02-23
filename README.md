## LeapKit Template

This is the LeapKit template for building web applications with Go, HTMX and Tailwind CSS.

### Getting Started
#### Setup tools

To get going make sure you have Go 1.24 installed. Once you have Go installed, you can download the required dependencies:

```sh
go mod download
```

#### Running the application

To run the application while in development please use the following command:

```sh
go tool dev
```

This will use the `dev` tool to read the Procfile at the root of the project and start the application. It will automatically restart the app process when changes are detected in the .go files. The rest of the processes defined in the file will be running in parallel.
Once the application is running, you can access it at http://localhost:3000.


### Building the application

One important part of the build process is to build the TailwindCSS styles. The tailo tool will take care of this.

```sh
go tool tailo -i internal/system/assets/tailwind.css -o internal/system/assets/application.css
```

Then building the application can be done with the following command:

```sh
go build -o bin/app ./cmd/app
```

That will create a binary file named `app` in the `bin` folder at the root of the project.
