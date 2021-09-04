# go-fundementals

Learn Go in Docker

https://www.youtube.com/watch?v=jpKysZwllVw&t=519s

`docker run -it -v ${PWD}:/work go sh` --- Volume mount associated with the WORKDIR in Dockerfile

`go build` --> `./app`

`go install <module>` ---> stored in your $HOME/go/bin/app and can be run as a command line tool

Variables must first be instantiated:

1. var <variable> = "string"
2. <variable> := "string"

Control flows, decision trees, business logic:

Infinite loops to run batch jobs over certain intervals

## Docker target

`docker run --target dev` --- will only build the top two lines in your dockerfile