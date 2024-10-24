import { BookCardSkeleton } from "./components/book-card";
import { Recomendations } from "./components/recomendations";
import { Suspense } from "react";
import { BooksSection } from "./components/books";
import { OrderSection } from "./components/last-order";

export default async function Home({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}) {
  const s = await searchParams;

  const q = s["q"] as string | "";
  const genres = s["genre"] as string[] | string | undefined;
  const type = s["type"] || "";

  let query = "";
  if (q)
    if (type) query += `&${type}=${q}`;
    else query += `&name=${q}`;
  if (genres)
    if (Array.isArray(genres)) query += `${genres.map((g) => `&genre=${g}`)}}`;
    else query += `&genre=${genres}`;

  return (
    <main className="w-full">
      <Suspense>
        <OrderSection />
      </Suspense>
      <Suspense
        fallback={
          <section className="flex flex-col items-center relative">
            <div className="grid grid-cols-[repeat(auto-fill,180px)] gap-5 p-5 size-fit max-w-[1020px] w-full place-content-center">
              {new Array(5).fill(0).map((_, i) => (
                <BookCardSkeleton key={i} />
              ))}
            </div>
          </section>
        }
      >
        <BooksSection query={query} />
      </Suspense>
      <Suspense>
        <Recomendations />
      </Suspense>
    </main>
  );
}
