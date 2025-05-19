use crate::utils::bench;
use polars::prelude::*;
use polars::prelude::{col, lit};
use std::env;
use std::fs::File;
use std::io::BufWriter;

pub mod utils;

fn divide_by2(df: &DataFrame) -> DataFrame {
    return df
        .clone()
        .lazy()
        .select([
            col("param_0") / lit(2.0),
            col("param_1") / lit(2.0),
            col("param_2") / lit(2.0),
            col("param_3") / lit(2.0),
            col("param_4") / lit(2.0),
            col("param_5") / lit(2.0),
            col("param_6") / lit(2.0),
            col("param_7") / lit(2.0),
            col("param_8") / lit(2.0),
            col("param_9") / lit(2.0),
        ])
        .collect()
        .unwrap();
}

fn sqrt(df: &DataFrame) -> DataFrame {
    return df
        .clone()
        .lazy()
        .select([
            col("param_0").sqrt()
                + (col("param_0") / lit(2.0))
                + col("param_0").pow(2.0)
                + (col("param_0") % lit(360.0)).cos(),
            col("param_1").sqrt()
                + (col("param_1") / lit(2.0))
                + col("param_1").pow(2.0)
                + (col("param_1") % lit(360.0)).cos(),
            col("param_2").sqrt()
                + (col("param_2") / lit(2.0))
                + col("param_2").pow(2.0)
                + (col("param_2") % lit(360.0)).cos(),
            col("param_3").sqrt()
                + (col("param_3") / lit(2.0))
                + col("param_3").pow(2.0)
                + (col("param_3") % lit(360.0)).cos(),
            col("param_4").sqrt()
                + (col("param_4") / lit(2.0))
                + col("param_4").pow(2.0)
                + (col("param_4") % lit(360.0)).cos(),
            col("param_5").sqrt()
                + (col("param_5") / lit(2.0))
                + col("param_5").pow(2.0)
                + (col("param_5") % lit(360.0)).cos(),
            col("param_6").sqrt()
                + (col("param_6") / lit(2.0))
                + col("param_6").pow(2.0)
                + (col("param_6") % lit(360.0)).cos(),
            col("param_7").sqrt()
                + (col("param_7") / lit(2.0))
                + col("param_7").pow(2.0)
                + (col("param_7") % lit(360.0)).cos(),
            col("param_8").sqrt()
                + (col("param_8") / lit(2.0))
                + col("param_8").pow(2.0)
                + (col("param_8") % lit(360.0)).cos(),
            col("param_9").sqrt()
                + (col("param_9") / lit(2.0))
                + col("param_9").pow(2.0)
                + (col("param_9") % lit(360.0)).cos(),
        ])
        .collect()
        .unwrap();
}

fn main() {
    let project_root = env::var("PROJECT_ROOT").unwrap_or(String::new());
    let file_name = format!("{project_root}/fixtures/sample_001.parquet");
    let mut file = File::open(file_name).unwrap();

    let df = ParquetReader::new(&mut file)
        .finish()
        .unwrap()
        .drop("__index_level_0__") // remove index column since Polars doesn't use it
        .unwrap();

    println!("Loaded {} rows and {} cols", df.height(), df.width());

    let mut r1 = bench::simple_bench(String::from("Divide By 2"), divide_by2, &df);
    let mut r2 = bench::simple_bench(String::from("Simple Sqrt"), sqrt, &df);
    // println!("{}", df.head(Some(1)));

    // Write results
    let r1_writer = &mut BufWriter::new(
        File::create(format!("{project_root}/results/rust_divide_by2.parquet")).unwrap(),
    );
    ParquetWriter::new(r1_writer)
        .finish(&mut r1)
        .expect("Cannot write result 1");
    let r2_writer = &mut BufWriter::new(
        File::create(format!("{project_root}/results/rust_sqrt.parquet")).unwrap(),
    );
    ParquetWriter::new(r2_writer)
        .finish(&mut r2)
        .expect("Cannot write result 2");
}
