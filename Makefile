# build
build:
	GOOS=linux go build -o bin/main ./src

# deploy
deploy:
	GOOS=linux go build -o bin/main ./src && sls deploy -v --profile kzk-serverless