TAG_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell git log -1 --format=%cd --date=format:"%Y%m%d")
VERSION := $(TAG:v%=%)

ifneq ($(COMMIT), $(TAG_COMMIT))
	VERSION := $(VERSION)-next-$(COMMIT)-$(DATE)
endif

ifeq ($(VERSION),)
	VERSION := $(COMMIT)-$(DATA)
endif

version:
	@echo $(VERSION)

nightly:
	git tag -d nightly &> /dev/null
	git tag nightly
	git push origin :nightly &> /dev/null
	git push origin nightly
	git fetch --tags