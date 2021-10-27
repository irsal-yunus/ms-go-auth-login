module github.com/irsal-yunus/ms-go-auth-login/authentication

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/irsal-yunus/ms-go-auth-login v0.0.0-20211027065908-bcdc6157551d // indirect
	github.com/irsal-yunus/ms-go-auth-login/rpc v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/irsal-yunus/ms-go-auth-login/rpc => ../rpc
