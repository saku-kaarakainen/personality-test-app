.PHONY: test
test:
	$(MAKE) -C api test && $(MAKE) -C app test

.PHONY: run
run:
	$(MAKE) -C api run & $(MAKE) -C app run
