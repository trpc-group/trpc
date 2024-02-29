DIR=pb/go
INCLUDE_GOOGLE_PB_PATH=/usr/local/include
all: 
	rm -rf $(DIR)
	mkdir -p $(DIR)
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=paths=source_relative:./$(DIR) \
		./trpc/api/annotations.proto \
		./trpc/api/http.proto \
		./trpc/api/openapiv2.proto 
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=./$(DIR) \
		./trpc/trpc.proto \
		./trpc/proto/trpc_options.proto \
		./trpc/swagger/swagger.proto \
		./trpc/validate/validate.proto 
	mv $(DIR)/git.woa.com/trpc/trpc-protocol/pb/go/trpc/* $(DIR)/trpc
	rm -rf $(DIR)/git.woa.com
	cd ./$(DIR)/trpc && go mod init git.woa.com/trpc/trpc-protocol/$(DIR)/trpc && go mod tidy && cd - 

	# trpc v2
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=paths=source_relative:./$(DIR) \
		./trpc/v2/api/annotations.proto \
		./trpc/v2/api/http.proto 
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=./$(DIR) \
		./trpc/v2/trpc.proto \
		./trpc/v2/proto/trpc_options.proto \
		./trpc/v2/swagger/swagger.proto \
		./trpc/v2/validate/validate.proto 
	mv $(DIR)/git.woa.com/trpc/trpc-protocol/pb/go/trpc/v2/* $(DIR)/trpc/v2
	rm -rf $(DIR)/git.woa.com
	cd ./$(DIR)/trpc/v2 && go mod init git.woa.com/trpc/trpc-protocol/$(DIR)/trpc/v2 && go mod tidy && cd - 

.PHONY: test clean

test:
	./testbuild.sh

clean:
	rm -rf $(DIR)
