module github.com/irsal-yunus/ms-go-auth-login/testGrpc

go 1.16

replace github.com/irsal-yunus/ms-go-auth-login/rpc => ../rpc

require (
	github.com/irsal-yunus/ms-go-auth-login/rpc v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
)