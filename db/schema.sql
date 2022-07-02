CREATE TABLE pitches
(
    id        BIGSERIAL PRIMARY KEY,
    name      TEXT    NOT NULL,
    number    INTEGER NOT NULL,
    frequency FLOAT   NOT NULL
);

CREATE TABLE scales
(
    id     BIGSERIAL PRIMARY KEY,
    name   TEXT    NOT NULL,
    number INTEGER NOT NULL
);

CREATE TABLE scale_pitches
(
    id       BIGSERIAL PRIMARY KEY,
    scale_id BIGINT NOT NULL REFERENCES scales (id),
    tonic_id BIGINT NOT NULL REFERENCES pitches (id),
    pitch_id BIGINT NOT NULL REFERENCES pitches (id)
)