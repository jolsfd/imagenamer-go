export GOARCH=amd64

export GOOS=darwin
go build -o ./bin/mac/imagenamer-go .

export GOOS=linux
go build -o ./bin/linux/imagenamer-go .

export GOOS=windows
go build -o ./bin/windows/imagenamer-go.exe .