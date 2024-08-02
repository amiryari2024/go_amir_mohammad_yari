INSERT INTO operators (id, name, created_at, updated_at) VALUES
(1, 'Operator 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Operator 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Operator 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Operator 4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Operator 5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO product_types (id, name, created_at, updated_at) VALUES
(1, 'Product Type 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Product Type 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Product Type 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
(1, 1, 3, 'P001', 'Product 1A', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 1, 3, 'P002', 'Product 1B', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
(3, 2, 1, 'P003', 'Product 2A', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 2, 1, 'P004', 'Product 2B', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 2, 1, 'P005', 'Product 2C', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
(6, 3, 4, 'P006', 'Product 3A', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(7, 3, 4, 'P007', 'Product 3B', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(8, 3, 4, 'P008', 'Product 3C', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO product_descriptions (id, description, created_at, updated_at) VALUES
(1, 'Description for Product 1A', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Description for Product 1B', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Description for Product 2A', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Description for Product 2B', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Description for Product 2C', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 'Description for Product 3A', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(7, 'Description for Product 3B', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(8, 'Description for Product 3C', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO payment_methods (id, name, status, created_at, updated_at) VALUES
(1, 'Credit Card', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'PayPal', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Bank Transfer', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO users (id, status, dob, gender, created_at, updated_at) VALUES
(1, 1, '1990-01-01', 'M', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 1, '1991-02-02', 'F', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 1, '1992-03-03', 'M', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 1, '1993-04-04', 'F', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 1, '1994-05-05', 'M', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at) VALUES
-- User 1
(1, 1, 1, 'completed', 3, 30.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 1, 2, 'completed', 2, 20.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 1, 3, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
-- User 2
(4, 2, 1, 'completed', 3, 30.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 2, 2, 'completed', 2, 20.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 2, 3, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
-- User 3
(7, 3, 1, 'completed', 3, 30.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(8, 3, 2, 'completed', 2, 20.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(9, 3, 3, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
-- User 4
(10, 4, 1, 'completed', 3, 30.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(11, 4, 2, 'completed', 2, 20.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(12, 4, 3, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
-- User 5
(13, 5, 1, 'completed', 3, 30.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(14, 5, 2, 'completed', 2, 20.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(15, 5, 3, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, created_at, updated_at) VALUES
-- Transaction 1
(1, 1, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 2, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 3, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
-- Transaction 2
(2, 4, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 5, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 6, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
-- Transaction 3
(3, 7, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 8, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 1, 'completed', 1, 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
-- Additional transactions follow the same pattern
-- ...

SELECT id, gender, created_at, updated_at FROM users WHERE gender = 'M';

SELECT * FROM products WHERE id = 3;

SELECT * FROM users WHERE created_at >= NOW() - INTERVAL 7 DAY AND username LIKE '%a%';

SELECT COUNT(*) FROM users WHERE gender = 'F';

SELECT * FROM users ORDER BY username ASC;

SELECT * FROM products LIMIT 5;

UPDATE products SET name = 'product dummy' WHERE id = 1;

UPDATE transaction_details SET qty = 3 WHERE product_id = 1;

DELETE FROM products WHERE id = 1;

DELETE FROM products WHERE product_type_id = 1;

SELECT * FROM transactions WHERE user_id IN (1, 2);

SELECT SUM(total_price) FROM transactions WHERE user_id = 1;

SELECT COUNT(*) FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
WHERE p.product_type_id = 2;

SELECT p.*, pt.name AS product_type_name
FROM products p
JOIN product_types pt ON p.product_type_id = pt.id;

SELECT t.*, p.name AS product_name, u.username AS user_name
FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
JOIN users u ON t.user_id = u.id;

CREATE TRIGGER after_transaction_delete
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
  DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END;

CREATE TRIGGER after_transaction_detail_delete
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
  UPDATE transactions
  SET total_qty = total_qty - OLD.qty
  WHERE id = OLD.transaction_id;
END;

SELECT * FROM products WHERE id NOT IN (SELECT DISTINCT product_id FROM transaction_details);