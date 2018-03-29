install:
	go get -u github.com/dgrijalva/jwt-go

run: app.rsa
	go run main.go

keypair/openssl:
	openssl genrsa 4096 > app.rsa
	openssl rsa -pubout < app.rsa > app.rsa.pub

app.rsa:
	$(MAKE) keypair/openssl

clean:
	rm -f app.rsa app.rsa.pub
