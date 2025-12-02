use fancy_regex::Regex;
use std::error::Error;
use std::fs::read_to_string;

fn main() {
    let ranges = read_ranges("assets/input.txt").unwrap();
    let value = ranges.iter().fold(0i64, |mut acc, range| {
        for i in range.start..range.end {
            let num_digits = (i as f64).log10().ceil() as i64;
            if num_digits % 2 != 0 {
                continue;
            }
            // we have an even number of digits
            let divisor = 10i64.pow(num_digits as u32 / 2);
            let l = i / divisor;
            let r = i % divisor;
            if l == r {
                // left and right half match!
                acc += i
            }
        }
        acc
    });
    println!("{}", value);
    let re = Regex::new(r"^(\d+)\1+$").unwrap();

    let p2_value = ranges.iter().fold(0i64, |mut acc, range| {
        for i in range.start..range.end {
            // look at all the values in the ranges
            let num_str = i.to_string();
            if re.is_match(&num_str).unwrap() {
                acc += i;
            }
        }
        acc
    });
    println!("{}", p2_value);
}

fn read_ranges(fpath: &str) -> Result<Vec<Range>, Box<dyn Error>> {
    let data = read_to_string(fpath)?;
    let out = data
        .split(",")
        .map(|l| -> Range {
            let (s, e) = l.trim_end().trim_start().split_once("-").unwrap();
            let start: i64 = s.parse().unwrap();
            let end: i64 = e.parse().unwrap();
            Range {
                start,
                end: end + 1,
            }
        })
        .collect();

    Ok(out)
}

#[derive(Debug)]
struct Range {
    start: i64,
    end: i64,
}
