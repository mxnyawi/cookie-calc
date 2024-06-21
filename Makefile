.PHONY: clone build test run clean

# Repository URL
REPO=https://github.com/mxnyawi/cookie-calc.git

# Default CSV file and date for running the application
CSV_FILE=./cookie_log.csv
DATE=2018-12-08

# Clone the repository
clone:
	git clone $(REPO)

# Build the project
build:
	cd cmd/cookie-calc && go build -o cookie-calc

# Run tests
test:
	go test ./...

# Run the application with default parameters
run:
	cd cmd/cookie-calc && ./cookie-calc -f $(CSV_FILE) -d $(DATE) -log

# Clean up build artifacts
clean:
	cd cmd/cookie-calc && rm -f cookie-calc
	cd cmd/cookie-calc && rm -f cookie_calc.log