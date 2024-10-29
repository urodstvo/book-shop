import { cookies } from "next/headers";

import { API_URL } from "@/env";
import { BookCard } from "./book-card";
import { Book, Cart, Genre } from "models";

export const Recomendations = async () => {
  if (!cookies().has("session_id")) return null;

  const session = cookies().get("session_id");

  const response = await fetch(API_URL + "/books/recomendations", {
    credentials: "include",
    headers: {
      Cookie: `session_id=${session?.value}`,
    },
  });

  if (!response.ok) {
    return null;
  }

  const books = (await response.json()) as (Book & {
    userRating: number | null;
    genres: Genre[];
  })[];

  if (books.length === 0) return null;

  const cartsResponse = await fetch(API_URL + "/carts", {
    credentials: "include",
    cache: "no-store",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });
  const carts = (await cartsResponse.json()) as { item: Cart; stock_count: number }[];

  return (
    <section className="flex flex-col items-center mt-20">
      <h2 className="max-w-[1020px] w-full font-bold text-xl">Подобрано для Вас</h2>
      <div className="grid grid-cols-[repeat(auto-fill,180px)] gap-5 p-5 size-fit max-w-[1020px] w-full place-content-center">
        {books.map((book, i) => (
          <BookCard
            key={`book-rec-${i}`}
            {...book}
            inCart={carts.some((cart) => cart.item.book_id === book.id)}
          />
        ))}
      </div>
    </section>
  );
};
