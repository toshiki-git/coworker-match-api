PHONY: gen-type-file

# Variables
OPENAPI_URL=https://raw.githubusercontent.com/toshiki-git/coworker-match-openapi/main/dist/combined-from-ci.yml
GENERATOR_IMAGE=openapitools/openapi-generator-cli
OUTPUT_DIR=/local/gen/go

define GENERATE_CMD
    docker run --rm \
        -v ${CURDIR}:/local ${GENERATOR_IMAGE} generate \
        -i ${OPENAPI_URL} \
        -g go \
        -o ${OUTPUT_DIR} \
        --global-property models,supportingFiles=utils.go
endef

gen-type-file:
	$(GENERATE_CMD)
