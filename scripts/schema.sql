CREATE TABLE IF NOT EXISTS `ticks`
(
    `timestamp` bigint unsigned NOT NULL,
    `symbol`    varchar(8) NOT NULL,
    `bid`       float     NOT NULL,
    `ask`       float     NOT NULL,
    CONSTRAINT ticks_pk
        PRIMARY KEY (`timestamp`, `symbol`)
    # TODO maybe need some indexes ?
) ENGINE = InnoDB;