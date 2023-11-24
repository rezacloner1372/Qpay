CREATE TABLE IF NOT EXISTS tariffs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price VARCHAR(255) NOT NULL,
    currency VARCHAR(255) NOT NULL,
    validity_days INT NOT NULL,
    isDefault INT NOT NULL DEFAULT 0,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);
