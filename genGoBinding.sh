#!/bin/sh

base_dir="abigenBindings"
contracts=$(find $base_dir/abi -mindepth 1 -maxdepth 1 -exec basename {} \;)

for contract in $contracts; do
    echo $contract
    abigen --abi $base_dir/abi/$contract --pkg smc --out $base_dir/golang/${contract%.*}.go
done

