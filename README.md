# ZombEvasion Discord Bot
![forestBoard](/forestBoard.png)
![hospitalBoard](/hospitalBoard.png)
![graveYardBoard](/graveYardBoard.png)


# Instructions
This is a turn base game. One player will be able to make an action at a time and then it will go to the next player.

Goal as a Human:
get to the green safe house on the board.
Goals as a Zombie: 
bite all humans and change them into zombies

Humans start on the blue hex sector and Zombies start on the light green hex sector.

Gray hex is a Dangerous Sector. If you move to this location one of three things can happen.

    40% chance it will sound the alarm at this location.
    40% chance you get to set off an alarm in a different location.
    20% chance to not trip the alarm (as if you landed on a white secure sector)

White hex is a secure sector. If you move to this location no alarm will sound (The message will say "Silence No alarm")

Zombies can choose to either move or attack. Moving works the same as a Human and alarms can trigger. When attacking you pick the location you would like to attack. You will move there and attack the location. A message will show that the sector was attacked. It will let the zombie know if they hit a target (zombie or human). Zombies can move 2 sectors. If they turn a human into a zombie they will be upgraded and will be able to move 3 spaces. Humans can only move 1.

Players that are attacked will respond at the zombie sector.



# Admin Commands

`/setup-server`
> Adds the roles, channels, and images in the game

`/announce-next-game`
> Adds buttons so users can join or leave the queue for the next game

`/start-game <board-name>`
> Select the board and starts the game. Adds the inGame role to players in the queue and marks players as zombies or humans. Sends message for the first turn to player.

`/end-game`
> Ends the current game.


