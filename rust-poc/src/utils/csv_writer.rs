use std::env;
use std::error::Error;
use std::fs::OpenOptions;
use std::path::Path;

pub(crate) fn write_timing(elapsed_list: [u128; 3]) -> Result<(), Box<dyn Error>> {
    let project_root = env::var("PROJECT_ROOT").unwrap_or(String::new());
    let csv_path = format!("{project_root}/results/rust_timing.csv");

    let needs_headers = Path::new(&csv_path).exists() == false;

    let csv_file = OpenOptions::new()
        .write(true)
        .create(true)
        .append(true)
        .open(csv_path.clone())?;
    let mut wtr = csv::Writer::from_writer(csv_file);

    if needs_headers {
        wtr.write_record(["divide_by_2[ms]", "sqrt[ms]", "ewma[ms]"])?;
    }

    // Convert u128 numbers to strings before writing
    let string_fields: Vec<String> = elapsed_list.iter().map(|value| value.to_string()).collect();

    wtr.write_record(&string_fields)?;
    wtr.flush()?;

    Ok(())
}
