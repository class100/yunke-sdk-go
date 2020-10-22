module github.com/class100/yunke-sdk-go

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-resty/resty/v2 v2.3.0
	github.com/jlaffaye/ftp v0.0.0-20201021201046-0de5c29d4555 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/storezhang/gox v1.2.6
	github.com/storezhang/replace v1.0.3
	github.com/storezhang/transfer v1.0.1
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20201020230747-6e5568b54d1a // indirect
)

// replace github.com/storezhang/gox => ../../storezhang/gox
