NIGHTLY_TAG := v0.0.0

nightly:
	git tag -d $(NIGHTLY_TAG) &> /dev/null
	git push origin :$(NIGHTLY_TAG) &> /dev/null

	git tag $(NIGHTLY_TAG) &> /dev/null
	git push origin $(NIGHTLY_TAG) &> /dev/null
	git fetch --tags