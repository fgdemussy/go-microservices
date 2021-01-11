check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	swagger generate spec --scan-models -o ./swagger.yaml