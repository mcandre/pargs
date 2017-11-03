VERSION=0.0.3

.PHONY: port clean clean-ports

all: integration-test

integration-test:
	find shakespeare -type f -print | pargs -n 2 echo 'plays:'

govet:
	find . -path "*/vendor*" -prune -o -name "*.go" -type f -exec go tool vet -shadow {} \;

golint:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec golint {} \;

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

errcheck:
	errcheck -blank

nakedret:
	nakedret -l 0 ./...

bashate:
	find . \( -wholename '*/.git/*' -o -wholename '*/node_modules*' -o -name '*.bat' \) -prune -o -type f \( -wholename '*/lib/*' -o -wholename '*/hooks/*' -o -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs bashate

shlint:
	find . \( -wholename '*/.git/*' -o -wholename '*/node_modules*' -o -name '*.bat' \) -prune -o -type f \( -wholename '*/lib/*' -o -wholename '*/hooks/*' -o -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs shlint

checkbashisms:
	find . \( -wholename '*/.git/*' -o -wholename '*/node_modules*' -o -name '*.bat' \) -prune -o -type f \( -wholename '*/lib/*' -o -wholename '*/hooks/*' -o -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs checkbashisms -n -p

shellcheck:
	find . \( -wholename '*/.git/*' -o -wholename '*/node_modules*' -o -name '*.bat' \) -prune -o -type f \( -wholename '*/lib/*' -o -wholename '*/hooks/*' -o -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs shellcheck

editorconfig:
	flcl . | xargs -n 100 editorconfig-cli check

lint: govet golint gofmt goimport errcheck nakedret bashate shlint checkbashisms shellcheck editorconfig

port: archive-ports

archive-ports: bin
	zipc -C bin "pargs-$(VERSION).zip" "pargs-$(VERSION)"

bin:
	gox -output="bin/pargs-$(VERSION)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./cmd...

clean: clean-ports

clean-ports:
	rm -rf bin
