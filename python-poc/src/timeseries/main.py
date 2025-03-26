import math
import os
import time

import numpy as np
import pandas as pd

from pandas.testing import assert_frame_equal


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


def bench(fn, *args, **kwargs) -> pd.DataFrame:
    start = time.monotonic_ns()
    r = fn(*args, **kwargs)
    end = time.monotonic_ns()
    elapsed = (end - start) / 1e+6
    print("Took {:0.2f} ms".format(elapsed))
    return r


if __name__ == '__main__':
    timeseries = pd.read_parquet(os.path.join(os.environ.get("PROJECT_ROOT", "/"), "fixtures", "sample_001.parquet"))

    r1 = bench(divide_by_2, timeseries)
    r2 = bench(sqrt_formula, timeseries)
    r3 = bench(vectorized_sqrt_formula, timeseries)

    assert_frame_equal(r2, r3, check_exact=False, rtol=1e-5)
