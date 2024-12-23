CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    userId INT NOT NULL,
    total NUMERIC(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    address TEXT NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users(id)
);
