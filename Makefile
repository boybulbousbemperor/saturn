build-server:
#Note that the build output is explicitly set to web/app.wasm. 
# The reason why is that the Handler expects the client to be a static resource 
# -located at the /web/app.wasm path.
  cd server && GOARCH=wasm GOOS=js go build -o /~/saturn-js/web/app.wasm
	go build

run-server: build-server
	./main