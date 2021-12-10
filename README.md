# Advent of code 2021


## Build
```
go build
./<name_of_module>
```

## Adding a module
``` bash
go mod init
go mod tidy
go mod edit -replace <name>=../<name>
go get <name>
```