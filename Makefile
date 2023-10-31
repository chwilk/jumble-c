# COLORS
RED    := $(shell tput -Txterm setaf 1)
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
VIOLET := $(shell tput -Txterm setaf 5)
AQUA   := $(shell tput -Txterm setaf 6)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


## Show help
help:
	@echo ''
	@echo 'Makefile to build and run jumble-c'
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET}\t${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
# magic help script borrowed from https://github.com/kubevirt/kubevirt.github.io/blob/main/Makefile

## Build image
build: 
	@echo "${GREEN}Makefile: Building Image ${RESET}"
	podman build -t jumble .
	@echo

## Run site.  App available @ http://localhost:8080
run: | stop
	@echo "${GREEN}Makefile: Run site${RESET}"
	podman run -d --rm --publish 8080:8080 --name jumble jumble
	@echo

## Stop running container
stop:
	@echo "${GREEN}Makefile: Stop running container${RESET}"
	podman rm -f jumble 2> /dev/null; echo
	@echo -n