CREATE TABLE "cattles" (
    "id"                        text        NOT NULL,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "animal_name"               text        NOT NULL,
    "device_id"                 text        NOT NULL,
    "birth_date"                timestamptz NOT NULL,
    "sex"                       text        NOT NULL,
    "color"                     text        NOT NULL,
    "breed"                     text        NOT NULL,
    "birth_weight"              INT         NOT NULL,
    "chest"                     INT         NOT NULL,
    "weaning_weight"            INT         NOT NULL,
    "build_number"              text        NOT NULL,
    "breeder_id"                text        NOT NULL,
    "contact_id"                text        NOT NULL,
    "owner_id"                  text        NOT NULL,

    PRIMARY KEY("id"),
    CONSTRAINT "fk_cattles_cattles"     FOREIGN KEY("breeder_id")   REFERENCES "cattles" ("id")     ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_cattles_contacts"    FOREIGN KEY("contact_id")   REFERENCES "contacts" ("id")    ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_cattles_users"       FOREIGN KEY("owner_id")     REFERENCES "users" ("id")       ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX "idx_cattles_deleted_at" ON "cattles" ("deleted_at");