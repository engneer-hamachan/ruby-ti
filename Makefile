.PHONY: install apply-rbs generate-test

install:
	sh ./shell/install.sh

apply-rbs:
	sh ./shell/rbs_to_json.sh $(filter-out $@,$(MAKECMDGOALS)) && ./shell/install_myjson.sh

%:
	@:

generate-test:
	sh ./shell/generate_test_from_sample.sh
