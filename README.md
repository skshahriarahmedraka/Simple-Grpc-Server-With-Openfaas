### protoc installation (Protocol Buffer Compiler)
For Linux :
```
sudo apt install -y protobuf-compiler
```

### protoc-gen-go (the code generator) installation

We need to generate the server’s code and the messages’ encoding.

To do that, we will need a code generator :

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
```
sudo apt install golang-goprotobuf-dev
```
