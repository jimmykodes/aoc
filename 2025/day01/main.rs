use std::error::Error;
use std::fs;
use std::io;
use std::io::BufRead;

fn main() -> Result<(), Box<dyn Error>> {
    let inst = read_instructions("assets/input.txt")?;
    let solution = inst.iter().fold(
        Counter {
            passed: 0,
            landed: 0,
            val: 50,
        },
        |mut c, v| {
            let mut val = c.val;
            val += v.delta;

            let mut passed = 0;

            c.passed += v.full_rotations;

            if val < 0 {
                val = val + 100;
                if c.val != 0 {
                    passed += 1;
                }
            } else if val > 99 {
                val = val - 100;
                passed += 1;
            }

            if val == 0 {
                if passed == 0 {
                    c.passed += 1;
                }
                c.landed += 1;
            }

            c.val = val;
            c.passed += passed;
            println!("{:?}", v);
            println!("{:?}", c);
            c
        },
    );

    println!("puzzle 1: {}", solution.landed);
    println!("puzzle 2: {}", solution.passed);
    Ok(())
}

#[derive(Debug)]
struct Counter {
    landed: i32,
    passed: i32,
    val: i32,
}

#[derive(Debug)]
struct Inst {
    /// delta represents how much the count should change after any full rotations are completed
    delta: i32,

    /// rotations represents the number of full rotations in the instruction
    full_rotations: i32,
}

fn read_instructions(fname: &str) -> Result<Vec<Inst>, Box<dyn Error>> {
    let file = fs::File::open(fname)?;
    let reader = io::BufReader::new(file);

    let mut out: Vec<Inst> = Vec::new();

    for line in reader.lines() {
        let line = line?;
        let dec = &line[0..1] == "L";
        let val: i32 = line[1..].parse()?;
        let full_rotations = val / 100;
        let mut delta = val % 100;
        if dec {
            delta *= -1
        }
        let inst: Inst = Inst {
            delta,
            full_rotations,
        };
        out.push(inst)
    }

    Ok(out)
}
