CREATE TABLE IF NOT EXISTS borrows (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL,
    borrow_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP NULL,
    is_returned BOOL NOT NULL,
    CONSTRAINT fk_user_id_borrows FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_book_id_borrows FOREIGN KEY (book_id) REFERENCES books(id)
);
