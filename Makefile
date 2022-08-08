DIR=pbgo
all: 
	mkdir -p $(DIR)
	protoc --go_out=paths=source_relative:./$(DIR) \
		./trpc/proto/trpc.proto \
		./trpc/proto/swagger.proto \
		./trpc/proto/trpc_options.proto \
		./trpc/proto/validate.ext.proto \
		./trpc/api/annotations.proto \
		./trpc/api/http.proto 
	cd ./pbgo && go mod init git.woa.com/trpc/trpc-protocol/$(DIR) && go mod tidy && cd - 

clean:
	rm -rf $(DIR)
