// use std::io;
use std::fs::read_to_string;

fn take_input(ele: String) -> Vec<i32> {
    // let mut input = String::new();
    // io::stdin()
    //     .read_line(&mut input)
    //     .unwrap();

    let arr: Vec<i32> = ele
                            .trim()
                            .split_whitespace()
                            .map(|x| x.parse().unwrap())
                            .collect();

    return arr;
}


fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string())
    }

    result
}


fn print_arr(arr: Vec<i32>) {
    for ele in arr {
        println!("{}", ele);
    }
    
    println!("----------------");
}

fn main() {
    let file_path = "/home/batman/aoc-2024/day-1/input.txt";
    let set: Vec<String> = read_lines(file_path);
    let mut f_list: Vec<i32> = Vec::new();
    let mut s_list: Vec<i32> = Vec::new();
    for el in set{
        //let x = el;
        // println!("{}", el);
        
        let arr: Vec<i32> = take_input(el);
        f_list.push(arr[0]);
        s_list.push(arr[1]);
        // for ele in arr {
        //     println!("{}", ele);
        // }
        // println!("----------------");
    }

    print_arr(f_list.clone());
    print_arr(s_list.clone());

    f_list.sort();
    s_list.sort();

    print_arr(f_list.clone());
    print_arr(s_list.clone());

    let n = f_list.len();
    let mut ans = 0;

    for i in 0..n {
        ans += (f_list[i] - s_list[i]).abs();
    }

    // let _arr: Vec<usize> = take_input();
    println!("Hello World! {}", ans);
}

/*
3   4
4   3
2   5
1   3
3   9
3   3
*/
