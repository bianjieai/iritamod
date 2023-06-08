module github.com/bianjieai/iritamod

go 1.16

require (
	cosmossdk.io/simapp v0.0.0-20230426205644-8f6a94cd1f9f
	github.com/cometbft/cometbft v0.37.1
	github.com/cosmos/cosmos-sdk v0.47.2
	github.com/cosmos/gogoproto v1.4.8
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.14.0
	github.com/stretchr/testify v1.8.2
	github.com/tjfoc/gmsm v1.4.0
	golang.org/x/crypto v0.7.0
	google.golang.org/genproto v0.0.0-20230216225411-c8e22ba71e44
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0

)

replace (
	github.com/cometbft/cometbft => github.com/bianjieai/cometbft v0.37.1-irita-230607
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.34.4-0.20230608013811-adaeef8bb569
)
