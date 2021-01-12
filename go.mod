module github.com/bianjieai/iritamod

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.40.0-rc3
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/keybase/go-keychain v0.0.0-20191114153608-ccd67945d59e // indirect
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc6
	github.com/tendermint/tm-db v0.6.3
	github.com/tjfoc/gmsm v1.4.0
	google.golang.org/genproto v0.0.0-20201111145450-ac7456db90a6
	google.golang.org/grpc v1.33.2
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.28.2-0.20210112055458-b53a7d5a7c9c
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-210104.0.20210112015006-57e95aa6402f
)
