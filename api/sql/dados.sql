insert into usuarios (nome, nick, email, senha) 
values 
('usuario 1', 'usuario_1', 'usuario1@gmail', '$2a$10$qAPRqZ77kemC8tN/7KQf2eNMPTBWKiTARTGwhoVwm5i2iVrcljEFO'),
('usuario 2', 'usuario_2', 'usuario2@gmail', '$2a$10$qAPRqZ77kemC8tN/7KQf2eNMPTBWKiTARTGwhoVwm5i2iVrcljEFO'),
('usuario 3', 'usuario_3', 'usuario3@gmail', '$2a$10$qAPRqZ77kemC8tN/7KQf2eNMPTBWKiTARTGwhoVwm5i2iVrcljEFO');

insert into seguidores (usuario_id, seguidor_id) 
values 
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes (titulo, conteudo, autor_id) 
values 
('Publicação 1', 'Conteúdo da publicação 1', 1),
('Publicação 2', 'Conteúdo da publicação 2', 1),
('Publicação 3', 'Conteúdo da publicação 3', 2),
('Publicação 4', 'Conteúdo da publicação 4', 2),
('Publicação 5', 'Conteúdo da publicação 5', 3),
('Publicação 6', 'Conteúdo da publicação 6', 3);