.PHONY: install install-skills generate-test test version-up

install:
	bash ./shell/install.sh

install-skills:
	@mkdir -p ~/.claude/skills
	@cp -r ./skills/* ~/.claude/skills/
	@echo "Installed skills to ~/.claude/skills/"

generate-test:
	bash ./shell/generate_test_from_sample.sh $(ARGS)

test:
	bash ./shell/test.sh

version-up:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Usage: make version-up vX.X.XX"; \
		exit 1; \
	fi
	$(eval VERSION := $(filter-out $@,$(MAKECMDGOALS)))
	sed -i 's/const Version = ".*"/const Version = "$(VERSION)"/' cmd/version.go
	git add cmd/version.go
	git commit -m "minor version up"
	git tag -a $(VERSION) -m "$(VERSION)"
	@echo "Version updated to $(VERSION), committed and tagged"

%:
	@:

