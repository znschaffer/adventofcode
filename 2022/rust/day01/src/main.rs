use core::fmt;

#[derive(Debug)]
struct Answer {
    part1: u64,
    part2: u64,
}

impl fmt::Display for Answer {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Part 1: {}\nPart 2: {}", self.part1, self.part2)
    }
}

fn main() {
    let mut answer = Answer { part1: 0, part2: 0 };

    let mut groups = Vec::new();

    // part 1
    for group in include_str!("input.txt").split("\n\n") {
        let mut sum = 0;

        for line in group.lines() {
            match line.parse::<u64>() {
                Ok(n) => sum += n,
                _ => continue,
            }
        }

        if sum > answer.part1 {
            answer.part1 = sum;
        }

        groups.push(sum);
    }

    // part 2
    groups.sort();
    groups.reverse();
    answer.part2 = groups.iter().take(3).sum::<u64>();

    println!("{}", answer);
}
