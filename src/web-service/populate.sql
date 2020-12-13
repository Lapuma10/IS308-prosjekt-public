USE project;

INSERT INTO Credentials (email, password) VALUES 
("test1@mail.com", "rootroot"),
("test2@mail.com", "rootroot"),
("test3@mail.com", "rootroot"),
("test4@mail.com", "rootroot"),
("test5@mail.com", "rootroot");

INSERT INTO Users (user_id, shopify_name, api_key_shopify, api_key_fiken, company_slug) VALUES
(1, "", "", "", ""),
(2, "", "", "", ""),
(3, "", "", "", ""),
(4, "", "", "", ""),
(5, "", "", "", "");


INSERT INTO Cronjobs (job_type, last_called_date, interval_days, user_id) VALUES
("update-user", DATE('2020-11-05'), 4, 1),
("update-transaction", DATE('2020-11-06'), 5, 2);

