


use std::fmt::Debug;
use crate::*;
use rand::Rng;
use rand::seq::SliceRandom;


pub struct NodePerception{
    ltn: i64, // last time this node perception got notified
}

#[derive(Clone, Debug, Default)]
pub struct SignalMemory{
    signal_producer: SignalPorducerInfo,
    signal_type: SignalType,
    last_time_sp_seen: i64,
}

#[derive(Clone, Debug, Default)]
pub struct SignalPorducerInfo;

#[derive(Clone, Debug, Default)]
pub struct SignalType;

#[derive(Clone, Debug, Default)]
pub struct Enemy{
    pub damage_rate: u8
}

pub struct Game<'a>{
    pub players: Vec<Player<'a>>,
    pub game: Box<dyn Round<Deck<Card<String>>>>
}

#[derive(Clone, Debug, Default, PartialEq)]
pub struct Card<R>{
    pub suit: R,
    pub rank: R
}

#[derive(Clone, Debug, Default)]
pub struct Deck<C>{
    pub cards: Vec<C>
}

pub struct Player<'s>{
    pub nickname: &'s str,
    pub score: u16,
}

#[derive(Clone, Debug, Default)]
pub struct Col{
    x: u8,
    y: u8
}

#[derive(Clone, Debug, Default)]
pub struct Row{
    x: u8,
    y: u8
}

#[derive(Clone, Debug, Default)]
pub struct Board<'b>{
    col: &'b [Col],
    row: &'b [Row]
}

#[derive(Clone, Debug, Default)]
pub struct Node<T>{
    // parent   ----> children is strong
    // children ----> parent is weak
    pub value: T, 
    pub parent: Option<std::sync::Arc<std::rc::Weak<Node<T>>>>,
    pub children: Option<std::sync::Arc<tokio::sync::Mutex<Vec<Node<T>>>>>
}

#[derive(Clone, Debug, Default)]
pub struct Graph<T>{
    pub nodes: Vec<Node<T>>
}