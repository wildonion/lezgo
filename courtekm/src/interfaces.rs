

use crate::*;
use self::models::Player;


pub trait Round<G>{
    fn get_cards(&self) -> Vec<Card<String>>;
    fn init_new_round(&mut self) -> G; // without &mut self it won't be object safe
    fn set_players(&mut self,  players: Vec<Player>) -> G;
    fn start(&mut self);
}

pub trait DeckExt<G>{
    fn log(&self);
}