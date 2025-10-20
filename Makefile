.PHONY: install apply-rbs reset-apply generate-test generate-rbs-test test rbs-test

install:
	bash ./shell/install.sh

apply-rbs:
	bash ./shell/rbs_to_json.sh $(filter-out $@,$(MAKECMDGOALS)) && bash ./shell/install_myjson.sh

reset-apply:
	git checkout builtin/builtin_config/
	git clean -f builtin/builtin_config/

generate-test:
	bash ./shell/generate_test_from_sample.sh

generate-rbs-test:
	bash ./shell/generate_test_from_rbs.sh $(filter-out $@,$(MAKECMDGOALS)) 

test:
	bash ./shell/test.sh

rbs-test:
	go test ./rbs_test/... -count=1

%:
	@:

