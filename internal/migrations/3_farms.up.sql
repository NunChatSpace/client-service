CREATE TABLE "farms"(
    "id"                        TEXT        NOT NULL,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "farmer_id"                 TEXT        NOT NULL,
    "farm_name"                 TEXT        NOT NULL,
    "farm_logo"                 TEXT        NOT NULL,
    "farm_picture"              TEXT        NOT NULL,
    "farm_address"              TEXT        NOT NULL,
    "farm_manager"              TEXT        NOT NULL,
    "animal_husbandry"          TEXT        NOT NULL,

    PRIMARY KEY("id"),
    CONSTRAINT "fk_farms_users_farmer_id"           FOREIGN KEY("farmer_id") REFERENCES "users" ("id")          ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_farms_users_farm_manager"        FOREIGN KEY("farm_manager") REFERENCES "users" ("id")       ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_farms_users_animal_husbandry"    FOREIGN KEY("animal_husbandry") REFERENCES "users" ("id")   ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX "idx_farms_deleted_at" ON "farms" ("deleted_at");