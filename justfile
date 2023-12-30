build:
    @go build -o benthic

bin:
    GOOS=darwin GOARCH=amd64 go build -o benthic-mac
    GOOS=darwin GOARCH=arm64 go build -o benthic-mac-arm
    GOOS=windows GOARCH=amd64 go build -o benthic-windows.exe
    GOOS=linux GOARCH=amd64 go build -o benthic-linux