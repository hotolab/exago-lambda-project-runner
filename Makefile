.PHONY: build deploy

ZIP := exago-project-runner.zip

build: cleanup | binary 
	@echo "Adding directories to zipfile"
	zip -9 -r $(ZIP) bin/ git/

# Setup build dependencies (not related to project)
cleanup:
	@echo "Cleaning up directories"
	@find . -name ".DS_Store" -exec rm {} \;
	@rm -f $(ZIP)

binary:
	@echo "Creating AWS lambda binary"
	docker run --rm -v $(GOPATH):/go -v $(PWD):/tmp eawsy/aws-lambda-go -function index -handler handler -package $(ZIP)