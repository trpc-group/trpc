DIR=pb/go
all: 
	rm -rf $(DIR)
	mkdir -p $(DIR)
	protoc -I/usr/local/include -I.\
		--go_out=paths=source_relative:./$(DIR) \
		./trpc/api/annotations.proto \
		./trpc/api/http.proto 
	protoc -I/usr/local/include -I.\
		--go_out=./$(DIR) \
		./trpc/proto/trpc.proto \
		./trpc/proto/swagger.proto \
		./trpc/proto/trpc_options.proto \
		./trpc/proto/validate.proto 
	mv $(DIR)/git.woa.com/trpc/trpc-protocol/pb/go/trpc/* $(DIR)/trpc
	rm -rf $(DIR)/git.woa.com
	cd ./$(DIR)/trpc && go mod init git.woa.com/trpc/trpc-protocol/$(DIR)/trpc && go mod tidy && cd - 

clean:
	rm -rf $(DIR)
