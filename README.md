# Guess-it-1

## Overview

The **Guess-it-1** project is a command-line tool written in Go designed to predic a range for the next input number base on the previous inputs. T in which the next number will likely fall, given a sequence of numbers as input. **statistical and probabilistic** calculations based on the **mean** and **standard deviataion** of the input data. The goal is to generate the most accurate predictions with the smallest possible range.

The program takes a number as input (representing a graph's y-axis value) and calculates the probable range for the next input number based on the previous values.

## Key Features

+ **Input via Standard Input (Stdin)**: Numbers are fed to the program through standard input.
+ **Predictive Algorithm**: Uses mean and standard deviation to calculate the next number's probable range.
+ **Range Display**: For each input, the program outputs a lower and upper bound for the next number.
+ **Performance Optimization**: The program is designed to minimize range size while maximizing accuracy.

### Example Usage

Below is an example of how the program interacts with user input and predicts the range for the next number:
```bash
$ ./your_program
189 --> the standard input
120 200 --> the range for the next input (113)
113 --> the standard input
160 230 --> the range for the next input (121)
121 --> the standard input
110 140 --> the range for the next input (114)
```
As seen in the example, the program takes each number and produces a predicted range for the subsequent number.

## Project Structure
+ **Main Program:** Located in the `main.go` file. It calls the `Reader()` function to handle input and outputs the predicted range.
+ **Utilities:** The `utils` package handles reading inputs and managing the flow of the program.
+ **Statistics:** The `stats` package calculates the mean and standard deviation, which are used to predict the range of the next number.

### Project Directory Structure
```
guess-it-1/
│
├── student/
│   │
│   ├── main.go               # Entry point of the program 
│   │
│   ├── stats/
│   │   ├── guesser_test.go
│   │   └── guesser.go              # Contains statistical functions (Mean, Standard Deviation, and Guesser)
│   │
│   ├── utils/
│   │   └── reader.go             # Handles input reading and output formatting
│   │
│   ├── test_data/
│   │   └── data.txt              # Sample data used for testing
│   │
│   ├── script.sh             # Shell script to run the program for testing purposes
│   │
│   └── go.sum                    # Dependencies file
│
├── README.md                 # Project documentation
└── LICENSE                   # License file (if applicable)
```

## How It Works

The prediction of the next number is based on the following statistical formula:

1. **Mean Calculation:** The program calculates the mean of the input data. The mean represents the central tendency of the input numbers.

2. **Standard Deviation Calculation:** The program calculates the standard deviation, which measures the variability or spread of the input numbers.

3. **Range Prediction:** Using the mean and standard deviation, the program predicts the next number's range:
    + Lower bound = Mean - (3.1 × Standard Deviation)
    + Upper bound = Mean + (3.1 × Standard Deviation)

    The `3.1` multiplier helps to adjust the prediction window for a balance between accuracy and range size.

## Installation & Usage

### Prerequisites

 + **Go:** Ensure Go is installed on your system. You can download it [here](https://go.dev/doc/).

### Steps to Run the Program

1. Clone the repository:
```bash
git clone https://learn.zone01kisumu.ke/git/oumouma/guess-it-1
cd guess-it-1
```

2. Build and run the program:
```bash
cd cmd/
go run main.go
```

3. Provide input via the command line:
```
189
113
121
114
```

4. View the predicted ranges as output.

### Running Tests

To run the tests, ensure that your test data is in the `test/test_data/` folder. You can run the tests with:
```bash
go test ./...
```

## Shell Script

To facilitate automated testing, the script.sh file has been provided. It runs the program from the root directory:
```bash
#!/bin/bash

# navigate into student directory
cd student/

# run the program
go run main.go
```

## How to Contribute

If you encounter any issues or have suggestions for improvements, feel free to submit an issue or propose changes through pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE]() file for more details.

Author

This project is created and maintained by [Ouma Ouma](https://github.com/garveyshah).