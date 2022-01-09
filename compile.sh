
#!/bin/bash

set -e

env GOOS=darwin go build -o ./build/main_darwin  main.go
env GOOS=linux go build -o ./build/main_linux  main.go
env GOOS=windows go build -o ./build/main_windows.exe  main.go