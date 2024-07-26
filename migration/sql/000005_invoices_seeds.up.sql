INSERT INTO invoices (id, subject, customer_id, due_date, status, created_at, updated_at, deleted_at)
VALUES
    (1, 'Spring Marketing Campaign', 101, '2024-07-15 00:00:00', 'Paid', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (2, 'Summer Marketing Campaign', 102, '2024-07-16 00:00:00', 'Unpaid', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (3, 'Autumn Marketing Campaign', 103, '2024-07-17 00:00:00', 'Paid', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);
