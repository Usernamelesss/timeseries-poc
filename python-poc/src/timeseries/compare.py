import os

import pandas as pd
from pandas.testing import assert_frame_equal


if __name__ == '__main__':
    results_dir = os.path.join(os.environ.get("PROJECT_ROOT", "/"), "results")

    python_res = [
        pd.read_parquet(os.path.join(results_dir, "python_divide_by2.parquet")),
        pd.read_parquet(os.path.join(results_dir, "python_sqrt.parquet")),
        pd.read_parquet(os.path.join(results_dir, "python_ema.parquet")),
    ]

    golang_res = [
        pd.read_parquet(os.path.join(results_dir, "golang_divide_by2.parquet")),
        pd.read_parquet(os.path.join(results_dir, "golang_sqrt.parquet")),
        pd.read_parquet(os.path.join(results_dir, "golang_ema.parquet")),
    ]

    rust_res = [
        pd.read_parquet(os.path.join(results_dir, "rust_divide_by2.parquet")),
        pd.read_parquet(os.path.join(results_dir, "rust_sqrt.parquet")),
        pd.read_parquet(os.path.join(results_dir, "rust_ema.parquet")),
    ]

    mapping = {
        0: "divide by 2",
        1: "sqrt",
        2: "exponential moving average",
    }

    for i, (python_df, golang_df, rust_df) in enumerate(zip(python_res, golang_res, rust_res)):
        # Fix indexes because Go and Rust doesn't use it
        golang_df.set_index("__index_level_0__", inplace=True, drop=True)
        golang_df.index.name = None
        golang_df.index = golang_df.index.tz_localize(None)  # remove UTC timezone that Go ads by default
        rust_df.set_index("__index_level_0__", inplace=True, drop=True)
        rust_df.index.name = None

        print("----------------------------------------------------")
        print(f"Comparing {mapping[i]} between Python and Golang...")
        assert_frame_equal(python_df, golang_df, check_exact=False, rtol=1e-10)
        print(f"Comparing {mapping[i]} between Golang and Rust...")
        assert_frame_equal(golang_df, rust_df, check_exact=False, rtol=1e-10)
        print("----------------------------------------------------")
