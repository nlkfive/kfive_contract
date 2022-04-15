#!/bin/sh
base_dir=$(pwd)/abigenBindings
mkdir -p $base_dir/golang
contracts=$(find $base_dir/abi -mindepth 1 -maxdepth 1 -exec basename {} \;)

for contract in $contracts; do
    echo $contract
    name=${contract%.*}
    mkdir -p $base_dir/golang/$name
    abigen --abi $base_dir/abi/$contract --pkg "$name" --out $base_dir/golang/$name/${name}.go
done

