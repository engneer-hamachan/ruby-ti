.PHONY: install generate-test test

install:
	bash ./shell/install.sh

generate-test:
	bash ./shell/generate_test_from_sample.sh

test:
	bash ./shell/test.sh

%:
	@:

