module github.com/class100/yunke-sdk-go

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-resty/resty/v2 v2.3.0
	github.com/json-iterator/go v1.1.10
	github.com/storezhang/gox v1.2.5
	github.com/storezhang/replace v1.0.1
	github.com/storezhang/transfer v1.0.0
)

replace github.com/storezhang/gox => ../../storezhang/gox
