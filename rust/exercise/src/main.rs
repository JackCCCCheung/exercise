fn main() {
    println!("Hello, world!");
}

pub fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {
    let mut ans = 0;
    let mut cnt = 1;
    for f in flowerbed {
        if f == 0 {
            cnt += 1;
        } else {
            ans += if cnt > 2 { (cnt - 1) / 2 } else { 0 };
            if ans >= n {
                return true;
            }
            cnt = 0;
        }
    }
    ans += cnt / 2;
    ans >= n
}
