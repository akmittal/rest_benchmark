-- CreateTable
CREATE TABLE "Movie" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL DEFAULT 'unknown',
    "director" TEXT NOT NULL,
    "year" INTEGER NOT NULL,
    "actors" TEXT NOT NULL,
    "genre" TEXT NOT NULL,
    "language" TEXT NOT NULL,
    "country" TEXT NOT NULL
);
