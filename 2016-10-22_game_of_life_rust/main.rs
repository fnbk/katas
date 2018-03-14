// Conditionally compile `main` only when the test-suite is *not* being run.
//
// to execute unit-tests:
// rustc --test main.rs
// ./main
//
// to execute main function:
// rustc main.rs
// ./main
//
#[cfg(not(test))]
fn main() {
    println!("Hello World");
}

pub fn new_grid(width: usize, height: usize) -> Vec<Vec<i32>> {
    let mut grid = Vec::new();

    for _ in 0..width {
        grid.push(vec![0i32;height])
    }

    grid
}

pub fn number_neighbors_alive(grid: &Vec<Vec<i32>>, x: usize, y: usize) -> i32 {
    let mut alive_around :i32 = 0;

    let tt = neighbors(grid, x as i32, y as i32);

    for t in tt {
        let x = t[0] as usize;
        let y = t[1] as usize;
        alive_around += grid[x][y];
    }

    alive_around
}

pub fn neighbors(grid: &Vec<Vec<i32>>, x: i32, y: i32) -> Vec<Vec<i32>> {
    let xm :i32 = grid.len() as i32;
    let ym :i32 = grid[0].len() as i32;

    let left :i32 = x - 1;
    let right :i32 = x + 1;
    let above :i32 = y - 1;
    let below :i32 = y + 1;

    let mut tt = Vec::new();

    if left >= 0 {
        tt.push(vec![left, y])
    }

    if right < xm {
        tt.push(vec![right, y]);
    }

    if above >= 0 {
        tt.push(vec![x, above]);

        if left >= 0 {
            tt.push(vec![left, above]);
        }

        if right < xm {
            tt.push(vec![right, above]);
        }
    }

    if below < ym {
        tt.push(vec![x, below]);

        if right < xm {
            tt.push(vec![right, below]);
        }

        if left >= 0 {
            tt.push(vec![left, below]);
        }
    }

    tt
}

pub fn calculate_cell_state(old: Vec<Vec<i32>>, x: usize, y: usize) -> i32 {

    let alive_around = number_neighbors_alive(&old, x, y);

    println!("neighbours alive:{:?} ", alive_around);

    let current_state = old[x][y];

    if current_state == 1 {
        if alive_around == 2 {
            return 1;
        }
    }
    if alive_around == 3 {
        return 1;
    }

    0
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn cell_is_dead_when_all_neighbours_are_dead() {
        let grid = new_grid(3, 3);
        assert!(calculate_cell_state(grid, 1, 1) == 0);
    }

    #[test]
    fn living_cell_dies_when_all_neighbours_are_dead() {
        let mut grid = new_grid(3, 3);
        grid[1][1] = 1;
        assert!(calculate_cell_state(grid, 1, 1) == 0);
    }


    #[test]
    fn dead_cell_stays_dead_if_two_neighbors_are_alive() {
        let mut grid = new_grid(3, 3);
        grid[0][0] = 1;
        grid[0][1] = 1;
        assert!(calculate_cell_state(grid, 1, 1) == 0);
    }

    #[test]
    fn living_cell_stays_alive_if_two_neighbors_are_alive() {
        let mut grid = new_grid(3, 3);
        grid[0][0] = 1;
        grid[0][1] = 1;
        grid[1][1] = 1;
        assert!(calculate_cell_state(grid, 1, 1) == 1);
    }

    #[test]
    fn living_cell_stays_alive_if_two_neighbors_are_alive_even_at_the_left_border() {
        let mut grid = new_grid(3, 3);
        grid[0][1] = 1;

        grid[0][0] = 1;
        grid[0][2] = 1;
        assert!(calculate_cell_state(grid, 0, 1) == 1);
    }

    #[test]
    fn living_cell_stays_alive_if_two_neighbors_are_alive_even_at_the_lower_right_border() {
        let mut grid = new_grid(3, 3);
        grid[2][2] = 1;

        grid[1][2] = 1;
        grid[2][1] = 1;
        assert!(calculate_cell_state(grid, 2, 2) == 1);
    }

    #[test]
    fn living_cell_stays_alive_if_three_neighbors_are_alive_even_at_the_border() {
        let mut grid = new_grid(3, 3);
        grid[0][1] = 1;

        grid[0][0] = 1;
        grid[0][2] = 1;
        grid[1][0] = 1;
        assert!(calculate_cell_state(grid, 0, 1) == 1);
    }

    #[test]
    fn dead_cell_reborn_if_three_neighbors_are_alive_even_at_the_border() {
        let mut grid = new_grid(3, 3);
        grid[0][1] = 0;

        grid[0][0] = 1;
        grid[0][2] = 1;
        grid[1][0] = 1;
        assert!(calculate_cell_state(grid, 0, 1) == 1);
    }

    #[test]
    fn living_cell_dies_if_four_neighbors_are_alive() {
        let mut grid = new_grid(3, 3);
        grid[0][1] = 1;

        grid[0][0] = 1;
        grid[0][2] = 1;
        grid[1][0] = 1;
        grid[1][1] = 1;
        assert!(calculate_cell_state(grid, 0, 1) == 0);
    }
}
