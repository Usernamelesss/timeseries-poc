import math
import os
import time

import numpy as np
import pandas as pd


def divide_by_2(df: pd.DataFrame) -> pd.DataFrame:
    """
    Divide all DataFrame elements by 2
    """
    return df / 2  # vectorized
    # return df.map(lambda x: x / 2)


def sqrt_formula(df: pd.DataFrame) -> pd.DataFrame:
    return df.map(lambda x: math.sqrt(x) + x/2 + math.pow(x, 2) + math.cos(x%360))


def vectorized_sqrt_formula(df: pd.DataFrame) -> pd.DataFrame:
    return np.sqrt(df) + df/2 + np.power(df, 2) + np.cos(df % 360)


def exponential_moving_average(df: pd.DataFrame) -> pd.DataFrame:
    return df.ewm(span=10, adjust=False, min_periods=10).mean()


def bench(description, fn, *args, **kwargs) -> (pd.DataFrame, float):
    start = time.monotonic_ns()
    r = fn(*args, **kwargs)
    end = time.monotonic_ns()
    elapsed = (end - start) / 1e+6
    print("{} took {:0.2f} ms".format(description, elapsed))
    return r, elapsed


if __name__ == '__main__':
    timeseries = pd.read_parquet(os.path.join(os.environ.get("PROJECT_ROOT", "/"), "fixtures", "sample_001.parquet"))
    output_dir = os.path.join(os.environ.get("PROJECT_ROOT", "/"), "results")

    os.makedirs(output_dir, exist_ok=True)

    r1, elapsed1 = bench("Divide By 2", divide_by_2, timeseries)
    r1.to_parquet(os.path.join(output_dir, "python_divide_by2.parquet"))

    r3, elapsed2 = bench("Sqrt", vectorized_sqrt_formula, timeseries)
    r3.to_parquet(os.path.join(output_dir, "python_sqrt.parquet"))

    r4, elapsed3 = bench("EWMA", exponential_moving_average, timeseries)
    r4.to_parquet(os.path.join(output_dir, "python_ema.parquet"))

    timing_path = os.path.join(output_dir, "python_timing.csv")
    timing_df = pd.DataFrame({
            "divide_by_2[ms]": [elapsed1],
            "sqrt[ms]": [elapsed2],
            "ewma[ms]": [elapsed3],
        })
    if os.path.exists(timing_path):
        timing_df.to_csv(timing_path, mode="a", index=False, header=False)
    else:
        timing_df.to_csv(os.path.join(output_dir, "python_timing.csv"), mode="x", index=False)
