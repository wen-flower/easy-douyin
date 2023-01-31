.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build build-api build-user build-video
build: tidy build-api build-user build-video
	@echo "build: api user video"
build-api:
	$(MAKE) -C ./cmd/api build
build-user:
	$(MAKE) -C ./cmd/user build
build-video:
	$(MAKE) -C ./cmd/video build

.PHONY: format format-api format-user format-video
format: tidy format-api format-user format-video
format-api:
	$(MAKE) -C ./cmd/api format
format-user:
	$(MAKE) -C ./cmd/user format
format-video:
	$(MAKE) -C ./cmd/video format

.PHONY: clean clean-api clean-user clean-video
clean: clean-api clean-user clean-video
clean-api:
	$(MAKE) -C ./cmd/api clean
clean-user:
	$(MAKE) -C ./cmd/user clean
clean-video:
	$(MAKE) -C ./cmd/video clean
