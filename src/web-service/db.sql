Create database project;
use project;

CREATE TABLE Credentials
(
user_id INTEGER AUTO_INCREMENT PRIMARY KEY NOT NULL,
email  varchar(255) unique,
password varchar(255)
);


Create table Users
(
user_id INTEGER,
shopify_name VARCHAR(255) NOT NULL,
api_key_shopify VARCHAR(255) NOT NULL,
api_key_fiken VARCHAR(255) NOT NULL,
company_slug VARCHAR(255) NOT NULL,
FOREIGN KEY (user_id) REFERENCES Credentials(user_id)
);



Create table Cronjobs
(
cronjob_id INTEGER auto_increment PRIMARY KEY NOT NULL,
job_type VARCHAR(255) NOT NULL,
last_called_date date,
interval_days INTEGER NOT NULL,
user_id INTEGER NOT NULL,
FOREIGN KEY (user_id) REFERENCES Credentials(user_id)
);
