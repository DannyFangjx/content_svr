
#!/bin/bash

# 递归查找所有的 .proto 文件
project=content_svr
PROTO_FILES=$(find ./$project/protobuf -name "*.proto")

# 递归编译所有的 .proto 文件，并生成对应的 .pb.go 文件
for file in ${PROTO_FILES}; do
    # 获取文件名（不包含路径和扩展名）
    file_name=$(basename -- "${file}" .proto)
    # 获取文件所在目录
    dir_name=$(dirname -- "${file}")
    # 生成 .pb.go 文件的路径
    pb_file_path="${dir_name}/${file_name}.pb.go"
    # 编译 .proto 文件并生成 .pb.go 文件
    protoc --go_out=. "${file}" && echo "Generated ${pb_file_path}"
done