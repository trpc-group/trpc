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
		./trpc/trpc.proto \
		./trpc/proto/trpc_options.proto \
		./trpc/swagger/swagger.proto \
		./trpc/validate/validate.proto 
	mv $(DIR)/git.woa.com/trpc/trpc-protocol/pb/go/trpc/* $(DIR)/trpc
	rm -rf $(DIR)/git.woa.com
	cd ./$(DIR)/trpc && go mod init git.woa.com/trpc/trpc-protocol/$(DIR)/trpc && go mod tidy && cd - 

	# trpc v2
	protoc -I/usr/local/include -I.\
		--go_out=paths=source_relative:./$(DIR) \
		./trpcv2/api/annotations.proto \
		./trpcv2/api/http.proto 
	protoc -I/usr/local/include -I.\
		--go_out=./$(DIR) \
		./trpcv2/trpc.proto \
		./trpcv2/proto/trpc_options.proto \
		./trpcv2/swagger/swagger.proto \
		./trpcv2/validate/validate.proto 
	mv $(DIR)/git.woa.com/trpc/trpc-protocol/pb/go/trpcv2/* $(DIR)/trpcv2
	rm -rf $(DIR)/git.woa.com
	cd ./$(DIR)/trpcv2 && go mod init git.woa.com/trpc/trpc-protocol/$(DIR)/trpcv2 && go mod tidy && cd - 

.PHONY: test clean

test:
	./testbuild.sh

clean:
	rm -rf $(DIR)
