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

CREATE TABLE scale_pitches
(
    id         BIGSERIAL PRIMARY KEY,
    scale_id   BIGINT  NOT NULL REFERENCES scales (id),
    tonic_id   BIGINT  NOT NULL REFERENCES pitches (id),
    pitch_id   BIGINT  NOT NULL REFERENCES pitches (id),
    is_perfect BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX ON scale_pitches (scale_id, tonic_id, pitch_id);

CREATE TABLE scale_centers
(
    id           BIGSERIAL PRIMARY KEY,
    scale_id     BIGINT  NOT NULL REFERENCES scales (id),
    tonic_id     BIGINT  NOT NULL REFERENCES pitches (id),
    balanced     BOOLEAN NOT NULL,
    coordinate_x FLOAT   NOT NULL,
    coordinate_y FLOAT   NOT NULL
);

CREATE UNIQUE INDEX ON scale_centers (scale_id, tonic_id);