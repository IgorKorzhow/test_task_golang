CREATE TABLE currency_courses_history (
    id INT AUTO_INCREMENT PRIMARY KEY,
    currency_type VARCHAR(3) NOT NULL,
    currency_scale int NOT NULL,
    currency_name VARCHAR(50) NOT NULL,
    on_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);