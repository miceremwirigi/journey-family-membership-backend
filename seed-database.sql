-- Drop and recreate database
DROP DATABASE IF EXISTS journey_family_db;
CREATE DATABASE journey_family_db;

\c journey_family_db;

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

-- Insert all 12 members at once
INSERT INTO members (id, created_at, updated_at, deleted_at, first_name, middle_name, last_name, email, gender, zone, marital_status, mobile, residence, birthday, wedding_anniversary, special_celebration, special_celebration_description, family_id, family_role, small_group_id, join_date, password, role) VALUES
(1, NOW(), NOW(), NULL, 'Michael', 'Otieno', 'Omondi', 'michael-omondi12@example.com', 'male', 'East Zone', 'married', '254770000012', 'Nairobi', '1985-03-15', '2010-06-20', NULL, NULL, NULL, 'Husband', 5, '2020-01-15', '', 'admin'),
(2, NOW(), NOW(), NULL, 'Nancy', 'Wanjiku', 'Wairimu', 'nancy-wairimu23@example.com', 'female', 'West Zone', 'single', '254770000023', 'Nairobi', '1992-08-22', NULL, NULL, NULL, NULL, NULL, 3, '2021-03-10', '', 'member'),
(3, NOW(), NOW(), NULL, 'Martin', 'Kamau', 'Thuo', 'martin-thuo22@example.com', 'male', 'East Zone', 'married', '254770000022', 'Eldoret', '1980-11-30', '2005-12-15', NULL, NULL, NULL, 'Husband', 5, '2019-06-20', '', 'admin'),
(4, NOW(), NOW(), NULL, 'Joyce', 'Wambui', 'Nyambura', 'joyce-nyambura21@example.com', 'female', 'South Zone', 'divorced', '254770000021', 'Nairobi', '1988-04-12', NULL, '2026-07-01', 'Ministry Anniversary', NULL, NULL, 4, '2022-01-05', '', 'member'),
(5, NOW(), NOW(), NULL, 'Samuel', 'Mwangi', 'Njoroge', 'samuel-njoroge20@example.com', 'male', 'North Zone', 'widowed', '254770000020', 'Eldoret', '1975-09-08', NULL, NULL, NULL, NULL, NULL, 1, '2018-02-14', '', 'member'),
(6, NOW(), NOW(), NULL, 'Margaret', 'Nyokabi', 'Wanjiru', 'margaret-wanjiru19@example.com', 'female', 'Central Zone', 'widowed', '254770000019', 'Eldoret', '1970-02-28', NULL, NULL, NULL, NULL, NULL, 2, '2017-08-30', '', 'member'),
(7, NOW(), NOW(), NULL, 'Charles', 'Maina', 'Kariuki', 'charles-kariuki18@example.com', 'male', 'West Zone', 'widowed', '254770000018', 'Mombasa', '1965-07-14', NULL, NULL, NULL, NULL, NULL, 3, '2016-05-22', '', 'member'),
(8, NOW(), NOW(), NULL, 'Catherine', 'Njeri', 'Muthoni', 'catherine-muthoni17@example.com', 'female', 'East Zone', 'single', '254770000017', 'Nakuru', '1995-12-03', NULL, NULL, NULL, NULL, NULL, 5, '2023-04-18', '', 'member'),
(9, NOW(), NOW(), NULL, 'Robert', 'Ochieng', 'Mutua', 'robert-mutua16@example.com', 'male', 'South Zone', 'married', '254770000016', 'Nairobi', '1983-05-25', '2012-04-10', NULL, NULL, NULL, NULL, 4, '2020-09-01', '', 'member'),
(10, NOW(), NOW(), NULL, 'Jane', 'Akinyi', 'Atieno', 'jane-atieno15@example.com', 'female', 'North Zone', 'widowed', '254770000015', 'Nairobi', '1978-10-17', NULL, NULL, NULL, NULL, NULL, 1, '2019-11-11', '', 'member'),
(11, NOW(), NOW(), NULL, 'Stephen', 'Owino', 'Odhiambo', 'stephen-odhiambo14@example.com', 'male', 'Central Zone', 'single', '254770000014', 'Nakuru', '1990-06-09', NULL, NULL, NULL, NULL, NULL, 2, '2021-07-25', '', 'member'),
(12, NOW(), NOW(), NULL, 'Ruth', 'Adhiambo', 'Achieng', 'ruth-achieng13@example.com', 'female', 'West Zone', 'widowed', '254770000013', 'Mombasa', '1968-01-21', NULL, NULL, NULL, NULL, NULL, 3, '2015-03-08', '', 'member');

-- Insert small groups
INSERT INTO small_groups (id, created_at, updated_at, deleted_at, name, zone, location, description, meeting_day, meeting_time) VALUES
(1, NOW(), NOW(), NULL, 'Garden Estate CG', 'North Zone', 'Garden Estate', 'Community group for Garden Estate residents', 'Wednesday', '6:30 PM'),
(2, NOW(), NOW(), NULL, 'Kilimani Home Fellowship', 'Central Zone', 'Kilimani', 'Fellowship group in Kilimani area', 'Thursday', '7:00 PM'),
(3, NOW(), NOW(), NULL, 'Westlands Bible Study', 'West Zone', 'Westlands', 'Weekly Bible study and prayer group', 'Friday', '6:00 PM'),
(4, NOW(), NOW(), NULL, 'South B Young Adults', 'South Zone', 'South B', 'Young adults fellowship and mentorship', 'Saturday', '4:00 PM'),
(5, NOW(), NOW(), NULL, 'Eastleigh Family CG', 'East Zone', 'Eastleigh', 'Family-focused community group', 'Sunday', '3:00 PM');

-- Insert families WITHOUT constraints first
INSERT INTO families (id, created_at, updated_at, deleted_at, name, head_id) VALUES
(1, NOW(), NOW(), NULL, 'The Omondi Family', 1),
(2, NOW(), NOW(), NULL, 'The Thuo Family', 3);

-- Update members with family references
UPDATE members SET family_id = 1 WHERE id = 1;
UPDATE members SET family_id = 2 WHERE id = 3;

-- Insert visitors
INSERT INTO visitors (id, created_at, updated_at, deleted_at, first_name, middle_name, last_name, gender, contact, email, visits, first_visit, last_visit, interest, status) VALUES
(1, NOW(), NOW(), NULL, 'Philip', 'Kiprono', 'Gachoki', 'male', '254770000024', 'philip-gachoki24@example.com', 2, 'Feb 14, 2026', 'Feb 13, 2026', 'N/A', 'Not Contacted'),
(2, NOW(), NOW(), NULL, 'Agnes', 'Wangari', 'Njeri', 'female', '254770000027', 'agnes-njeri27@example.com', 4, 'Feb 6, 2026', 'Feb 12, 2026', 'N/A', 'Not Contacted'),
(3, NOW(), NOW(), NULL, 'Simon', 'Gitau', 'Maina', 'male', '254770000026', 'simon-maina26@example.com', 3, 'Jan 18, 2026', 'Feb 8, 2026', 'N/A', 'Not Contacted'),
(4, NOW(), NOW(), NULL, 'Esther', 'Nyambura', 'Wambui', 'female', '254770000025', 'esther-wambui25@example.com', 4, 'Jan 25, 2026', 'Feb 7, 2026', 'N/A', 'Not Contacted');

-- Insert messages
INSERT INTO messages (id, created_at, updated_at, deleted_at, date, time, type, recipients, message, delivered) VALUES
(1, NOW(), NOW(), NULL, 'Feb 21, 2026', '06:45 PM', 'sms', 'Sunday service reminder', 'Enim quibusdam fugit illo et et eos quis voluptatem omnis possimus labore voluptates suscipit dolores nesciunt molestias.', 2),
(2, NOW(), NOW(), NULL, 'Feb 21, 2026', '06:45 PM', 'whatsapp', 'Fellowship meeting this Saturday', 'Cupiditate aut consequatur vitae iusto qui sit ut numquam exercitationem nostrum et cumque similique.', 1),
(3, NOW(), NOW(), NULL, 'Feb 21, 2026', '06:45 PM', 'sms', 'Thanks for your service', 'Delectus saepe dicta ratione accusamus deleniti consequuntur et.', 3),
(4, NOW(), NOW(), NULL, 'Feb 21, 2026', '06:45 PM', 'whatsapp', 'Prayer request from the church', 'Earum natus ut vitae numquam ut et voluptas ut excepturi eveniet voluptas quia ex sit consectetur molestias.', 2),
(5, NOW(), NOW(), NULL, 'Feb 21, 2026', '06:45 PM', 'whatsapp', 'Upcoming baptism class', 'Natus et ab perferendis aut ipsam hic sed accusamus autem.', 2);

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

-- Grant permissions
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO journey_family_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO journey_family_user;