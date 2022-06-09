# ./build-abi.sh selfAuth SelfAuth
solc --abi --bin ../contracts/$1.sol -o ../contracts/build --overwrite
# abigen --abi=../contracts/build/$2.abi --bin=../contracts/build/$2.bin --pkg=contracts --type=SelfAuth --out=../contracts/$1.go
abigen --abi=../contracts/build/I$2.abi --bin=../contracts/build/I$2.bin --pkg=contracts --type=ISelfAuth --out=../contracts/$1.go
