CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    gateway_id INT NOT NULL,
    amount INT NOT NULL,
    status VARCHAR(255) NOT NULL,
    transaction_time TIMESTAMP NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (gateway_id) REFERENCES payment_gateways(id)
    CallbackURL VARCHAR(255) NOT NULL,
    Description VARCHAR(255),
    Email VARCHAR(255),
    Phone VARCHAR(255) NOT NULL,
    Authority VARCHAR(255) UNIQUE NOT NULL,
);
