# sneat-mod-template

A template for starting a new Sneat module.

To customize the template, run the following command:

```shell
./customize.sh
```

This will do the following:

1. Detect the name of the module from the current directory name.
   The name of the repo should starts with `sneat-mod-` following by a module ID, like: `sneat-mod-{MODULE_ID}`.

2. Rename directory `go/module` to `go/{MODULE_ID}`.

3. Alter `go.mod` to use the module ID as module name. 

<!-- dev-approach:v1 -->
## Our approach to development

We build with our own tooling:

- **[SpecScore](https://specscore.md)** — specify requirements as `SpecScore.md` artifacts
- **[SpecStudio](https://specscore.studio)** — author & manage specs across their lifecycle
- **[inGitDB](https://ingitdb.com)** — store structured data in Git where applicable
- **[DALgo](https://dalgo.io)** — data access layer for Go
- **[cover100.dev](https://cover100.dev)** — drive toward 100% test coverage
- **[DataTug](https://datatug.io)** — query & explore data
<!-- /dev-approach -->
