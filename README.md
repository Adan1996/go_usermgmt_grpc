# BELAJAR GRPC MENGGUNAKAN GOLANG

## Membuat user management dengan GRPC
1. Inisialisasikan go module dengan perintah
> go mod init github.con/Adan1996/go_usermgmt_grpc
> go get google.golang.org/grpc

2. compile grpc dengan perintah
> protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. usermgmt/usermgmt.proto
