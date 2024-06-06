# Chess

This project is a chess game for the Dero blockchain.

## Acknowledgements

This project builds upon the original work done by [Andy Williams](https://github.com/andydotxyz). We would like to thank them for developing the initial version and making it available as open source. Their work has been instrumental in the development of this Dero version.
You can find the original repository [here](https://github.com/andydotxyz/chess)

## Running

Fallow these commands to run Locally.

    $ git clone https://github.com/Dero-Dev-101/chess.git
    $ cd chess
    $ go build
    $ ./chess

## Installing

To install alongside the other applications on your system use the `fyne` tool.

    $ go get fyne.io/fyne/v2/cmd/fyne
    $ fyne install

## Screenshot

<img src = "/img/screenshot.png" style="max-width: 488px" />

## Status

- [x] Renders board
- [x] Animate moves
- [x] Polish board and colours etc
- [x] Handle user input
- [x] Drag and drop for moves
- [x] Take turns against a computer player
- [x] Save state and restore on app launch

### Dero integration

- [ ] Matrix chat
- [ ] Dero Smart Contracts
- [ ] Group/bundle chess moves into transactions

TODO

- [ ] Add game summary info (who to move etc)
- [ ] Remove dependency on external algorithm
