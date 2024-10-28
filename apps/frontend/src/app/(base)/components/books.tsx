import { API_URL } from "@/env";
import { Book, Cart } from "models";
import { cookies } from "next/headers";
import { BookCard } from "./book-card";

export const BooksSection = async ({ query = "" }: { query?: string }) => {
  const booksResponse = await fetch(API_URL + "/books?limit=10&page=1" + query, {
    cache: "no-store",
  });
  let books = [] as Book[];

  if (booksResponse.ok) {
    books = (await booksResponse.json()) as Book[];
  }

  let carts: { item: Cart; stock_count: number }[] = [];

  if (cookies().has("session_id")) {
    const cartsResponse = await fetch(API_URL + "/carts", {
      credentials: "include",
      cache: "no-store",
      headers: {
        Cookie: `session_id=${cookies().get("session_id")?.value}`,
      },
    });
    carts = (await cartsResponse.json()) as { item: Cart; stock_count: number }[];
  }

  return (
    <section className="flex flex-col items-center relative">
      <h2 className="max-w-[1020px] w-full font-bold text-xl">Каталог книг</h2>
      <p className="text-muted-foreground max-w-[1020px] w-full text-sm">
        Найдено книг: {books.length}
      </p>
      <div className="grid grid-cols-[repeat(auto-fill,180px)] gap-5 p-5 size-fit max-w-[1020px] w-full place-content-center">
        {books.map((book, i) => (
          <BookCard
            key={`book${i}`}
            {...book}
            inCart={carts.some((cart) => cart.item.book_id === book.id)}
          />
        ))}
      </div>
    </section>
  );
};
