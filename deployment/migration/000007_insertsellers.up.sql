INSERT INTO sellers (tenant_id, name, phone, email, created_by_id, updated_by_id) VALUES
    (1, 'Seller A1', '1234567890', 'sellera1@example.com', 1, 1),
    (2, 'Seller A2', '0987654321', 'sellera2@example.com', 2, 2),
    (3, 'Seller B1', '1112223333', 'sellerb1@example.com', 3, 3),
    (4, 'Seller C1', '2223334444', 'sellerc1@example.com', 4, 4),
    (5, 'Seller D1', '3334445555', 'sellerd1@example.com', 5, 5),
    (6, 'Seller E1', '4445556666', 'sellere1@example.com', 6, 6),
    (7, 'Seller F1', '5556667777', 'sellerf1@example.com', 7, 7),
    (8, 'Seller G1', '6667778888', 'sellerg1@example.com', 8, 8),
    (9, 'Seller H1', '7778889999', 'sellerh1@example.com', 9, 9),
    (10, 'Seller I1', '8889990000', 'selleri1@example.com', 10, 10);
ON CONFLICT (tenant_id, email, deleted_at) DO NOTHING;    