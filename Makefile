.PHONY: build compile

ZIP := exago-project-runner.zip
GOROOT := ./bin/linux-amd64/goroot
RUNNER_PATH := $(GOPATH)/src/github.com/hotolab/exago-runner/cmd/exago-runner

build: cleanup
	@echo "Adding directories to zipfile"
	zip -r9 $(ZIP) bin/ git/ index.js

# Setup build dependencies (not related to project)
cleanup:
	@echo "Cleaning up directories"
	@find . -name ".DS_Store" -exec rm {} \;
	@rm -f $(ZIP)

strip:
	@echo "Stripping Go from unecessary files"
	@rm -fR "$(GOROOT)/pkg/bootstrap" \
		"$(GOROOT)/pkg/obj" \
		"$(GOROOT)/pkg/linux_amd64_shared" \
		"$(GOROOT)/pkg/tool/linux_amd64/addr2line" \
		"$(GOROOT)/pkg/tool/linux_amd64/api" \
		"$(GOROOT)/pkg/tool/linux_amd64/cgo" \
		"$(GOROOT)/pkg/tool/linux_amd64/dist" \
		"$(GOROOT)/pkg/tool/linux_amd64/doc" \
		"$(GOROOT)/pkg/tool/linux_amd64/fix" \
		"$(GOROOT)/pkg/tool/linux_amd64/nm" \
		"$(GOROOT)/pkg/tool/linux_amd64/objdump" \
		"$(GOROOT)/pkg/tool/linux_amd64/pack" \
		"$(GOROOT)/pkg/tool/linux_amd64/pprof" \
		"$(GOROOT)/pkg/tool/linux_amd64/trace" \
		"$(GOROOT)/pkg/tool/linux_amd64/yacc"

compile:
	@echo "Compiling dependencies"
	cd $(RUNNER_PATH) && CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -i -tags netgo
	mv "$(RUNNER_PATH)/exago-runner" /var/task/bin/linux-amd64/ && cd "-"
