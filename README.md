# IMPORTANT
This little program is for funsies and only written to practice implementing a heap in a program. As of 12/13/23 that's pretty much all I needed it to do. Maybe some day I will add a server and give the ability to register and actually level... again could be fun and this project will stay as something for funsies.... funsies!

# Match-MakingBazaar

Match-MakingBazaar is a command-line multiplayer game matchmaking system designed to bring a seamless and competitive gaming experience. The application utilizes a max-heap to efficiently match players based on their skill levels, ensuring fair and engaging matchups.

## Features

- **Player Registration:** Players can register with a username and skill level, with the information stored in a max-heap.
   - This is done in the code. There is no actual player registration... although it could be added later on.
- **Joining Match Queue:** Players can join the matchmaking queue to find opponents, with the system pairing players with similar skill levels using the max-heap.
- **Matchmaking Algorithm:** The system extracts the top two players from the max-heap for fair matchups and displays match details.
- **Player Statistics:** Display and update player statistics, including matches played, won, and lost.
- **Interactive Interface:** An intuitive command-line interface allows users to register, join queues, view statistics, and engage in matches seamlessly.

## Example Workflow

1. **Player Registration:** Players register with usernames and skill levels. (See above. Players are hard-coded at the moment)
2. **Join Matchmaking Queue:** Players join the matchmaking queue for fair matchups.
3. **Match Found:** The system pairs players based on skill levels for an engaging match.
4. **Update Player Statistics:** Player statistics are updated after each match.
5. **View Player Statistics:** Players can view their individual statistics and track their progress.

## Usage

1. Clone the repository.
   ```bash
   git clone https://github.com/cs50-romain/Match-MakingBazaar.git
