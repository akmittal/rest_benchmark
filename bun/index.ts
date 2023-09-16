import { PrismaClient } from "@prisma/client";
import { Hono } from "hono";
const app = new Hono({});

const prisma = new PrismaClient();

app.get("/movie", async (c) => {
  let limit = c.req.query("limit") ? Number(c.req.query("limit")) : 100;
  return c.json(await prisma.movie.findMany({ take: limit }));
});

app.post("/movie", async (c) => {
  const newMovie = await prisma.movie.create({
    data: {
      name: "Titanic",
      year: 1997,
      actors: "Leonardo DiCaprio, Kate Winslet, Billy Zane, Kathy Bates",
      director: "James Cameron",
      genre: "Drama, Romance",
      country: "USA",
      language: "English, Swedish, Italian, French",
    },
  });
  return c.json(newMovie);
});



export default app;
