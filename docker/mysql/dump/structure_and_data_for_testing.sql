use api;

create table orders (
    id int auto_increment primary key,
    payment_method varchar(20),
    total float(10, 2),
    date_purchase timestamp,
    status varchar(20),
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
);

insert into orders (payment_method, total, date_purchase, status) values 
    ('Credito', 152.51, '2025-01-30 11:23:25', 'Confirmacao'),
    ('Debito', 55.03, '2025-01-30 11:24:08', 'Confirmado'),
    ('Dinheiro', 100.15, '2025-01-30 11:24:36', 'Confirmado'),
    ('Dinheiro', 10.40, '2025-01-30 11:30:00', 'Confirmado'),
    ('Debito', 19.96, '2025-01-30 11:52:12', 'Confirmado')