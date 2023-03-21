PROTOBUF_INPUT_DIR=.\RPC\protobuf
PROTOBUF_OUTPUT_DIR=.\RPC
allpbf: clipboard.proto

clipboard.proto:
				protoc -I=${PROTOBUF_INPUT_DIR}\clipboard --go_out=${PROTOBUF_OUTPUT_DIR} ${PROTOBUF_INPUT_DIR}\clipboard\clipboard.proto