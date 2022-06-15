# ./build-abi.sh rsdefi SelfAuth
if [[ "$2" = "" ]] ;then
	solc --abi --bin ../contracts/$1.sol -o ../contracts/build --overwrite
	abigen --abi=../contracts/build/$1.abi --bin=../contracts/build/$1.bin --pkg=contracts --type=$1 --out=../contracts/$1.go
else
	solc --abi --bin ../contracts/$1/$2.sol -o ../contracts/$1/build --overwrite
	abigen --abi=../contracts/$1/build/$2.abi --bin=../contracts/$1/build/$2.bin --pkg=$1 --type=$2 --out=../contracts/$1/$2.go
fi
