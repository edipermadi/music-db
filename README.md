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
### Listing Pitches of a key

```postgresql
SELECT k.name AS key,
       p.name    AS pitch,
       kp.pitch_id
FROM key_pitches kp
         JOIN keys k ON kp.key_id = k.id
         JOIN pitches p ON kp.pitch_id = p.id
WHERE k.name = 'CNaturalIonian'
ORDER BY kp.id;
```

### List All Chords for C Major

```postgresql
SELECT k.name AS key,
       p.name AS pitch,
       c.name AS chord
FROM key_pitch_chords spc
         JOIN keys k ON spc.key_id = k.id
         JOIN chords c ON spc.chord_id = c.id
         JOIN pitches p ON spc.pitch_id = p.id
WHERE k.name = 'CNaturalIonian';
```