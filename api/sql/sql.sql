CREATE DATABASE IF NOT EXISTS devbook;
use devbook;

drop table if exists publicacoes;
drop table if exists usuarios;
drop table if exists seguidores;

CREATE TABLE usuarios(
    id int primary key auto_increment,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=InnoDB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id) 
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY (usuario_id, seguidor_id)
)ENGINE=InnoDB;

CREATE TABLE publicacoes(
    id int primary key auto_increment,
    titulo varchar(150) not null,
    conteudo text not null,
   
    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
   
    curtidas int default 0,
    criadoEm timestamp default current_timestamp()
)ENGINE=InnoDB;
