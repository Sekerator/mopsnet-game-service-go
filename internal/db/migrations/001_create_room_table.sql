-- Write your migrate up statements here

    CREATE TABLE room (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        user1_id UUID NOT NULL,
        user2_id UUID,
        status INTEGER NOT NULL DEFAULT 0 CHECK (status BETWEEN 0 AND 9),
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
    CREATE INDEX idx_room_user1_id ON room(user1_id);
    CREATE INDEX idx_room_user2_id ON room(user2_id);

    CREATE OR REPLACE FUNCTION update_updated_at()
        RETURNS TRIGGER AS $$
    BEGIN
        NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
    END;
    $$ LANGUAGE plpgsql;

    CREATE TRIGGER set_updated_at
        BEFORE UPDATE ON room
        FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

---- create above / drop below ----

---- create above / drop below ----

    DROP TRIGGER IF EXISTS set_updated_at ON room;
    DROP FUNCTION IF EXISTS update_updated_at;
    DROP TABLE IF EXISTS room;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
