# run swagger
on MacOs
export GOPATH=/Users/(username)/go
export PATH=$GOPATH/bin:$PATH
swag init 

if you want to generate continously 
swag init -g {path router}

report bug
1. jika type input string, tapi diisi integer masih lolos validatenya (done)
2. login, register, socialmedia, photo, comment v2 is done (pending to test)