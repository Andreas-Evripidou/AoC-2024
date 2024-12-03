# Advent of Code
Welcome to my attempt to solve the daily advent of code challenges using Go.
To help me get ready faster each year for the challenges, I wrote the init.go script. 
The script takes as a year parameter and then initialise a folder with subfolders for each day for the given year.

## Prerequisites
- Have Go 1.23.1 or higher installed in your machine.

## Usage

1. Clone or download this repository to your local machine.
2. Navigate to the directory where you saved the repository.
3. Run the following command to execute the script and generate the structure:
```bash
go run main.go -year 2024
```
Replace 2024 with the appropriate year if you are setting up a different year's challenge.

You must specify the year using the -year flag. This will create a root directory for the year and populate it with files and directories for the challenge.

## File Structure

After running the script, the following structure will be created:

```graphql
2024/                  # Root directory for the year
├── day01/             # Directory for Day 1 challenge
│   ├── main.go        # Solution file for Day 1
│   ├── main_test.go   # Tests for Day 1 solution
│   ├── input.txt      # Input file for Day 1
│   └── testInput.txt  # Test input file for Day 1
├── day02/             # Directory for Day 2 challenge
│   ├── main.go        # Solution file for Day 2
│   ├── main_test.go   # Tests for Day 2 solution
│   ├── input.txt      # Input file for Day 2
│   └── testInput.txt  # Test input file for Day 2
...
└── day25/             # Directory for Day 25 challenge
    ├── main.go        # Solution file for Day 25
    ├── main_test.go   # Tests for Day 25 solution
    ├── input.txt      # Input file for Day 25
    └── testInput.txt  # Test input file for Day 25

```
