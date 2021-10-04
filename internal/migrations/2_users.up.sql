CREATE TABLE "users" (
    "id"                        text        NOT NULL,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "name"                      text        NOT NULL,
    "type"                      text        NOT NULL,
    "contact_id"                text        NOT NULL,

    PRIMARY KEY("id"),
    CONSTRAINT "fk_users_contacts" FOREIGN KEY("contact_id") REFERENCES "contacts" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
