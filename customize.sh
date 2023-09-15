#!/usr/bin/env bash

current_dir=$(basename "$(pwd)")
MODULE_ID="${current_dir#sneat-module-}"
echo "Module id: $MODULE_ID"

# ======== BEGIN: Customization of Go related files & directories ========

# Renames module name in go.mod file using actual module id taken from $MODULE_ID environment variable
sed -i '' "s/sneat-mod-module/sneat-mod-$MODULE_ID/g" go.mod

# Renames module sub-directories using actual module id taken from $MODULE_ID environment variable
for dir in go/module/*module; do mv "$dir" "${dir%module}$MODULE_ID"; done

# Replace 4module with actual module id taken from $MODULE_ID environment variable
find "go/module" -type f -exec sed -i '' "s/4module/4$MODULE_ID/g" {} +


# Changes package name "module" to actual module id taken from $MODULE_ID environment variable
find "go/module" -type f -name "*.go" -exec sed -i '' "s/package module/package $MODULE_ID/g" {} +

# Replace {MODULE_ID} with actual module id taken from $MODULE_ID environment variable
find "go/module" -type f -name "*.go" -exec sed -i '' "s/{MODULE_ID}/$MODULE_ID/g" {} +


# Renames module directory using actual module id taken from $MODULE_ID environment variable
mv go/module go/"$MODULE_ID"

# ======== END: Customization of Go related files & directories ========

# ======== BEGIN: Clean-up ========

# Replace README.md with README-MODULE.md
mv README-MODULE.md README.md

# Remove customization script
rm customize.sh

git add .
# ======== END: Clean-up ========
