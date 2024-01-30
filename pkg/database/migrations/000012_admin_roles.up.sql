CREATE TABLE IF NOT EXISTS admin_roles(
    id SERIAL PRIMARY KEY,
    admin_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    CONSTRAINT fk_admin FOREIGN KEY(admin_id) REFERENCES admins(id) ON DELETE CASCADE,
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE
);