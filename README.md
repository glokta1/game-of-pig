## Goal
Write a command-line program that simulates the game of Pig. The strategy for each of the players should be fed as arguments.

## Rules
Pig is a two-player game played with a 6-sided die.
The game has the following rules:
Each turn, a player repeatedly rolls a die until either a 1 is rolled or the player decides to “hold”:

* If the player rolls a 1, they score nothing, and it becomes the next player’s turn.

* If the player rolls any other number, it is added to their turn total, and their turn continues. The player can then decide to continue to roll again or “hold”.

* If a player chooses to “hold”, their current turn total is added to their score, and it becomes the next player’s turn.

The first player to score 100 or more points wins.
For example, the first player, Donald, begins a turn with a roll of 5. Donald could hold and score 5 points but chooses to roll again. Donald rolls a 2 and could hold with a turn total of 7 points but decides to roll again. Donald rolls a 1 and must end his turn without scoring.
The next player, Alexis, rolls the sequence 4-5-3-5-6, after which she chooses to hold and adds her turn total of 23 points to her score.
The game continues, Donald taking the turn. In this turn, Donald rolls 3-2-4 and decides to hold. His total score so far is 9, and the turn passes to Alexis.
The player reaching a score of 100 or more total points wins.