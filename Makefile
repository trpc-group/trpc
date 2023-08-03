DIR=pb/go
INCLUDE_GOOGLE_PB_PATH=/usr/local/include
all: 
	rm -rf $(DIR)
	mkdir -p $(DIR)
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=paths=source_relative:./$(DIR) \
		./trpc/api/annotations.proto \
		./trpc/api/http.proto 
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=./$(DIR) \
		./trpc/trpc.proto \
		./trpc/proto/trpc_options.proto \
		./trpc/swagger/swagger.proto \
		./trpc/validate/validate.proto 
	mv $(DIR)/trpc.group/trpc/trpc-protocol/pb/go/trpc/* $(DIR)/trpc
	rm -rf $(DIR)/trpc.group
	cd ./$(DIR)/trpc && go mod init trpc.group/trpc/trpc-protocol/$(DIR)/trpc && go mod tidy && cd -

.PHONY: test clean

test:
	./testbuild.sh

clean:
	rm -rf $(DIR)
