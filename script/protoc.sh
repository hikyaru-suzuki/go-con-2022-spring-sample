#!/bin/sh

proto_file_dir=./proto
server_proto_dir=${proto_file_dir}/server

out_proto_dir=./pkg/domain/proto
mkdir -p ${out_proto_dir}

call_protoc() {
  protoc "$@" || exit $?
}

# 3. 1の生成物をもとにprotoc
server_enums_proto_file=proto/server/enums/enums.proto
server_master_proto_files=$(find ${server_proto_dir}/master -type f -name '*.proto')
server_transaction_proto_files=$(find ${server_proto_dir}/transaction -type f -name '*.proto')
call_protoc \
  --proto_path=${proto_file_dir} \
	--sample_out=gen_enum,gen_master,gen_transaction,paths=source_relative:. \
  ${server_enums_proto_file} ${server_master_proto_files} ${server_transaction_proto_files}

gen-file-remover
