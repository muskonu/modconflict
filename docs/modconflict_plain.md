## modconflict plain

Print out the chain of possible dependency conflicts.

### Synopsis

Print out the chain of possible dependency conflicts.

```
modconflict plain [flags]
```

### Options

```
  -h, --help   help for plain
```

### Options inherited from parent commands

```
  -f, --file string   select the file where the results of the go mod graph are saved as input
```

### example
```shell
//use pipe as input
go mod graph | modconflict plain
//use file as input
modconflict plain -f a.txt
```
#### output
```shell
find Confict in package github.com/golang/protobuf:
demo -> github.com/golang/protobuf@v1.5.4
demo -> github.com/grpc-ecosystem/grpc-gateway/v2@v2.18.0 -> github.com/golang/protobuf@v1.5.3
```

### SEE ALSO

* [modconflict](modconflict.md)	 - Tool to quickly locate paths with conflicting dependencies.