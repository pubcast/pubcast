-- Automatic updated_at
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Groups Table
CREATE TABLE "groups"
(
    "id" SERIAL PRIMARY KEY,
    "slug" VARCHAR(64) NOT NULL UNIQUE,
    "name" VARCHAR(512) NOT NULL,
    "note" TEXT DEFAULT '',
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);

--- Set timestamps for groups
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "groups"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- Organizations Table
CREATE TABLE "organizations"
(
    "id" SERIAL PRIMARY KEY,
    "slug" VARCHAR(64) NOT NULL UNIQUE,
    "name" VARCHAR(512) NOT NULL,
    "note" TEXT DEFAULT '',
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "group_id" INT references groups(id) ON DELETE CASCADE
);

--- Set timestamps for organizations
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "organizations"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- Podcasts Table
CREATE TABLE "podcasts"
(
    "id" SERIAL PRIMARY KEY,
    "slug" VARCHAR(64) NOT NULL UNIQUE,
    "name" VARCHAR(512) NOT NULL,
    "note" TEXT DEFAULT '',
    "thumbnail_url" VARCHAR(2048) NOT NULL,
    "audio_url" VARCHAR(2048) NOT NULL,
    "media_type" VARCHAR(512) NOT NULL,
    "posted_at" TIMESTAMPTZ NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "organization_id" INT references organizations(id) ON DELETE CASCADE
);

--- Set timestamps for podcasts
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "podcasts"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();