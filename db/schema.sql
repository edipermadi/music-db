CREATE TABLE pitches
(
    id        BIGSERIAL PRIMARY KEY,
    name      TEXT    NOT NULL,
    number    INTEGER NOT NULL,
    frequency FLOAT   NOT NULL
);

CREATE UNIQUE INDEX ON pitches (name);
CREATE UNIQUE INDEX ON pitches (number);

CREATE TABLE scales
(
    id           BIGSERIAL PRIMARY KEY,
    name         TEXT    NOT NULL,
    cardinality  INTEGER NOT NULL,
    number       INTEGER NOT NULL,
    perfection   INTEGER NOT NULL,
    imperfection INTEGER NOT NULL
);

CREATE UNIQUE INDEX ON scales (name);
CREATE UNIQUE INDEX ON scales (number);

CREATE TABLE scale_pitches
(
    id         BIGSERIAL PRIMARY KEY,
    scale_id   BIGINT  NOT NULL REFERENCES scales (id),
    tonic_id   BIGINT  NOT NULL REFERENCES pitches (id),
    pitch_id   BIGINT  NOT NULL REFERENCES pitches (id),
    is_perfect BOOLEAN NOT NULL
)

CREATE UNIQUE INDEX ON scale_pitches (scale_id, tonic_id, pitch_id);