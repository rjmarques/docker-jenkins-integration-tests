CREATE SCHEMA testing
    AUTHORIZATION postgres;

CREATE TABLE testing."my_table"
(
    id bigserial NOT NULL,
    my_text_value text NOT NULL,
    PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE testing."my_table"
    OWNER to postgres;