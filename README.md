# Cookie Frequency Calculator

This program calculates the most frequently occurring cookies in a given CSV file for a specific date.

## Installation

First, clone the repository to your local machine:

```bash
git clone https://github.com/mxnyawi/cookie-calc.git
```

Then, navigate to the cmd/cookie-calc directory:
```bash
cd cookie-calc/cmd/cookie-calc
```

## Usage

To run the program, use the go run command followed by a -f flag to specify the CSV file, a -d flag to specify the date, and optionally a -log flag to enable logging. For example:

```bash
./cookie-calc -f .\cookie_log.csv -d 2018-12-08 -log
```

This command will calculate the most frequently occurring cookies in the cookie_log.csv file for the date 2018-12-08 and enable logging. When logging is enabled with the `-log` flag, a log file named `cookie_calc.log` will be created in the `cmd/cookie-calc` directory.

You can also use the `-h` flag to display help information about the command usage:

```bash
./cookie-calc -h
```

## CSV File Format

The CSV file should have the following format:

```csv
cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
...
```

Each line represents a cookie with a unique ID and the timestamp when the cookie was accessed

## Output

The program outputs the cookie(s) that were accessed the most on the specified date. If multiple cookies have the same maximum number of accesses, all of them are output, one per line. For example:

```bash
AtY0laUfhglK3lC7
SAZuXPGUrfbcn5UA
```

If the `-log` flag is used, detailed execution logs will be written to `cookie_calc.log` in the `cmd/cookie-calc` directory.

## Testing

To run the tests, run:

```bash
cd /pkg/calculator
go test

cd api/csvreader
go test
```