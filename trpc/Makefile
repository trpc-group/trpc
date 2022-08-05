all: trpc.proto
	protoc --go_out=paths=source_relative:. trpc.proto swagger.proto

clean:
	rm *.pb.go
