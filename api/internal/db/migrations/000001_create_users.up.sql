-- Migração 000001: criação da tabela de usuários
CREATE TABLE IF NOT EXISTS users (
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);
