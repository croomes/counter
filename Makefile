
# .PHONY build
build:
	go get github.com/mitchellh/gox
	gox -verbose -output="release/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		-osarch="linux/amd64 linux/arm darwin/amd64 windows/amd64"
