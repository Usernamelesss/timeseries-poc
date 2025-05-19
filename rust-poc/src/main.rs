use crate::utils::bench;
use polars::prelude::*;
use polars::prelude::{col, lit};
use std::env;
use std::fs::File;
use std::io::BufWriter;

pub mod utils;

fn divide_by2(df: &DataFrame) -> DataFrame {
    let lazy = df.clone().lazy();
    // Get schema to extract all column names
    let schema = df.schema();

    // Create expressions: col(name) / 2.0
    let exprs: Vec<Expr> = schema
        .iter_fields()
        .map(|field| {
            let name = field.name().as_str();
            if field.dtype != DataType::Float64 {
                return col(name);
            }
            (col(name) / lit(2.0)).alias(name)
        })
        .collect();

    // Select transformed expressions
    return lazy.select(exprs).collect().unwrap();

}

fn sqrt(df: &DataFrame) -> DataFrame {
    let lf = df.clone().lazy();
    // Get the schema so we know which columns to operate on
    let schema = df.schema();

    // Build expressions for each column
    let exprs: Vec<Expr> = schema
        .iter_fields()
        .map(|field| {
            let name = field.name().as_str();
            let col = col(name);

            if field.dtype != DataType::Float64 {
                return col; // polars::prelude::col(name);
            }

            // Apply the full expression: sqrt(col) + col / 2 + col^2 + cos(col % 360)
            (col.clone().sqrt()
                + col.clone() / lit(2.0)
                + col.clone().pow(lit(2.0))
                + (col % lit(360.0)).cos())
                .alias(name)
        })
        .collect();

    // Project all transformed columns
    return lf.select(exprs).collect().unwrap();
}

fn main() {
    let project_root = env::var("PROJECT_ROOT").unwrap_or(String::new());
    let file_name = format!("{project_root}/fixtures/sample_001.parquet");
    let mut file = File::open(file_name).unwrap();

    let df = ParquetReader::new(&mut file)
        .finish()
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
