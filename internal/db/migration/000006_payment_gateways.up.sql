CREATE TABLE IF NOT EXISTS payment_gateways (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    personalized_url VARCHAR(255) NOT NULL,
    tariff INT NOT NULL,
    status INT NOT NULL,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    tariff_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (tariff_id) REFERENCES tariffs(id)
);

