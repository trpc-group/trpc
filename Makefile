DIR=pb/go
INCLUDE_GOOGLE_PB_PATH=/usr/local/include
all: clean_dir create_dir trpc_api trpc_reflection trpc_core trpc_v2_api trpc_v2_core 

clean_dir:
	rm -rf $(DIR)

create_dir:
	mkdir -p $(DIR)

trpc_api:
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=paths=source_relative:./$(DIR) \
		./trpc/api/annotations.proto \
		./trpc/api/http.proto \
		./trpc/api/openapiv2.proto 

trpc_reflection:
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=paths=source_relative:./$(DIR) \
		./trpc/reflection/reflection.proto
	cd ./$(DIR)/trpc/reflection && rm -f go.mod go.sum && go mod init trpc.group/trpc/trpc-protocol/$(DIR)/trpc/reflection && go mod edit -go=1.18 && go mod tidy && cd -

trpc_core:
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=./$(DIR) \
		./trpc/trpc.proto \
		./trpc/proto/trpc_options.proto \
		./trpc/swagger/swagger.proto \
		./trpc/validate/validate.proto
	mv $(DIR)/trpc.group/trpc/trpc-protocol/pb/go/trpc/* $(DIR)/trpc
	rm -rf $(DIR)/trpc.group
	cd ./$(DIR)/trpc && go mod init trpc.group/trpc/trpc-protocol/$(DIR)/trpc && go mod tidy && cd -

trpc_v2_api:
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=paths=source_relative:./$(DIR) \
		./trpc/v2/api/annotations.proto \
		./trpc/v2/api/http.proto 

trpc_v2_core:
	protoc -I$(INCLUDE_GOOGLE_PB_PATH) -I.\
		--go_out=./$(DIR) \
		./trpc/v2/trpc.proto \
		./trpc/v2/proto/trpc_options.proto \
		./trpc/v2/swagger/swagger.proto \
		./trpc/v2/validate/validate.proto 
	mv $(DIR)/trpc.group/trpc/trpc-protocol/pb/go/trpc/v2/* $(DIR)/trpc/v2
	rm -rf $(DIR)/trpc.group
	cd ./$(DIR)/trpc/v2 && go mod init trpc.group/trpc/trpc-protocol/$(DIR)/trpc/v2 && go mod tidy && cd - 

.PHONY: test clean

test:
	./testbuild.sh

clean:
	rm -rf $(DIR)
