build:
    tinygo flash -target=pico2-w .

run: build
    tinygo monitor

editor device:
    export GOROOT=$(tinygo info -target={{ device }} | grep 'cached GOROOT' | cut -d: -f2 | xargs); \
    export GOFLAGS="-tags=$(tinygo info -target={{ device }} | grep 'build tags' | cut -d: -f2 | xargs | tr ' ' ',')"; \
    zeditor .
