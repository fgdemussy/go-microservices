check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	GO111MODULE=off swagger generate spec --scan-models -o ./swagger.yaml