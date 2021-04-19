[![PkgGoDev](https://pkg.go.dev/badge/github.com/CortexFoundation/rosetta-cortex)](https://pkg.go.dev/github.com/CortexFoundation/rosetta-cortex)
[![license](https://img.shields.io/github/license/CortexFoundation/rosetta-cortex.svg)](https://github.com/CortexFoundation/rosetta-cortex/master/LICENSE)
[![LoC](https://tokei.rs/b1/github.com/CortexFoundation/rosetta-cortex)](https://github.com/CortexFoundation/rosetta-cortex)

# Cortex Rosetta Shared Lib

Shared lib used in Cortex-SDK. This includes the libraries that are used by
different versions of the SDK, including Launchpad and Stargate.

Cortex full node

https://github.com/CortexFoundation/CortexTheseus
```
./cortex --rpc --rpcapi "net,web3,debug,admin,personal,ctxc"
```

Starting the proxy service
```
go run cmd/main.go
```

Do testing

```
./test-insall.sh
./test-run.sh
```
![image](https://user-images.githubusercontent.com/22344498/115228567-853abf80-a144-11eb-9374-4c5611f31f6a.png)
