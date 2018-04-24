CREATE TABLE IF NOT EXISTS `coinprice` (
  `coinpriceno` int(11) NOT NULL AUTO_INCREMENT,
  `pricedate` int(11) NOT NULL,
  `no` int(11) NOT NULL,
  `name` varchar(40) NOT NULL,
  `symbol` varchar(10) NOT NULL,
  `price` float NOT NULL,
  `marketcap` bigint(20) NOT NULL,
  `circulating_supply` bigint(20) NOT NULL,
  `volume` bigint(20) NOT NULL,
  PRIMARY KEY (`coinpriceno`),
  UNIQUE KEY `pricedateno` (`pricedate`,`no`),
  KEY `symbol` (`symbol`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS `globaldata` (
  `globaldatano` int(11) NOT NULL AUTO_INCREMENT,
  `pricedate` int(11) NOT NULL,
  `totalmarketcap` bigint(20) NOT NULL,
  `bitcoinmarketcap` bigint(20) NOT NULL,
  `volume` bigint(20) NOT NULL,
  PRIMARY KEY (`globaldatano`),
  UNIQUE KEY (`pricedate`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=1;


