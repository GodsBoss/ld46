all: dist/main.wasm dist/wasm_exec.js dist/index.html dist/gfx.png

dist/main.wasm: Makefile
	GOOS=js GOARCH=wasm go build -o $@

dist/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js dist
	cp $< $@

dist/index.html: static/index.html
	cp $< $@

dist/gfx.png: gfx/gfx.xcf Makefile gfx/gfx.sh
	gfx/gfx.sh $< $@

dist:
	mkdir -p dist

clean:
	rm -rf dist

serve: all
	SERVE_DIR=$${PWD}/dist go run ./serve

.PHONY: \
	all \
	clean \
	dist/main.wasm \
	serve
