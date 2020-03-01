CREATE TABLE task
(
    id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
    text            VARCHAR(255) NOT NULL,                # Name of the cat
    status          VARCHAR(150) NOT NULL,                # Name of the cat
    updated_at      TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_at      TIMESTAMP NOT NULL,
    PRIMARY KEY     (id)                                  # Make the id the primary key
);