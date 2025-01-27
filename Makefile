TARGET=codelab

all: fmt clean build

clean:
    rm -rf $(TARGET)

depends:
    go get -u -v

build:
    go build -v -o $(TARGET) main.go

fmt:
    go fmt ./...

test:
    go test -v ./...
run:
    go run main.go