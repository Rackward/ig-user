-- schema
create database if not exists invest_graph;

-- users
create table invest_graph.users (
id bigint unsigned not null auto_increment,
username varchar(32),
primary key (id)
);

-- investments
-- a unique investment, ie. a certain vanguard fund, a share or a property
create table invest_graph.investments(
    id bigint unsigned not null auto_increment,
    user_id bigint unsigned not null,
    primary key (id),
    constraint `fk_investments_1` foreign key (user_id) references invest_graph.users(id)
);

-- values
-- records the changing value of the investment over time
create table invest_graph.`values` (
    id bigint unsigned not null auto_increment,
    investment_id bigint unsigned not null,
    `date` date not null,
    amount float,
    units int,
    unitPurchasePrice float,
    primary key (id),
    constraint `fk_values_1` foreign key (investment_id) references invest_graph.investments(id)
);

-- transactions
-- records the money you have put into an investment, eg. when you buy more shares in a fund
create table invest_graph.`transactions` (
    id bigint unsigned not null auto_increment,
    investment_id bigint unsigned not null,
    `date` date not null,
    amount float,
    units int,
    unitPrice float,
    type varchar(24), -- this could be one way to record dividends or regular payments
    primary key (id),
    constraint `fk_transactions_1` foreign key (investment_id) references invest_graph.investments(id)
);
