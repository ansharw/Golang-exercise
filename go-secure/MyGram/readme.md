# run swagger
on MacOs
export GOPATH=/Users/(username)/go
export PATH=$GOPATH/bin:$PATH
swag init 

if you want to generate continously 
swag init -g {path router}