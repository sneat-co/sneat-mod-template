# sneat-module-template

A template for starting a new Sneat module.

To customize the template, run the following command:

```shell
./customize.sh
```

This will do the following:

1. Detect the name of the module from the current directory name.
   The name of the repo should starts with `sneat-mod-` following by a module ID, like: `sneat-module-{MODULE_ID}`.

2. Rename directory `go/module` to `go/{MODULE_ID}`.

3. Alter `go.mod` to use the module ID as module name. 