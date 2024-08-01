-- Insert sample tenants
INSERT INTO tenants (name) VALUES ('TenantCompany 1');
INSERT INTO tenants (name) VALUES ('Sample Tenant 2');
INSERT INTO tenants (name) VALUES ('Sample Tenant 3');

-- Insert sample users
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Mark', 'Grayson', 'user1', 'user1@example.com', 'password1', 'author', 1);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Nolan', 'Grayson', 'Nolan', 'user2@example.com', 'password2', 'author', 1);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Atom','Eve','Eve', 'user3@example.com', 'password3', 'author', 1);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Adam','West', 'Batman', 'user4@example.com', 'password4', 'author', 1);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Carlos','Rodriguez','Elcarlo', 'user5@example.com', 'password5', 'author', 2);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Ana','Olsen','user6', 'user6@example.com', 'password6', 'author', 2);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('ASd','Dome','user7', 'user7@example.com', 'password7', 'author', 2);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Juan','Perez','user8', 'user8@example.com', 'password8', 'author', 2);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Jaim','Ito','user9', 'user9@example.com', 'password9', 'author', 3);
INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id) VALUES ('Susana','Rodriguez','user10', 'user10@example.com', 'password10', 'author', 3);


-- Insert sample categories
INSERT INTO categories (name, tenant_id) VALUES ('Technology', 1);
INSERT INTO categories (name, tenant_id) VALUES ('Science', 1);
INSERT INTO categories (name, tenant_id) VALUES ('Business', 1);
INSERT INTO categories (name, tenant_id) VALUES ('Health', 2);
INSERT INTO categories (name, tenant_id) VALUES ('Sports', 2);