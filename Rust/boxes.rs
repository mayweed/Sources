use std::io;

macro_rules! print_err {
    ($($arg:tt)*) => (
        {
            use std::io::Write;
            writeln!(&mut ::std::io::stderr(), $($arg)*).ok();
        }
    )
}

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

#[derive(Debug)]
struct boxes{
    weight:f64,
    volume:f64,
    }
    
fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let box_count = parse_input!(input_line, i32);
    let mut boxes_queue:Vec<boxes>=Vec::new();
    
    for i in 0..box_count as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let weight = parse_input!(inputs[0], f64);
        let volume = parse_input!(inputs[1], f64);
        boxes_queue.push(boxes{weight,volume});
    }

    // Write an action using println!("message...");
    print_err!("{:?}",boxes_queue);

    println!("0 0 0 0 0 ...");
}
