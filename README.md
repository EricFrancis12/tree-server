# Tree Server

## Vision

The goal of this project will be to create an executable that a user can run in any directory. It will then spin up an http server and serve all files nested within per the tree structure of the directory.

The idea is that this will be a quick way to get a server up and running for a large amount of files.

## To Do:

- (DONE) build the basic application
- (DONE) add ability to pass in args to the executable to specify PORT and DIR_PATH
- Dockerize it

## Diagram:

```base

$ tree-server.exe --PORT=3000

|-- videos
    |-- movies
        |-- saving-private-ryan.mp4     -->     http://localhost:3000/videos/movies/saving-private-ryan.mp4
        |-- moana.mp4                   -->     http://localhost:3000/videos/movies/moana.mp4
    |-- tv-shows
        |-- how-i-met-your-mother.mp4   -->     http://localhost:3000/videos/tv-shows/how-i-met-your-mother.mp4
|-- pictures
    |-- family-photos
        |-- first-birthday.jpeg         -->     http://localhost:3000/pictures/family-photos/first-birthday.jpeg

```
