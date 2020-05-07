
all:	
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w"

run:
	bee run -gendoc=true -downdoc=true

deploy:
	fun deploy -y

clean:
	rm go-serv*
