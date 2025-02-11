-- Write your migrate up statements here

    CREATE TABLE game (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        room_id UUID NOT NULL,
        user_id UUID NOT NULL,
        hp INTEGER NOT NULL DEFAULT 100,
        mp INTEGER NOT NULL DEFAULT 100,
        speed INTEGER NOT NULL DEFAULT 10,
        FOREIGN KEY (room_id) REFERENCES room(id) ON DELETE CASCADE
    );

---- create above / drop below ----

    DROP TABLE IF EXISTS game;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
