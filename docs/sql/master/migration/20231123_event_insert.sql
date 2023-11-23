INSERT INTO event (name, reset_hour, repeat_setting, repeat_start_at, start_at, end_at, created_at, updated_at)
VALUES ('Event1', 9, 1, '2023-11-23 08:00:00', NULL, NULL, NOW(), NOW()),
       ('Event2', 9, 0, NULL, '2023-11-23 08:00:00', '2023-12-23 08:00:00', NOW(), NOW());
