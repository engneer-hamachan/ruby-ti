.PHONY: install apply-rbs reset-apply generate-test test

install:
	sh ./shell/install.sh

apply-rbs:
	sh ./shell/rbs_to_json.sh $(filter-out $@,$(MAKECMDGOALS)) && ./shell/install_myjson.sh

reset-apply:
	git checkout builtin/builtin_config/
	git clean -f builtin/builtin_config/

%:
	@:

generate-test:
	sh ./shell/generate_test_from_sample.sh

test:
	sh ./shell/test.sh
