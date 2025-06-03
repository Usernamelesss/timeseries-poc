import os

import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt


if __name__ == '__main__':
    results_dir = os.path.join(os.environ.get("PROJECT_ROOT", "/"), "results")

    python = pd.read_csv(os.path.join(results_dir, "python_timing.csv"))
    rust = pd.read_csv(os.path.join(results_dir, "rust_timing.csv"))
    go = pd.read_csv(os.path.join(results_dir, "golang_timing.csv"))

    python["id"] = "python"
    rust["id"] = "rust"
    go["id"] = "go"

    df_all = pd.concat([python, go, rust], ignore_index=True)

    df_melted = df_all.melt(id_vars='id',
                            var_name='benchmark',
                            value_name='time_ms')

    # Plot: use boxplot to compare distributions
    plt.figure(figsize=(10, 6))
    sns.boxplot(data=df_melted, x='benchmark', y='time_ms', hue='id')
    plt.title('Benchmark Comparison Across Runs')
    plt.ylabel('Time (ms)')
    plt.xlabel('Benchmark')
    plt.tight_layout()
    plt.show()