-- Create table users
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE,
  bio TEXT,
  name VARCHAR(255),
  userName VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create table blogs
CREATE TABLE blogs (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users (id),
  title VARCHAR(255),
  sub_title VARCHAR(255),
  body TEXT,
  views_count INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create table likes
CREATE TABLE likes (
  id SERIAL PRIMARY KEY,
  blog_id INTEGER REFERENCES blogs (id),
  user_id INTEGER REFERENCES users (id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create table authentications
CREATE TABLE authentications (
  id SERIAL PRIMARY KEY,
  token VARCHAR(255),
  user_id INTEGER REFERENCES users (id),
  expiry_at TIMESTAMP,
  is_active BOOLEAN,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create table otps
CREATE TABLE otps (
  id SERIAL PRIMARY KEY,
  otp INTEGER UNIQUE,
  failed_reason VARCHAR(255),
  user_id INTEGER REFERENCES users (id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
