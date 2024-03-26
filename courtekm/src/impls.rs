


use rand::{thread_rng, Rng, SeedableRng};
use crate::*;
use crate::models::*;
use rand::seq::SliceRandom;
use crate::interfaces::*;
use self::consts::gen_random_chars;


impl Enemy{
    pub fn new() -> Self{
        Self { 
            damage_rate: {
                let mut rng = rand::thread_rng();
                let rate = rng.gen_range(0..10);
                rate
            } 
        }
    }
}

impl<R> Card<R>{

    pub fn new(suit: R, rank: R) -> Self{
        Self { suit, rank }
    }
}

impl<C> Deck<C>{

    pub fn new(cards: Vec<C>) -> Self{
        Self { cards }
    }

    pub fn distribute_cards(&mut self) -> Deck<Card<String>> {

        let _ = &self.cards.shuffle(&mut thread_rng());
        let p1_cards = &self.cards[0..4];
        let p3_cards = &self.cards[4..9];
        let p4_cards = &self.cards[9..14];
        let p5_cards = &self.cards[14..19];

        // 0 - distribute 52 cards among players
        // 1 - shuffle cards
        // 2 - select a dealer and the rule
        // 3 - spread cards
        // ...

        todo!()

    }
}

impl DeckExt<Deck<Card<String>>> for Deck<Card<String>>{
    fn log(&self) {
        let cards = &self.cards;
        for card in cards{
            println!("suit: {} \t | rank: {}", card.suit, card.rank);
        }
    }
}

impl Round<Deck<Card<String>>> for Deck<Card<String>>{

    fn get_cards(&self) -> Vec<Card<String>>{
        let cards = self.cards.clone();
        cards
    }

    fn set_players(&mut self,  players: Vec<Player>) -> Deck<Card<String>> {
        
        self.clone()
    }

    fn start(&mut self){

        loop{
            
        }

    } 

    fn init_new_round(&mut self) -> Deck<Card<String>>{
        
        let suits = vec![
            String::from("Diamonds"), // khesht
            String::from("Hearts"), // ghalb
            String::from("Clubs"), // gishniz
            String::from("Spades") // pick 
        ];
        let ranks = vec![
            String::from("A"), String::from("2"), String::from("3"), 
            String::from("4"), String::from("5"), String::from("6"), 
            String::from("7"), String::from("8"), String::from("9"), 
            String::from("10"), String::from("J"), String::from("Q"), 
            String::from("K")
        ];
        
        let mut deck = Deck::new(vec![]);
        for s in suits.clone(){
            for r in ranks.clone(){
               
                let random_chars = gen_random_chars(10);
                let hash_of_random_chars = wallexerr::misc::Wallet::generate_sha256_from(&random_chars);
                let mut crypto_seeded_rng = rand_chacha::ChaCha20Rng::from_seed(hash_of_random_chars);
               
                let random_suit = suits.choose(&mut crypto_seeded_rng);
                let random_rank = ranks.choose(&mut crypto_seeded_rng);
               
                let new_card = Card::new(
                    random_suit.unwrap().to_owned(),
                    random_rank.unwrap().to_owned()
                );
                if !deck.cards.contains(&new_card){
                    deck.cards.push(
                        // select random suit and rank if not already selected
                        new_card
                    )        
                }
            }
        }

        self.cards = deck.clone().cards;

        deck

    }

}