all: wasm serve

wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

serve:
	python3 -m http.server 8090

clean:
	rm -f main.wasm
	rm -rf code.txt

watch:
	reflex -r '\.go$$' -- bash -c "make wasm"


dev:
	make wasm
	reflex -r '\.go$$' -- bash -c "make wasm" &
	python3 -m http.server 8090


list_files:
	./utils/list_files.sh > code.txt

wasm-exec:
	bash utils/copy_wasm_exec.sh

