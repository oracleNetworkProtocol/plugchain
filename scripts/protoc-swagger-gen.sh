#!/usr/bin/env bash

set -eo pipefail

LIQUIDITY_VERSION=v0.2.1

mkdir -p ./tmp-swagger-gen

proto_dirs=$(find ./proto ./third_party/proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do

  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 -name 'query.proto')

  if [[ ! -z "$query_file" ]]; then
    buf protoc  \
        -I "proto" \
        -I "third_party/proto" \
      "$query_file" \
      --swagger_out=./tmp-swagger-gen \
      --swagger_opt=logtostderr=true --swagger_opt=fqn_for_swagger_name=true --swagger_opt=simple_operation_ids=true
  fi
done

#copy liquidity swagger.yml
chmod -R 755 ${GOPATH}/pkg/mod/github.com/oracle!network!protocol/liquidity@${LIQUIDITY_VERSION}/tmp-swagger-gen/tendermint/liquidity/v1beta1/query.swagger.json
mkdir -p  ./tmp-swagger-gen/tendermint/liquidity/v1beta1
cp -r ${GOPATH}/pkg/mod/github.com/oracle!network!protocol/liquidity@${LIQUIDITY_VERSION}/tmp-swagger-gen/tendermint/liquidity/v1beta1/query.swagger.json ./tmp-swagger-gen/tendermint/liquidity/v1beta1/query.swagger.json


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
sed -i '' 's/Gaia/Plug Chain Hub/g' ./client/static/openapi.yml
sed -i '' 's/gaia/plugchaind/g' ./client/static/openapi.yml
sed -i '' 's/cosmoshub/plugchainhub/g' ./client/static/openapi.yml
sed -i '' 's/uatom/uplugcn/g' ./client/static/openapi.yml
sed -i '' 's/atom/plugcn/g' ./client/static/openapi.yml

tendermintURL=https://github.com/tendermint/liquidity/blob/develop
onpURL=https://github.com/oracleNetworkProtocol/liquidity/tree/main
sed -i '' "s#${tendermintURL}#${onpURL}#g" ./client/static/openapi.yml

# generate proto doc  Use tools for protoc-gen-doc

buf protoc \
    -I "proto" \
    -I "third_party/proto" \
    --doc_out=./docs/endpoints \
    --doc_opt=./docs/endpoints/protodoc-markdown.tmpl,proto-docs.md \
    $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')
go mod tidy

cp ./docs/endpoints/proto-docs.md ./docs/zh/endpoints/proto-docs.md

# clean swagger files
rm -rf ./tmp-swagger-gen