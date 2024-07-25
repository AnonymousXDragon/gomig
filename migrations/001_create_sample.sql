
-- =migrate:up
CREATE TABLE Test(
    id SERIAL PRIMARY KEY ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE Person(
    id SERIAL PRIMARY KEY ,
    name VARCHAR(200),
    age INTEGER NOT NULL,
    address VARCHAR(500) NOT NULL,
    email VARCHAR(250) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);


-- =migrate:down
DROP TABLE IF EXISTS Test;
DROP TABLE IF EXISTS Person;