set GOOS=linux
set GOARCH=amd64
cd .\cmd\gin
go build -tags=jsoniter -o SixAdminServer

