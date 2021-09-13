#!/usr/bin/env bash

set -eo pipefail

SDK_VERSION=v0.42.7
#TENDERMINT_VERSION=v0.34.11
chmod -R 755 ${GOPATH}/pkg/mod/github.com/cosmos/cosmos-sdk@${SDK_VERSION}/proto
chmod -R 755 ${GOPATH}/pkg/mod/github.com/cosmos/cosmos-sdk@${SDK_VERSION}/third_party/proto

rm -rf ./tmp-swagger-gen ./tmp
mkdir -p ./tmp-swagger-gen ./tmp/proto ./tmp/third_party

cp -r ${GOPATH}/pkg/mod/github.com/cosmos/cosmos-sdk@${SDK_VERSION}/proto ./tmp
cp -r ${GOPATH}/pkg/mod/github.com/cosmos/cosmos-sdk@${SDK_VERSION}/third_party/proto ./tmp/third_party
cp -r ./proto ./tmp

proto_dirs=$(find ./tmp/proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
echo $dir
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'msg.proto' -name 'tx.proto' \))
  if [[ $dir =~ "cosmos" ]]; then
   query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  fi
  if [[ ! -z "$query_file" ]]; then
    buf protoc  \
        -I "tmp/proto" \
        -I "tmp/third_party/proto" \
      "$query_file" \
      --swagger_out=./tmp-swagger-gen \
      --swagger_opt=logtostderr=true --swagger_opt=fqn_for_swagger_name=true --swagger_opt=simple_operation_ids=true
  fi
done


# copy cosmos swagger_legacy.yaml
# chmod -R 755 ${GOPATH}/pkg/mod/github.com/cosmos/cosmos-sdk@${SDK_VERSION}/client/docs/swagger_legacy.yaml
# cp -r ${GOPATH}/pkg/mod/github.com/cosmos/cosmos-sdk@${SDK_VERSION}/client/docs/swagger_legacy.yaml ./client/static/cosmos_swagger_legacy.yml

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine  ./client/config.json  -o ./client/static/openapi.yml -f yml --continueOnConflictingPaths true --includeDefinitions true

# a=`uname  -a`
# b="Darwin"
# if [[ $a =~ $b ]]
# then
# SedP="''"
# else
# SedP=""
# fi

# replace APIs example
sed -r -i '' 's/cosmos1[a-z,0-9]+/gx1sltcyjm5k0edlg59t47lsyw8gtgc3nudklntcq/g' ./client/static/openapi.yml
sed -r -i '' 's/cosmosvaloper1[a-z,0-9]+/gxvalioper1sltcyjm5k0edlg59t47lsyw8gtgc3nudrwey98/g' ./client/static/openapi.yml
sed -r -i '' 's/cosmosvalconspub1[a-z,0-9]+/gxvalconspub1zcjduepqwhwqn4h5v6mqa7k3kmy7cjzchsx5ptsrqaulwrgfmghy3k9jtdzs6rdddm/g' ./client/static/openapi.yml
sed -i '' 's/Gaia/PLUGChainHub/g' ./client/static/openapi.yml
sed -i '' 's/gaia/plugchaind/g' ./client/static/openapi.yml
sed -i '' 's/cosmoshub/plugchainhub/g' ./client/static/openapi.yml


# generate proto doc  Use tools for protoc-gen-doc

buf protoc \
    -I "tmp/proto" \
    -I "tmp/third_party/proto" \
    --doc_out=./docs/endpoints \
    --doc_opt=./docs/endpoints/protodoc-markdown.tmpl,proto-docs.md \
    $(find "$(pwd)/tmp/proto" -maxdepth 5 -name '*.proto')
go mod tidy

cp ./docs/endpoints/proto-docs.md ./docs/zh/endpoints/proto-docs.md

# clean swagger files
rm -rf ./tmp-swagger-gen

# clean proto files
rm -rf ./tmp
