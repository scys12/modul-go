CREATE TABLE customer
(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR
(255) UNIQUE NOT NULL,
  password VARCHAR
(255) NOT NULL,
  email VARCHAR
(255) UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE seller
(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR
(255) UNIQUE NOT NULL,
  password VARCHAR
(255) NOT NULL,
  email VARCHAR
(255) UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product
(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR
(255),
  price INT,
  description VARCHAR
(255),
  is_bought BOOLEAN DEFAULT FALSE,
  seller_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);