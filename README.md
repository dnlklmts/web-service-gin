API that provides access to a store selling vintage recordings on vinyl. To keep things simple, data will be stored in memory.

Based on tutorial <https://go.dev/doc/tutorial/web-service-gin>

The endpoints:

`/albums`
 - `GET` – Get a list of all albums, returned as JSON.
 - `POST` – Add a new album from request data sent as JSON.

`/albums/:id`
 - `GET` – Get an album by its ID, returning the album data as JSON.
