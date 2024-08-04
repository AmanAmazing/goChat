BEGIN;

-- -- Insert test users
-- insert into users (email, username, password) values
-- ('john@example.com', 'johndoe', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('jane@example.com', 'janesmith', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('mike@example.com', 'mikejohnson', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('emily@example.com', 'emilybrown', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('david@example.com', 'davidmiller', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('sarah@example.com', 'sarahwilson', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('kevin@example.com', 'kevinjones', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('amy@example.com', 'amytaylor', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('james@example.com', 'jameslee', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a'),
-- ('olivia@example.com', 'oliviaharris', '$2a$14$9gXg5n7LWwSjY/LeCdzKU.V1nFxdmebzgZfuz.h65JVE3bFBzEg6a');
--


CREATE TABLE communities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO communities (name, description, created_at)
VALUES 
('TravelTips', 'A community for sharing travel tips and experiences', CURRENT_TIMESTAMP),
('TechTalk', 'Discuss the latest in technology and gadgets', CURRENT_TIMESTAMP),
('Call Of Duty', 'A community for activision loving degens', CURRENT_TIMESTAMP);


CREATE TABLE threads (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    author INT REFERENCES users(id),
    community_id INT NOT NULL REFERENCES communities(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


INSERT INTO threads (community_id,title, text, author, created_at, updated_at, deleted_at)
VALUES 
    (1,'Why is my poop black', 'Hey guys, not sure why my poop is this black. Will attach an image later', 31, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (3,'Cod Rumours', 'Any new rumours for the upcoming Cod?', 32, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (2,'Third Post', 'This is the text of the third post.', 31, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (1,'Fourth Post', 'This is the text of the fourth post.', 31, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (3,'Fifth Post', 'This is the text of the fifth post.', 36, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);



COMMIT;

