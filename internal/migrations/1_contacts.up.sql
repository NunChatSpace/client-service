CREATE TABLE "contacts" (
    "id"                        text        NOT NULL,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "tel"                       text        NOT NULL,
    "address"                   text        NOT NULL,
    "email"                     text        NOT NULL,

    PRIMARY KEY("id")
);

CREATE INDEX "idx_contacts_deleted_at" ON "contacts" ("deleted_at");