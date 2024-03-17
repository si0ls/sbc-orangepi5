TALOS_SRC ?= https://github.com/siderolabs/talos.git
TALOS_VERSION ?= v1.7.0-alpha.1
PKG_KERNEL ?= $(REGISTRY_AND_USERNAME)/kernel:$(TAG)
KERNEL_VERSION ?= 6.8.0-talos

SHA ?= $(shell git describe --match=none --always --abbrev=8 --dirty)
TAG ?= $(shell git describe --tag --always --dirty --match v[0-9]\*)
ABBREV_TAG ?= $(shell git describe --tag --always --match v[0-9]\* --abbrev=0 )
TAG_SUFFIX ?=
IMAGE_TAG ?= $(TAG)$(TAG_SUFFIX)
USERNAME ?= si0ls
REGISTRY ?= ghcr.io
REGISTRY_AND_USERNAME ?= $(REGISTRY)/$(USERNAME)
PUSH ?= false

GIT := git
PATCH := patch

TALOS_DIR := talos
PATCHES_DIR := patches
PATCHES := $(shell find $(PATCHES_DIR) -type f -name "*.patch" -o -name ".*.patch")

PLATFORM ?= linux/amd64, linux/arm64
INSTALLER_ARCH ?= targetarch

.PHONY: all
all: talos-imager

talos: $(PATCHES)
	@$(GIT) clone --depth 1 --single-branch --branch $(TALOS_VERSION) $(TALOS_SRC) $(TALOS_DIR)
	@for patch in $^; do \
		$(PATCH) -d $(TALOS_DIR) -p1 < $$patch; \
	done

.PHONY: imager
imager: talos-imager

.PHONY: talos-%
talos-%: talos
	@$(MAKE) -C $(TALOS_DIR) $* \
		SHA=$(SHA) \
		TAG=$(TAG) \
		ABBREV_TAG=$(ABBREV_TAG) \
		TAG_SUFFIX=$(TAG_SUFFIX) \
		IMAGE_TAG=$(IMAGE_TAG) \
		USERNAME=$(USERNAME) \
		REGISTRY=$(REGISTRY) \
		REGISTRY_AND_USERNAME=$(REGISTRY_AND_USERNAME) \
		PUSH=$(PUSH) \
		PLATFORM=$(PLATFORM) \
		INSTALLER_ARCH=$(INSTALLER_ARCH) \
		PKG_KERNEL=$(PKG_KERNEL) \
		KERNEL_VERSION=$(KERNEL_VERSION) \
		$(MAKEFLAGS)

.PHONY: push
push: @$(MAKE) -C $(TALOS_DIR) push

.PHONY: clean
clean:
	@rm -rf $(TALOS_DIR)
