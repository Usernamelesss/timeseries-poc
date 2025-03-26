use polars::prelude::*;
use std::env;
use std::time::{SystemTime, UNIX_EPOCH};
use polars::prelude::{col, lit};

fn bench_divide_by2(df: &DataFrame) {
    let start = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
    let _divided_df_2: DataFrame = df.clone().lazy().select([
        col("param_0") / lit(2.0),
        col("param_1") / lit(2.0),
        col("param_2") / lit(2.0),
        col("param_3") / lit(2.0),
        col("param_4") / lit(2.0),
        col("param_5") / lit(2.0),
        col("param_6") / lit(2.0),
        col("param_7") / lit(2.0),
        col("param_8") / lit(2.0),
        col("param_9") / lit(2.0)
    ]).collect().unwrap();
    let end = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
    println!("Took {} ms", (end - start).as_millis());
}

fn bench_sqrt(df: &DataFrame) {
    let start = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
    let _divided_df_2: DataFrame = df.clone().lazy().select([
        col("param_0").sqrt() + (col("param_0") / lit(2.0)) + col("param_0").pow(2.0) + (col("param_0") % lit(360.0)).cos(),
        col("param_1").sqrt() + (col("param_1") / lit(2.0)) + col("param_1").pow(2.0) + (col("param_1") % lit(360.0)).cos(),
        col("param_2").sqrt() + (col("param_2") / lit(2.0)) + col("param_2").pow(2.0) + (col("param_2") % lit(360.0)).cos(),
        col("param_3").sqrt() + (col("param_3") / lit(2.0)) + col("param_3").pow(2.0) + (col("param_3") % lit(360.0)).cos(),
        col("param_4").sqrt() + (col("param_4") / lit(2.0)) + col("param_4").pow(2.0) + (col("param_4") % lit(360.0)).cos(),
        col("param_5").sqrt() + (col("param_5") / lit(2.0)) + col("param_5").pow(2.0) + (col("param_5") % lit(360.0)).cos(),
        col("param_6").sqrt() + (col("param_6") / lit(2.0)) + col("param_6").pow(2.0) + (col("param_6") % lit(360.0)).cos(),
        col("param_7").sqrt() + (col("param_7") / lit(2.0)) + col("param_7").pow(2.0) + (col("param_7") % lit(360.0)).cos(),
        col("param_8").sqrt() + (col("param_8") / lit(2.0)) + col("param_8").pow(2.0) + (col("param_8") % lit(360.0)).cos(),
        col("param_9").sqrt() + (col("param_9") / lit(2.0)) + col("param_9").pow(2.0) + (col("param_9") % lit(360.0)).cos(),
    ]).collect().unwrap();
    let end = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
    println!("Took {} ms", (end - start).as_millis());
}


fn main() {
    let project_root = env::var("PROJECT_ROOT").unwrap_or(String::new());
    let file_name = format!("{project_root}/fixtures/sample_001.parquet");
    let mut file = std::fs::File::open(file_name).unwrap();

    let df = ParquetReader::new(&mut file).finish().unwrap()
        .drop("__index_level_0__")  // remove index column since Polars doesn't use it
        .unwrap();

    println!("Loaded {} rows and {} cols", df.height(), df.width());
    // println!("{}", df.head(Some(1)));
    bench_divide_by2(&df);
    bench_sqrt(&df);
}
