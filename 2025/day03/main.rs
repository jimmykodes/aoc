use std::{error::Error, fs};

fn main() -> Result<(), Box<dyn Error>> {
    let banks = parse_banks("assets/input.txt")?;

    let p1 = solve(banks.clone(), 2);
    println!("{}", p1);

    let p2 = solve(banks.clone(), 12);
    println!("{}", p2);

    Ok(())
}

fn solve(banks: Vec<Bank>, size: usize) -> i64 {
    let mut out: i64 = 0;

    for bank in banks.iter() {
        let mut left_idx = 0;
        let mut bank_val: i64 = 0;
        for i in 0..size {
            let right_idx_offset = size - (i + 1);
            let right_idx = bank.len() - right_idx_offset;
            let (idx, val) = find_max(&bank[left_idx..right_idx]);
            left_idx += idx + 1;
            bank_val += val as i64 * (10i64.pow(right_idx_offset as u32))
        }
        out += bank_val;
    }

    out
}

type Bank = Vec<i32>;

fn parse_banks(fname: &str) -> Result<Vec<Bank>, Box<dyn Error>> {
    let mut out = Vec::new();
    let data = fs::read_to_string(fname)?;

    for line in data.lines() {
        let bank = line
            .chars()
            .map(|c| c.to_digit(10).unwrap() as i32)
            .collect();
        out.push(bank)
    }

    Ok(out)
}

fn find_max(v: &[i32]) -> (usize, i32) {
    let mut max_val = 0;
    let mut max_idx = 0;
    for (idx, val) in v.iter().enumerate() {
        if *val > max_val {
            max_val = *val;
            max_idx = idx
        }
    }
    return (max_idx, max_val);
}
