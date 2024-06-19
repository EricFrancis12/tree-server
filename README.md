# Tree Server
Serve Files at Lightning Speed according to the Tree Structure of your File System

## Diagram
```base

$ ./tree-server.exe --PORT=3000

|-- Documents
    |-- spreadsheets
        |-- third-quarter-report.csv    -->     http://localhost:3000/Documents/spreadsheets/third-quarter-report.csv
|-- Pictures
    |-- family-photos
        |-- first-birthday.jpeg         -->     http://localhost:3000/Pictures/family-photos/first-birthday.jpeg
    |-- screenshots
        |-- screenshot.png              -->     http://localhost:3000/Pictures/screenshots/screenshot.png

```

## Getting Started
Make sure you have [Go](https://go.dev/doc/install) installed.

Clone the repo:
```
git clone https://github.com/EricFrancis12/tree-server
```

```
cd tree-server
```

### Option A: Quickstart (Run the Application without building)
Run the following command, optionally specifying a port (PORT) and working directory (WD):
```
go run . --PORT=3000 --WD=./
```

In the example above, `--WD=./` means the contents of this repo will be served on port 3000.

### Option B: Build and Run
Run the following command:
```
go build -o tree-server .
```

This will create a binary file located at `./tree-server`.

Run the binary, optionally specifying a port (PORT) and working directory (WD).

Note that if no working directory is specified, the working directory will be the location where the binary is ran. This means you can copy the binary to anywhere in the file system to serve its contents.
```
./tree-server --PORT=3000 --WD=./
```

### Option C: Run via Docker
Make sure you have [Docker](https://docs.docker.com/engine/install) installed.

Build the Docker Image:
```
docker build -t tree-server .
```

Run it using the following command:
```
docker run -p 3000:3000 -v .:/app/serve --name tree-server tree-server
```

The contents of the `/app/serve` directory located inside the container will be served. To bind `/app/serve` to a directory on your local machine, use a bind mount. In the above command, we are binding the `/app/serve` directory to the current local directory `.`.

For more information on bind mounts, [click here](https://docs.docker.com/storage/bind-mounts/).

For more information on port mapping using the `-p` flag, [click here](https://docs.docker.com/network/#published-ports)

## Use the Application in the browser
Once the Application is running, visit it at your chosen port:

[http://localhost:3000](http://localhost:3000)

### Downloading files
Use `dl=1` or `download=1` to download the file if it exists.

Example: http://localhost:3000/my/file.txt?dl=1
