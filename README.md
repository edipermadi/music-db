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

## Listing Pitches of a scale by given tonic

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
WHERE
    s.name='Minoric' AND
    p1.name = 'C'
ORDER BY sp.id;
```
