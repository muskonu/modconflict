## modconflict graph

Generate images of potentially conflicting dependency routes.

### Synopsis

Generate images of potentially conflicting dependency routes.

```
modconflict graph [flags]
```

### Options

```
  -h, --help           help for graph
  -o, --ouput string   the name of the output file, with different formats depending on the suffix.
```

### Options inherited from parent commands

```
  -f, --file string   select the file where the results of the go mod graph are saved as input
```

### example
```shell
//use pipe as input
go mod graph | modconflict graph -o demo.svg
//use file as input
modconflict graph -f a.txt
```

### SEE ALSO

* [modconflict](modconflict.md)	 - Tool to quickly locate paths with conflicting dependencies.