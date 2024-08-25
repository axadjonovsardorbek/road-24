-- Insert mock data into the users table
INSERT INTO users (id, first_name, last_name, username, password, role, is_admin, created_at, updated_at, deleted_at)
VALUES
    ('f3b1e5c0-3e6f-11ed-b878-0242ac120002', 'Sardorbek', 'Axadjonov', 'sardorbekaxadjonov', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'superadmin', 'true', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0);





    -- ('2b1a9cf0-3e70-11ed-b878-0242ac120002', 'Jane', 'Smith', 'janesmith', 'securePass!', 'user', 'false', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0),
    -- ('9f8e5c10-3e70-11ed-b878-0242ac120002', 'Alice', 'Johnson', 'alicej', 'alicePass!23', 'admin', 'true', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0),
    -- ('a9b1e3c0-3e70-11ed-b878-0242ac120002', 'Bob', 'Williams', 'bobw', 'bobbySecure$', 'moderator', 'false', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0),
    -- ('f1b4f6d0-3e70-11ed-b878-0242ac120002', 'Charlie', 'Brown', 'charlieb', 'charlieStrong45', 'user', 'false', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0);
