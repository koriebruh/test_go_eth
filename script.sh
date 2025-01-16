#CONNECTION
  go get github.com/ethereum/go-ethereum
  # if pengen pake localconnection usage ganache
  npm install -g ganache #first time aja kalo belum install

  #untuk menjalankan ganache nya
  ganache


# NOTE IF I BUY MAC I SHOULD SETUP FOR GENERATE BIN AND ABI
# install
# solidity compiler
# protobuf compiler
# last ->
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
$env:Path += ";$env:USERPROFILE\go\bin"


# script for generate bin and abi
npx solcjs --bin --abi -o build ./contract/lendingContract.sol

#generate mdoule
abigen --bin=build/contract_lendingContract_sol_LendingContract.bin \
       --abi=build/contract_lendingContract_sol_LendingContract.abi \
       --pkg=lendingContract \
       --out=gen/lending_contract.go

## misal udah di set alias / jika menjalnkan di linux or mac window busuk
#solc --bin --abi contract/lendingContract.sol -o build
