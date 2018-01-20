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
//yield max num of diff in a list
fn num_diff(cigars:&Vec<i32>) -> Vec<Vec<i32>>{
    let mut idx=0;
    let mut diffs_diffs:Vec<Vec<i32>>=Vec::new();
    for _ in cigars.iter(){
        let mut diffs:Vec<i32>=Vec::new();
        //WRONG : it should compare with ALL the elts remaining in the vector, including preceding ones!!
        //two vecs with pop()??
        //is there a simpler way??
        let mut start=cigars[idx];
        //should begin from index
        for cig in cigars.iter(){
            //ternary op in rust?
            let mut diff= if *cig > start {cig-start}else {start-cig};
            diffs.push(diff);
            }
        diffs_diffs.push(diffs);
        //diffs.clear();
        idx+=1;
        }
    diffs_diffs
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let n = parse_input!(input_line, i32);
    let mut cigars:Vec<i32>=Vec::new();
    for i in 0..n as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let lnt = parse_input!(input_line, i32);
        cigars.push(lnt);
    }

    print_err!("{:?} {:?}",cigars,num_diff(&cigars));

    println!("2");
}
