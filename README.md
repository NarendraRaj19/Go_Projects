# This repo contains a list of applications built using Go

# A WC-Tool - Coding Challenge #1

This is my own version of the Unix command line tool `wc`. It is a solution to the first problem in [John Crickett's Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc/).

## How to build

Execute `go build` to create the executable. It will create the executable called `ccwc`. 

## How to run

Execute the binary with a file as an argument to get the result.

The following are the possible arguments:

```bash
# Outputs the number of bytes

> ./ccwc -c WC_toolText.text
341836 test.txt

# Outputs the number of line breaks

> ./ccwc -l WC_toolText.text
7137 test.txt

# Outputs the number of words

> ./ccwc -w WC_toolText.text
58159 test.txt

# Outputs with -c, -l and -w flags
> ./ccwc WC_toolText.text
341836 7137 58159 WC_toolText.txt
```

You can also read from standard input

```bash
> cat WC_toolText.txt | ./ccwc -l
7137
```

## How to run the tests

Execute `go test`

You should see a result matching something like:

```bash
PASS
ok      _/C_/Users/naren/Downloads/Go_Projects/ccwc     0.788
```

## License

[MIT](LICENSE) © Naren