# OXO
> A naughts & crosses simulation implemented in Go.

## Contents

- [Introduction](#introduction)
- [Overview](#overview)
- [Usage](#usage)
- [Development](#development)
- [License](#license)

## Introduction
`OXO` is a simulation of the classic game of Noughts and Crosses, also known as Tic-Tac-Toe. 

A game consists of two players.

The objective is for one player to form a line of three of their symbols in a row, column, or diagonal while preventing their opponent from doing the same. The two players take turns to place their assigned marker (an "X" or a "O") on a playing board arranged as a 3 x 3 grid. The game continues until one player achieves a win or no more spaces are available on the board, resulting in a draw.

## Overview

The primary purpose of OXO is to evaluate and compare the effectiveness of different strategies in the game when employed by AI players.

### Key Features:

- **Recorded Games:** OXO can manage any number of games simultaneously, providing a record of each game for later analysis.

- **AI Players:** Players are capable of employing one or more tactics to choose the next empty space from a variety of implemented strategies.

- **Tactics Evaluation:** The application allows users to develop, test, and compare various tactics, offering insights into the most successful strategies.

## Usage

To use OXO, follow these steps:

1. **Install Go** if you haven't already. You can find installation instructions [here](https://go.dev/doc/install).

2. Clone the OXO repository to your local machine:
   ```sh
   git clone https://github.com/GoAtelier/oxo
   ```

3. Build and run the application using the provided commands:
    ```sh
    $ go build
    $ ./oxo
    ```

4. Set up AI players with predefined tactics or develop custom tactics.

5. Simulate games, tournaments, and analyze results to refine your tactics.

## Development

OXO is written in Go and designed to be extensible. If you wish to contribute or make modifications, follow these guidelines:

1. Fork the repository on GitHub.

2. Create a feature branch for your work.

3. Develop and test your changes.

4. Submit a pull request to the main repository for review.

5. Collaborate with the community to enhance the application.

## License

...