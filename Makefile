BUILD_DIR := build
SRC_DIR := src/

EXECS := $(shell find $(SRC_DIR) -mindepth 1 -maxdepth 1 -type d -exec basename {} \;)

all: $(addprefix $(BUILD_DIR)/, $(EXECS))

$(BUILD_DIR)/:
	mkdir -p $(BUILD_DIR)

$(BUILD_DIR)/%: $(SRC_DIR)%/*.go $(BUILD_DIR)
	go build -o $@ ./$<

clean:
	rm -Rf build

.PHONY: all clean
