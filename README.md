# csv_summary_benchmark
benchmarking CSV summary generation performance between R, Python, and Go.

Submitted assignment for Northwestern MSDS 431

Within the csv_summary directory I've split the functionality into separate files. 

- `csv_reader`: converts the list of comma separated strings into floats, skipping any non-numeric fields. 
- `summary_stats`: draws from the go-stats package and deploys the [Describe() function](https://github.com/montanaflynn/stats/blob/v0.7.1/describe.go#L23) that summarizes every numeric column, mirroring similar functionality in R and Python. 
- `errors`: contains common errors found during csv parsing, used to isolate where specific portions of the program could break
- `csv_test`: tests the two primary functions in `csv_reader` and `summary_stats` respectively. 
- `main`: opens the given csv file, calculates the summary and writes to a `.txt` file 100 times

Separated from this repository are similar programs in R and in Python. Here's what I did to operate the benchmark test:

1. `Create a virtual environment in the main directory (.venv is in the gitignore)`
2. `Create a directory 'given_files' that contains the R and Python scripts to execute`
3. run the following commands in the main directory three times each, storing the results for 'real time':
    - `time ./csv-summary.exe`
    - `time py ./given_files/py_file.py`
    - `time Rscripts ./given_files/r_file.R`

