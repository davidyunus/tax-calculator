--- Table Definition ---

CREATE TABLE "tax" (
    "taxId" SERIAL PRIMARY KEY NOT NULL,
	"name" CHARACTER varying(255) NOT NULL,
    "taxCode" CHARACTER varying(255) NOT NULL,
    "price" NUMERIC NOT NULL,
    "createdAt" timestamptz NOT NULL,
    "updatedAt" timestamptz NOT NULL,
    "deletedAt" timestamptz
);