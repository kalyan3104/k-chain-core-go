//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/kalyan3104/protobuf/protobuf  --gogoslick_out=. validatorStatistics.proto
package validator
