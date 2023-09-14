#!/usr/bin/env bash

current_dir=$(basename "$(pwd)")
MODULE_ID="${current_dir#sneat-module-}"
echo "Module id: $MODULE_ID"
mv go/module go/"$MODULE_ID"
sed -i '' "s/sneat-module-template/sneat-mod-$MODULE_ID/g" go.mod
find go/module -type d -name "*module" -exec bash -c 'mv "$1" "${1%module}$MODULE_ID"' bash {} \;
find "go" -type f -name "*.go" -exec sed -i '' "s/package module/package $MODULE_ID/g" {} +
find "go" -type f -name "*.go" -exec sed -i '' "s/{MODULE_ID}/$MODULE_ID/g" {} +
rm customize.sh
