#!/usr/bin/env bash

current_dir=$(basename "$(pwd)")
MODULE_ID="${current_dir#sneat-module-}"
echo "Module id: $MODULE_ID"
mv go/module go/"$MODULE_ID"
sed -i '' "s/sneat-module-template/sneat-mod-$MODULE_ID/g" go.mod
# rm customize.sh
