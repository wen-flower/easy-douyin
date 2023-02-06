.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build build-api build-user build-video build-chat
build: tidy build-api build-user build-video build-chat
	@echo "build: api user video"
build-api:
	$(MAKE) -C ./cmd/api build
build-user:
	$(MAKE) -C ./cmd/user build
build-video:
	$(MAKE) -C ./cmd/video build
build-chat:
	$(MAKE) -C ./cmd/chat build

.PHONY: format format-api format-user format-video format-chat
format: tidy format-api format-user format-video format-chat
format-api:
	$(MAKE) -C ./cmd/api format
format-user:
	$(MAKE) -C ./cmd/user format
format-video:
	$(MAKE) -C ./cmd/video format
format-chat:
	$(MAKE) -C ./cmd/chat format

.PHONY: clean clean-api clean-user clean-video clean-chat
clean: clean-api clean-user clean-video clean-chat
clean-api:
	$(MAKE) -C ./cmd/api clean
clean-user:
	$(MAKE) -C ./cmd/user clean
clean-video:
	$(MAKE) -C ./cmd/video clean
clean-chat:
	$(MAKE) -C ./cmd/chat clean
