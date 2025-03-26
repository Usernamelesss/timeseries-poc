import os
import random
import pandas as pd
from datetime import timedelta, datetime
from mockseries.trend import LinearTrend
from mockseries.seasonality import SinusoidalSeasonality
from mockseries.noise import RedNoise
from mockseries.utils import datetime_range


def get_generator(
        include_seasonality = True,
        linear_trend_coefficient = 2,
):
    trend = LinearTrend(coefficient=linear_trend_coefficient, time_unit=timedelta(days=4), flat_base=100)
    noise = RedNoise(mean=0, std=1.5, correlation=0.5)

    if include_seasonality:
        seasonality = SinusoidalSeasonality(amplitude=random.randrange(1, 10), period=timedelta(days=random.randrange(1, 100)))
        return trend + noise + seasonality

    return trend + noise


if __name__ == '__main__':
    columns_to_gen = 10
    granularity = timedelta(seconds=10)

    time_points = datetime_range(
        granularity=granularity,
        start_time=datetime(2021, 5, 31),
        end_time=datetime(2021, 8, 30),
    )

    ts = {}

    for i in range(columns_to_gen):
        ts[f"param_{i}"] = get_generator(
                 include_seasonality = i % 2 == 0,
                 linear_trend_coefficient = random.randrange(-2, 2),
             ).generate(time_points=time_points)
        print(f"Generated data for parameter {i}")

    df = pd.DataFrame(ts, index=time_points)
    df.to_parquet(path=os.path.join(os.environ["PROJECT_ROOT"], "fixtures", "sample_001.parquet"))
    print(f"Generated {len(df)} timestamps")
