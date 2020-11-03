.ONESHELL:

.PHONY: default
default: test

.PHONY: test
test:
	function tearDown {
		docker kill $${TEST_BENCH}
	}
	trap tearDown EXIT
	TEST_BENCH=$$(docker run --rm -d -p 8080:80 -e GALAXY_CONFIG_OVERRIDE_ALLOW_USER_DELETION=true quay.io/bgruening/galaxy:19.09)
	until curl -sS --fail -o /dev/null "http://localhost:8080/api/version"; do sleep 1; done
	GALAXY_HOST=http://localhost:8080 GALAXY_API_KEY=admin go test ./... || exit 1
