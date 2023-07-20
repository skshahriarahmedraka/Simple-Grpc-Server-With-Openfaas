package function

import (
	"fmt"
	server "github.com/skshahriarahmedraka/grpc-server"
)

// Handle a serverless request
func Handle(req []byte) string {

	server.Server()
	return fmt.Sprintf("Hello, Go. You said: %s", string(req))
}
