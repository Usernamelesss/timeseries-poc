# Timeseries-PoC

Simple repository to test very simple operations on a timeseries with different programming languages.

## Setup the environment
1) Install all the required tools with mise
2) Generate a timeseries with `python python-poc/src/timeseries/ts_generator.py`, it should create a parquet file under `fixtures/` directory

## How to run the comparison
1) Build the containers with `docker compose build`
2) Run the script that launch the containers `bash run.sh`

## Utilities
The `python-poc` folder contains also some scripts for:
- Each container can be configured to write its results into a `results/` directory, `compare.py` uses the generated parquet files to verify that all the outputs are the same.
- Each container outputs all execution times into a csv file into `results/` directory, the `visualize_timings.py` allows to print a boxplot to compare all the executions

## Implemented test cases
- `n/2`: perform a `n/2` operation on all elements of the timeseries
- `sqrt(n)+n/2+n^2+cos(n%360)`: perform a more complex element-wise operation on all elements of the timeseries.
- `exponential moving average (EMA)` is a kind of moving average (or rolling average) that smooth out short-term fluctuation. Mathematically speaking it's defined as:
    ```
    EMA = (value(t) * k) + (EMA(t-1) * (1 - k))
    Where:
    - value(t) is the value at time t
    - EMA(t-1) is the value of the EMA lagged by 1 (in other terms is the previous value of the EMA)
    - k is the smothing costant, usually 2/(n+1)
    - n is the number of observation of the EMA, that means how many data points we will lookback to compute the EMA
    ```
    Few things to note here:
    1) Since we have an EMA(t-1) there is a time dependency between values, that is an ordinary thing when working with timeseries.
    2) It differs from a simple moving average (SMA) because not all the data points in the rolling window are weighted equally: 
    due to the `k` multiplier which gives more "significance" to recent data points.
