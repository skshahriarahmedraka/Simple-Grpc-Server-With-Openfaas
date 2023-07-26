### NOTE:  Openfaas Doesn't support Grpc protocol yet



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

### First locally test client and server

run server 

```
go run server/main.go
```

run client 

```
go run client/main.go
```

# Method 1 : Using Docker

1. build docker image of server
   inside `/server`
   
   ```
   sudo docker build -t shahriarraka/simple-grpc . 
   ```

2. push the contianer to docker hub
   
   ```
   sudo docker push shahriarraka/simple-grpc:latest
   ```

3. create a deployment file for openfaas

![](./screenshort/Screenshot%20from%202023-07-20%2011-38-42.png)
4. see the log

![](./screenshort/Screenshot%20from%202023-07-20%2011-34-20.png)

# Method 2: using Openfaas template

### built  a github package (grpc-server) for running there server

```
https://github.com/skshahriarahmedraka/grpc-server
```

## Run OpenFass using this instructions

```
https://mickey.dev/posts/getting-started-with-openfaas/
```

## build the function for deploying openfaas

1. created a deploy server config file (simple-openfass-func.yml) for openfaas

2. build the function using this command 
   
   ```
   sudo faas-cli new  --lang go server
   ```

3. build the deployment file using this command 
   
   ```
   sudo faas-cli build -f simple-openfaas-func.yml
   ```

4. push the docker image to docker hub using this command 
   
   ```
   sudo faas-cli push -f simple-openfaas-func.yml
   ```

5. login to openfaas using this command 
   
   ```
   export PASSWORD=mysecretpassword 
   echo -n $PASSWORD | faas-cli login --username admin --password-stdin  --gateway 127.0.0.1:31112 
   ```

6. deploy the function using this command 
   
   ```
   sudo faas-cli deploy -f simple-openfaas-func.yml --gateway 127.0.0.1:31112 
   ```

![](./screenshort/Screenshot%20from%202023-07-20%2010-19-36.png)

![](./screenshort/Screenshot%20from%202023-07-20%2010-20-13.png)

![](./screenshort/Screenshot%20from%202023-07-20%2010-20-30.png)

![](./screenshort/Screenshot%20from%202023-07-20%2010-21-25.png)

![](./screenshort/Screenshot%20from%202023-07-20%2010-21-48.png)

![](./screenshort/Screenshot%202023-07-20%20at%2010-18-20%20OpenFaaS%20Portal.png)

![](./screenshort/Screenshot%20from%202023-07-20%2010-44-06.png)

For reference :

https://www.openfaas.com/blog/golang-deep-dive/

https://gcore.com/learning/create-serverless-functions-with-openfaas/

https://mickey.dev/posts/getting-started-with-openfaas/

https://docs.openfaas.com/cli/templates/
