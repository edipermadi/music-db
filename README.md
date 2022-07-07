# music-db

Music Database is an attempt to make an encyclopedic database of scales based on [William Zeitler](https://allthescales.org) classification.

## Running test

```shell
make test
```
## Building

To get database up and running

```shell
docker-compose up
```

## Example Queries
### Listing Pitches of a scale by given tonic

```postgresql
SELECT
    sp.id,
    s.name  AS scale,
    p1.name AS tonic,
    p2.name AS pitch
FROM scale_pitches sp
    JOIN scales s   ON s.id  = sp.scale_id
    JOIN pitches p1 ON p1.id = sp.tonic_id
    JOIN pitches p2 ON p2.id = sp.pitch_id
WHERE s.name  = 'Minoric'
    AND p1.name = 'C'
ORDER BY sp.id;
```

### List All Chords for C Major

```postgresql
SELECT 
    s.name     scale,
    p1.name AS tonic,
    p2.name AS pitch,
    c.name  AS chord,
    p3.name AS root
FROM scale_pitch_chords spc
    JOIN scales s ON spc.scale_id = s.id
    JOIN pitches p1 ON spc.tonic_id = p1.id
    JOIN chords c ON spc.chord_id = c.id
    JOIN pitches p3 ON spc.root_id = p3.id
    JOIN pitches p2 ON spc.pitch_id = p2.id
WHERE s.name = 'Ionian'
    AND p1.name = 'C'
```