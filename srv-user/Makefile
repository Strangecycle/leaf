# target: prerequisite 目标文件(target)依赖于前提文件(prerequisite)
.PROTOC: proto
proto:
	protoc --proto_path=./proto --micro_out=./proto/out --go_out=./proto/out ./proto/*.proto

.BUILD: build
build: proto
	go build -o leaf-user *.go
