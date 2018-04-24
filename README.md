# CryptoSim
Crypto Profit Simulator using data from coinmarketcap.com historical snapshots

The simulation checks profit when you invest in top N currencies equally for W weeks.


Related posts & results
-----------------------
* [part 1](https://steemit.com/cryptocurrency/@mix1009/project-crypto-investing-profit-simulation-from-historical-data-1)
* [part 2](https://steemit.com/cryptocurrency/@mix1009/project-crypto-investing-profit-simulation-from-historical-data-2)
* [interactive 3d chart](https://mix1009.com/crypto/simulation)

Requirements
------------
* go
* make
* MySQL server

Installation
------------
> git clone https://github.com/mix1009/CryptoSim

> cd CryptoSim

    get source from github

> make install-packages

    install required go packages

> make

    Build all executables.

> vi config.json

    edit config.json and fill database name, user, password.

Running
-------

> ./download

    downloads historical snapshots to data directory.

> ./parse

    parse historical snapshots from data directory and save it to database

> ./simulate

    do the simulation. Edit main_simulate.go to change simulation parameters.

> ./simulate > out.csv

    simulate output results to stdout. save data to out.csv

> ./simulate | tee out2.csv

    save data to out2.csv and show results to stdout at the same time.

License
-------
MIT License

