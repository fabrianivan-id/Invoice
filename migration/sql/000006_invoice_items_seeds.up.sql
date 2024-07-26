INSERT INTO invoice_items (id, invoice_id, product_id, item_name, quantity, unit_price, created_at, updated_at, deleted_at)
VALUES
    (1, 1, 201, 'Product A', 10, 15.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (2, 1, 202, 'Product B', 5, 30.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (3, 2, 203, 'Product C', 7, 20.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (4, 2, 204, 'Product D', 3, 50.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (5, 3, 205, 'Product E', 8, 25.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (6, 3, 206, 'Product F', 4, 35.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);

