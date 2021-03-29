NIGHTLY_TAG := v0.0.0

next:
	git tag -fa $(NIGHTLY_TAG) -m "Next release" && git push origin $(NIGHTLY_TAG) --force
