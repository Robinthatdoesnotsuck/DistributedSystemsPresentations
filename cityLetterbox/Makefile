rating-service:
	go run rating/cmd/main/main.go;
metadata-service:
	go run metadata/cmd/main/main.go;
movie-service:
	go run movie/cmd/main/main.go;

test-rating-service:

test-metadata-service:
	curl -X PUT -d id=1 -d localhost:8081/metadata
test-movie-service:
	curl -v localhost:8083/movie?id=1

run-consul-dev-server:
	docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul consul:1.15.4 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

generate-protobuf-code-only-for-model:
	protoc -I=api --go_out=. letterbox.proto

generate-protobuf-code-with-grpc:
	protoc -I=api --go_out=. --go-grpc_out=. letterbox.proto