CREATE TABLE pitches
(
    id        BIGSERIAL PRIMARY KEY,
    name      TEXT    NOT NULL,
    number    INTEGER NOT NULL,
    frequency FLOAT   NOT NULL
);

CREATE UNIQUE INDEX ON pitches (name);
CREATE UNIQUE INDEX ON pitches (number);

CREATE TABLE chord_qualities
(
    id               BIGSERIAL PRIMARY KEY,
    name             TEXT    NOT NULL,
    cardinality      INTEGER NOT NULL,
    number           INTEGER NOT NULL,
    pitch_class      JSONB   NOT NULL,
    interval_pattern JSONB   NOT NULL
);

CREATE UNIQUE INDEX ON chord_qualities (name);
CREATE UNIQUE INDEX ON chord_qualities (number);

CREATE TABLE chords
(
    id               BIGSERIAL PRIMARY KEY,
    chord_quality_id BIGINT  NOT NULL REFERENCES chord_qualities (id),
    root_id          BIGINT  NOT NULL REFERENCES pitches (id),
    name             TEXT    NOT NULL,
    number           INTEGER NOT NULL
);

CREATE UNIQUE INDEX ON chords (chord_quality_id, root_id);

CREATE TABLE chord_pitches
(
    id       BIGSERIAL PRIMARY KEY,
    chord_id BIGINT NOT NULL REFERENCES chords (id),
    pitch_id BIGINT NOT NULL REFERENCES pitches (id)
);

CREATE UNIQUE INDEX ON chord_pitches(chord_id, pitch_id);

CREATE TABLE scales
(
    id                         BIGSERIAL PRIMARY KEY,
    name                       TEXT    NOT NULL,
    cardinality                INTEGER NOT NULL,
    number                     INTEGER NOT NULL,
    perfection                 INTEGER NOT NULL,
    imperfection               INTEGER NOT NULL,
    pitch_class                JSONB   NOT NULL,
    interval_pattern           JSONB   NOT NULL,
    rotational_symmetric       BOOLEAN NOT NULL,
    rotational_symmetry_level  INTEGER NOT NULL,
    palindromic                BOOLEAN NOT NULL,
    reflectional_symmetric     BOOLEAN NOT NULL,
    reflectional_symmetry_axes JSONB   NOT NULL,
    balanced                   BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX ON scales (name);
CREATE UNIQUE INDEX ON scales (number);

CREATE TABLE keys
(
    id       BIGSERIAL PRIMARY KEY,
    scale_id BIGINT  NOT NULL REFERENCES scales (id),
    tonic_id BIGINT  NOT NULL REFERENCES pitches (id),
    name     TEXT    NOT NULL,
    number   INTEGER NOT NULL,
    balanced BOOLEAN NOT NULL,
    center_x FLOAT   NOT NULL,
    center_y FLOAT   NOT NULL
);

CREATE UNIQUE INDEX ON keys (scale_id, tonic_id);

CREATE TABLE key_pitches
(
    id         BIGSERIAL PRIMARY KEY,
    key_id     BIGINT  NOT NULL REFERENCES keys (id),
    pitch_id   BIGINT  NOT NULL REFERENCES pitches (id),
    is_perfect BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX ON key_pitches (key_id, pitch_id);

CREATE TABLE key_pitch_chords
(
    id       BIGSERIAL PRIMARY KEY,
    key_id   BIGINT NOT NULL REFERENCES keys (id),
    pitch_id BIGINT NOT NULL REFERENCES pitches (id),
    chord_id BIGINT NOT NULL REFERENCES chords (id)
);

CREATE UNIQUE INDEX ON key_pitch_chords (key_id, pitch_id, chord_id);