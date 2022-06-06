
build:
	mkdir build
	GOOS=linux go build -o build/main .

pack: build
	zip -j build/main.zip build/main

.PHONY: deploy
deploy: build pack
	aws lambda update-function-code \
    	--function-name poc_diver \
    	--zip-file fileb://build/main.zip 

.PHONY: clean
clean:
	rm -rf build