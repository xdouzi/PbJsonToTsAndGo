#!/bin/bash
#chmod u+x *.sh
#rm -rf bin
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./Toos/json-gen-ts
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./Toos/json-gen-ts

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./Toos/json-gen-go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./Toos/json-gen-go

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./Toos/json-gen-java
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./Toos/json-gen-java

#CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build ./Toos/pb-gen-JsonToos
#CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build ./Toos/pb-gen-JsonToos

#输出路径bin
output_dir="./bin"

# 删除文件夹
if [ -d "$output_dir" ]; then
    echo "Deleting folder..."
    rm -rf "$output_dir"
else
    echo "Folder does not exist."
fi

# 创建文件夹
echo "Creating folder..."
mkdir "$output_dir"

cp json-gen-ts "$output_dir/json-gen-ts"
cp json-gen-go "$output_dir/json-gen-go"
cp json-gen-java "$output_dir/json-gen-java"

cp json-gen-ts "$output_dir/json-gen-ts.exe"
cp json-gen-go "$output_dir/json-gen-go.exe"
cp json-gen-java "$output_dir/json-gen-java.exe"

echo "cp Protol..."
cp -rf Protol "$output_dir/"

rm json-gen-ts
rm json-gen-go
rm json-gen-java

rm json-gen-ts.exe
rm json-gen-go.exe
rm json-gen-java.exe