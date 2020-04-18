all: dist/main.wasm dist/wasm_exec.js dist/index.html

dist/main.wasm: Makefile
	GOOS=js GOARCH=wasm go build -o $@

dist/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js dist
	cp $< $@

dist/index.html: static/index.html
	cp $< $@

dist:
	mkdir -p dist

clean:
	rm -rf dist

.PHONY: \
	all \
	clean \
	dist/main.wasm
