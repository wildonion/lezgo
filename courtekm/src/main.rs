


// https://drive.google.com/file/d/1Es7Ew8fqkRKGFYfcmFZ8gGJWYOzncA6v/view?usp=sharing

use interfaces::{DeckExt, Round};
use models::{Card, Deck, Player};


mod models;
mod impls;
mod funcs;
mod consts;
mod interfaces;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync + 'static>>{

    let mut init_deck = Deck::new(vec![]);
    init_deck
        .init_new_round()
        .log();

    // a heap boxed type that implemets a trait is a trait object can be used for dynamic dispatching
    let mut round_interfaces: Vec<Box<dyn Round<Deck<Card<String>>>>> = vec![
        Box::new(
            init_deck.clone() // Round is implemented for Deck<Card<String>>
        )
    ];
    // let first_game = &round_interfaces[0]; // can't move out of a shared ref perhaps it's being used by other scopes

    init_deck.set_players(
        vec![
            Player{ nickname: "bot1", score: 0 },
            Player{ nickname: "bot2", score: 0 },
            Player{ nickname: "bot3", score: 0 },
            Player{ nickname: "wildonion", score: 0 }
        ]
    )
    .distribute_cards()
    .start();

    Ok(())

}