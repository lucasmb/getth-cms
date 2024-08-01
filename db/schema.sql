
-- Tenants Table
CREATE TABLE tenants (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

-- Users Table
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(20) NOT NULL,
    active   BOOLEAN   DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tenant_id INTEGER NOT NULL REFERENCES tenants(tenant_id)
);

-- Indexes for Users Table
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_tenant_id ON users(tenant_id);

-- Posts Table
CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER REFERENCES users(id),
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id)
);

-- Indexes for Posts Table
CREATE INDEX idx_posts_author_id ON posts(author_id);
CREATE INDEX idx_posts_published_date ON posts(published_date);
CREATE INDEX idx_posts_tenant_id ON posts(tenant_id);

-- Categories Table
CREATE TABLE categories (
    id INTEGER PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id)
);

-- Index for Categories Table
CREATE INDEX idx_categories_name ON categories(name);
CREATE INDEX idx_categories_tenant_id ON categories(tenant_id);

-- Tags Table
CREATE TABLE tags (
    id INTEGER PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id)
);

-- Index for Tags Table
CREATE INDEX idx_tags_name ON tags(name);
CREATE INDEX idx_tags_tenant_id ON tags(tenant_id);

-- Post Categories Table (Many-to-Many Relationship between Posts and Categories)
CREATE TABLE post_categories (
    post_id INTEGER REFERENCES posts(id),
    category_id INTEGER REFERENCES categories(id),
    PRIMARY KEY (post_id, category_id)
);

-- Indexes for Post Categories Table
CREATE INDEX idx_post_categories_post_id ON post_categories(post_id);
CREATE INDEX idx_post_categories_category_id ON post_categories(category_id);

-- Post Tags Table (Many-to-Many Relationship between Posts and Tags)
CREATE TABLE post_tags (
    post_id INTEGER REFERENCES posts(id),
    tag_id INTEGER REFERENCES tags(id),
    PRIMARY KEY (post_id, tag_id)
);

-- Indexes for Post Tags Table
CREATE INDEX idx_post_tags_post_id ON post_tags(post_id);
CREATE INDEX idx_post_tags_tag_id ON post_tags(tag_id);

-- Comments Table
CREATE TABLE comments (
    id INTEGER PRIMARY KEY,
    post_id INTEGER REFERENCES posts(id),
    user_id INTEGER REFERENCES users(id),
    content TEXT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id)
);

-- Indexes for Comments Table
CREATE INDEX idx_comments_post_id ON comments(post_id);
CREATE INDEX idx_comments_user_id ON comments(user_id);
CREATE INDEX idx_comments_comment_date ON comments(created_at);
CREATE INDEX idx_comments_tenant_id ON comments(tenant_id);

-- Media Table (Optional, for storing uploaded images and other media files)
CREATE TABLE media (
    id INTEGER PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    file_type VARCHAR(50) NOT NULL,
    file_size BIGINT NOT NULL,
    uploaded_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id)
);

-- Indexes for Media Table
CREATE INDEX idx_media_uploaded_by ON media(uploaded_by);
CREATE INDEX idx_media_upload_date ON media(created_at);
CREATE INDEX idx_media_tenant_id ON media(tenant_id);