-- Create Books table
CREATE TABLE books (
    id VARCHAR(27),
    title VARCHAR,
    publisher VARCHAR,
    year_published INT,
    call_number VARCHAR UNIQUE,
    cover_picture VARCHAR,
    isbn VARCHAR UNIQUE,
    book_collation TEXT,
    edition VARCHAR,
    description TEXT,
    loc_classification VARCHAR(2),
    quantity INT,
    added_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT books_pkey PRIMARY KEY (id)
)

-- Create Subjects table
CREATE TABLE subjects (
    id SERIAL,
    subject VARCHAR,
    CONSTRAINT subjects_pkey PRIMARY KEY (id)
)

-- Create Books_Subjects table
CREATE TABLE books_subjects (
    book_id VARCHAR(27) REFERENCES books (id),
    subject_id INT REFERENCES subjects (id),
    CONSTRAINT books_subjects_pkey PRIMARY KEY (book_id, subject_id) 
)

-- Create Authors table
CREATE TABLE authors (
    id VARCHAR(27),
    name VARCHAR,
    CONSTRAINT authors_pkey PRIMARY KEY (id)
)

-- Create Books_Authors table
CREATE TABLE books_authors (
    book_id VARCHAR(27) REFERENCES books (id),
    author_id INT REFERENCES authors (id),
    CONSTRAINT books_authors_pkey PRIMARY KEY (book_id, author_id)
)

-- Create Bookcopies table
CREATE TABLE bookcopies (
    id VARCHAR(27),
    barcode VARCHAR UNIQUE,
    book_id VARCHAR(27) REFERENCES books(id),
    condition VARCHAR,
    added_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT bookcopies_pkey PRIMARY KEY (id)
)

-- Create Users table
CREATE TABLE users (
    id VARCHAR(27),
    student_id VARCHAR(8) UNIQUE,
    role VARCHAR,
    username VARCHAR UNIQUE,
    email VARCHAR UNIQUE,
    password VARCHAR,
    total_fine INT,
    registered_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

-- Populate Books table

-- Populate Subjects table

-- Populate Books_Subjects table

-- Populate Authors table

-- Populate Books_Authors table

-- Populate Bookcopies table

-- Populate Users table

