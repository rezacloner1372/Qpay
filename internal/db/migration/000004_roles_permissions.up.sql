CREATE TABLE IF NOT EXISTS roles_permissions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    role_id INT,
    permission_id INT,
    guard_name VARCHAR(255) UNIQUE,
    FOREIGN KEY (role_id) REFERENCES roles(id),
    FOREIGN KEY (permission_id) REFERENCES permissions(id)
);
