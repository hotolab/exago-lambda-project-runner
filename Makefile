.PHONY: build compile

ZIP := exago-project-runner.zip
RUNNER_PATH := $(GOPATH)/src/github.com/hotolab/exago-runner/cmd/exago-runner

build: cleanup 
	@echo "Adding directories to zipfile"
	zip -r9 $(ZIP) bin/ git/

# Setup build dependencies (not related to project)
cleanup:
	@echo "Cleaning up directories"
	@find . -name ".DS_Store" -exec rm {} \;
	@rm -f $(ZIP)

compile:
	@echo "Compiling dependencies"
	cd $(RUNNER_PATH) && CGO_ENABLED=0 GOOS=linux go build -v -i -a -tags netgo
	mv "$(RUNNER_PATH)/exago-runner" /var/task/bin/linux-amd64/ && cd "-"