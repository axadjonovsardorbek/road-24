-- Insert mock data into the users table
INSERT INTO users (id, first_name, last_name, username, password, role, is_admin, created_at, updated_at, deleted_at)
VALUES
    ('f3b1e5c0-3e6f-11ed-b878-0242ac120002', 'Sardorbek', 'Axadjonov', 'sardorbekaxadjonov', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'superadmin', 'true', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0);

-- Mock data for the car_types table
INSERT INTO car_types (id, name) 
VALUES 
    ('b7e22dd8-b4b8-41df-bf93-f455abacbd36', 'Mercedes-Benz'),
    ('c2bfc9f7-e25c-4202-9a62-9371f9c62457', 'BMW'),
    ('d3e2fd76-879f-4f2b-b49d-3f72c4f85e91', 'Audi'),
    ('a4fb60d8-f2b6-49d3-b5cb-6587920d8f71', 'Tesla'),
    ('e4b68ef7-5213-4eb4-839c-4a6b7032b5f3', 'Toyota'),
    ('f6d183c2-2541-48bc-8f14-383760e42d77', 'Ford'),
    ('b9562f53-7da4-474c-9ee0-5debded452e1', 'Chevrolet'),
    ('5c1610ff-1043-40d5-868e-a9e20865a4d4', 'Honda'),
    ('f7d9eab1-dac0-4fa0-8ff7-3bfaa5987315', 'Lexus'),
    ('aa259a4d-0e58-445c-8893-6908b8a0cd0c', 'Nissan'),
    ('4ce1b2c8-0a2b-475a-8597-bcf48ae74397', 'Hyundai'),
    ('d37afe4b-37b9-4ef1-84b2-cf4261197ff4', 'Volkswagen');



INSERT INTO service_types (id, name)
VALUES
    ('6c08ca48-4582-4f5b-9f74-359fd23dbebc', 'Tonirovka'),
    ('7d76c19e-0a51-4ae9-bf2b-796161f10418', 'Moy almashtirish'),
    ('d6d490b6-a979-4f2c-844f-01706e353a6d', 'Dvigatel diagnostikasi'),
    ('44444444-4444-4444-4444-444444444444', 'Avtomobil yuvish'),
    ('7cb277f3-7352-4c85-8525-a1d4eb85c85f', 'Tormoz tizimi tekshiruvi'),
    ('66666666-6666-6666-6666-666666666666', 'Shinalarni almashtirish');
