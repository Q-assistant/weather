NIGHTLY_TAG := v0.0.0
TAG_EXISTS := $(shell git rev-parse "$(NIGHTLY_TAG)" >/dev/null)

nightly:
	@echo "Delete current nightly tag"
	git tag -d $(NIGHTLY_TAG) &> /dev/null
	git push --force origin :$(NIGHTLY_TAG) &> /dev/null
	@echo "Re-create nightly tag"
	git tag $(NIGHTLY_TAG)
	git push --force origin $(NIGHTLY_TAG)
	git fetch --tags