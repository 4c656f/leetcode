fn dfs(
    grid: &Vec<Vec<i32>>,
    cache: &mut Vec<i32>,
    x: usize,
    y: usize,
    height: usize,
    width: usize,
) -> i32 {
    let cache_index = x * (width + 1) + y;
    if cache[cache_index] != -1 {
        return cache[cache_index];
    }
    let mut min = -1;

    if x < height {
        min = dfs(grid, cache, x + 1, y, height, width)
    }
    if y < width {
        let res = dfs(grid, cache, x, y + 1, height, width);
        if min == -1 || res < min {
            min = res
        }
    }
    if min == -1 {
        min = 0;
    }
    min += grid[x][y];
    cache[cache_index] = min;

    return min;
}
pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
    let height = grid.len();
    let width = grid[0].len();
    let mut cache: Vec<i32> = vec![-1; height * width];
    return dfs(&grid, &mut cache, 0, 0, height - 1, width - 1);
}
pub fn min_path_sum_dp(mut grid: Vec<Vec<i32>>) -> i32 {
    let height = grid.len();
    let width = grid[0].len();
    let height_edge = height - 1;
    let width_edge = width - 1;
    for x in (0..height).rev() {
        for y in (0..width).rev() {
            let mut min: Option<i32> = Option::None;

            if x != height_edge {
                min = Some(grid[x + 1][y]);
            }
            if y != width_edge {
                match min {
                    Some(val) => {
                        if val > grid[x][y + 1] {
                            min = Some(grid[x][y + 1]);
                        }
                    }
                    None => {
                        min = Some(grid[x][y + 1]);
                    }
                }
            }

            grid[x][y] = min.unwrap_or(0) + grid[x][y];
        }
    }

    return grid[0][0];
}
#[cfg(test)]
mod tests {
    use crate::minimum_path_sum::{min_path_sum, min_path_sum_dp};

    #[test]
    fn basic_matrix_test() {
        let sample = vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1]];
        assert_eq!(min_path_sum(sample), 7);
    }
    #[test]
    fn min_path_sum_dp_test() {
        let sample = vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1]];
        assert_eq!(min_path_sum_dp(sample), 7);
    }
}
