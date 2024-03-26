

use crate::*;
use crate::models::*;
use crate::interfaces::*;




pub fn build_board(){    

    // every node has a board instance as its value
    let mut board = Board::default();
    let mut node = Node::<Board<'_>>::default();
    node.value = board.clone();

}

pub fn start_cp_game(deck: Deck<Card<String>>){



}

pub fn find_optimised_path(){

}