create table produtos (
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
);

insert into produtos (nome, descricao, preco, quantidade) values ('Camiseta', 'Camiseta preta', 19, 10), ('Fone', 'Muito', 99, 5);