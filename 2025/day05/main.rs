use std::fs;

fn main() {
    let puzzle = parse("assets/test.txt");
    let mut p1 = 0;
    for item in puzzle.items {
        for r in puzzle.inventory.iter() {
            if r.contains(item) {
                p1 += 1;
                break;
            }
        }
    }
    println!("p1 {}", p1);

    let p2: i64 = puzzle.inventory.iter().map(|r| r.size()).sum();
    println!("p2 {}", p2);
}

struct Range {
    start: i64,
    end: i64,
}

impl Range {
    fn contains(&self, v: i64) -> bool {
        self.start <= v && v <= self.end
    }

    fn size(&self) -> i64 {
        (self.end + 1) - self.start
    }

    fn overlaps(&self, other: Range) -> bool {
        self.start <= other.end || self.end >= other.start
    }

    fn grow(&mut self, other: Range) -> Range {
        Range {
            start: self.start.min(other.start),
            end: self.end.max(other.end),
        }
    }
}

struct Puzzle {
    inventory: Vec<Range>,
    items: Vec<i64>,
}

fn parse(fname: &str) -> Puzzle {
    let data = fs::read_to_string(fname).unwrap();
    let (inventory_str, items) = data.split_once("\n\n").unwrap();

    let inventory = inventory_str
        .lines()
        .map(|line| -> Range {
            let (start, end) = line.split_once("-").unwrap();
            Range {
                start: start.parse().unwrap(),
                end: end.parse().unwrap(),
            }
        })
        .collect();

    let items = items
        .lines()
        .map(|line| -> i64 { line.parse().unwrap() })
        .collect();

    Puzzle { inventory, items }
}

// fn consolidate_ranges(ranges: Vec<Range>) -> Vec<Range> {
//     let mut out: Vec<Range> = ranges.clone().collect();
//     loop {
//         let mut inner: Vec<Range> = Vec::new();
//         for r in out {
//             for i in inner {
//                 if r.overlaps(i) {}
//             }
//         }
//
//         if inner.len() != out.len() {
//             break;
//         }
//         out = inner;
//     }
//     out
// }
