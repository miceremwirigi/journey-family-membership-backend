-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create members table FIRST (no dependencies)
CREATE TABLE members (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    first_name VARCHAR(255),
    middle_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    gender VARCHAR(255),
    zone VARCHAR(255),
    marital_status VARCHAR(255),
    mobile VARCHAR(255),
    residence VARCHAR(255),
    birthday VARCHAR(255),
    wedding_anniversary VARCHAR(255),
    special_celebration VARCHAR(255),
    special_celebration_description VARCHAR(255),
    family_id BIGINT,
    family_role VARCHAR(255),
    small_group_id BIGINT,
    join_date VARCHAR(255),
    password VARCHAR(255),
    role VARCHAR(255)
);

-- Create families table
CREATE TABLE families (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(255),
    head_id BIGINT
);

-- Create small_groups table
CREATE TABLE small_groups (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(255),
    zone VARCHAR(255),
    location VARCHAR(255),
    leader_id BIGINT,
    description VARCHAR(255),
    meeting_day VARCHAR(255),
    meeting_time VARCHAR(255)
);

-- Create visitors table
CREATE TABLE visitors (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    first_name VARCHAR(255),
    middle_name VARCHAR(255),
    last_name VARCHAR(255),
    gender VARCHAR(255),
    contact VARCHAR(255),
    email VARCHAR(255),
    visits INT,
    first_visit VARCHAR(255),
    last_visit VARCHAR(255),
    interest VARCHAR(255),
    status VARCHAR(255)
);

-- Create messages table
CREATE TABLE messages (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    date VARCHAR(255),
    time VARCHAR(255),
    type VARCHAR(255),
    recipients VARCHAR(255),
    message TEXT,
    delivered INT
);

-- Create events table
CREATE TABLE events (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    title VARCHAR(255),
    category VARCHAR(255),
    date_time VARCHAR(255),
    location VARCHAR(255),
    status VARCHAR(255)
);

-- Add foreign key constraints AFTER all data is inserted
ALTER TABLE members ADD CONSTRAINT fk_members_family 
    FOREIGN KEY (family_id) REFERENCES families(id) ON DELETE SET NULL;

ALTER TABLE members ADD CONSTRAINT fk_members_small_group 
    FOREIGN KEY (small_group_id) REFERENCES small_groups(id) ON DELETE SET NULL;

ALTER TABLE families ADD CONSTRAINT fk_families_head_id 
    FOREIGN KEY (head_id) REFERENCES members(id) ON DELETE SET NULL;

ALTER TABLE small_groups ADD CONSTRAINT fk_small_groups_leader 
    FOREIGN KEY (leader_id) REFERENCES members(id) ON DELETE SET NULL;

-- Create indexes
CREATE INDEX idx_members_email ON members(email);
CREATE INDEX idx_members_zone ON members(zone);
CREATE INDEX idx_members_family_id ON members(family_id);
CREATE INDEX idx_members_small_group_id ON members(small_group_id);
CREATE INDEX idx_members_deleted_at ON members(deleted_at);
CREATE INDEX idx_families_head_id ON families(head_id);
CREATE INDEX idx_families_deleted_at ON families(deleted_at);
CREATE INDEX idx_small_groups_zone ON small_groups(zone);
CREATE INDEX idx_small_groups_deleted_at ON small_groups(deleted_at);
CREATE INDEX idx_visitors_email ON visitors(email);
CREATE INDEX idx_visitors_deleted_at ON visitors(deleted_at);
CREATE INDEX idx_messages_date ON messages(date);
CREATE INDEX idx_messages_deleted_at ON messages(deleted_at);
CREATE INDEX idx_events_deleted_at ON events(deleted_at);
