import { AspectRatio } from "@/components/ui/aspect-ratio";
import Image from "next/image";
import { BadgeIcon } from "lucide-react";
import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table";
import { Button } from "@/components/ui/button";

import { Cart, Genre, type Book } from "models";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { API_URL } from "@/env";
import { RateButton } from "./components/rate-button";
import { cookies } from "next/headers";
import { Badge } from "@/components/ui/badge";
import { CartButton } from "./components/cart-button";
import { notFound } from "next/navigation";
import Link from "next/link";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Книга | Книжный магазин",
};

async function getBook(id: string) {
  const response = await fetch(API_URL + "/books/" + id, {
    cache: "no-store",
    credentials: "include",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });

  if (!response.ok) {
    if (response.status === 404) {
      notFound();
    }
  }

  return (await response.json()) as Book & { user_rating: number | null; genres: Genre[] };
}

async function getCarts() {
  const response = await fetch(API_URL + "/carts", {
    cache: "no-store",
    credentials: "include",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });

  return (await response.json()) as { item: Cart }[];
}

const Rating = ({
  rating,
  ratingCount,
  userRating,
}: {
  rating: number | null;
  ratingCount: number;
  userRating: number | null;
}) => {
  return (
    <Popover>
      <PopoverTrigger className="relative">
        <BadgeIcon
          size={40}
          fill="currentColor"
          strokeWidth={1}
          className="fill-yellow-600 stroke-white hover:fill-yellow-500"
        />
        <span className="absolute size-full top-0 left-0 flex justify-center items-center font-bold text-sm text-white">
          {rating}
        </span>
      </PopoverTrigger>
      <PopoverContent align="end">
        <div className="flex justify-between">
          <p className="mb-5 text-sm">Общее количество оценок: </p>
          <span className="text-sm">{ratingCount}</span>
        </div>
        <p className="text-sm text-muted-foreground">
          {userRating ? "Ваша оценка" : "Оцените книгу"}
        </p>
        <div className="flex justify-between">
          {new Array(10).fill(0).map((_, i) => (
            <RateButton
              key={"rate" + i}
              rating={i + 1}
              isFilled={i < (userRating ?? 0)}
              session={cookies().get("session_id")?.value}
            />
          ))}
        </div>
      </PopoverContent>
    </Popover>
  );
};

export default async function BookPage({ params: { id } }: { params: { id: string } }) {
  const book = await getBook(id);
  let carts = [] as {
    item: Cart;
  }[];

  if (cookies().has("session_id")) {
    carts = await getCarts();
  }

  return (
    <main className="size-full">
      <section className="w-full grid grid-cols-1 md:grid-cols-[400px_1fr] xl:grid-cols-[400px_800px] grid-rows-1 gap-5 md:gap-0  place-content-center">
        <div className="flex justify-center md:py-5">
          <div className="w-[300px] h-[400px] sticky top-10">
            <AspectRatio ratio={2 / 3}>
              <Image
                src={book.cover}
                alt="Placeholder image"
                fill
                className="size-full rounded-lg object-cover"
              />
              <div className="absolute top-1/2 left-1/2 size-[150%] translate-x-[-50%] translate-y-[-50%] z-[-1] pointer-events-none blur-3xl bg-blue-900 bg-opacity-15" />
              <div className="absolute top-2 right-2">
                <Rating
                  rating={book.rating}
                  ratingCount={book.rating_count}
                  userRating={book.user_rating}
                />
              </div>
            </AspectRatio>
          </div>
        </div>
        <div className="flex flex-col gap-5 md:gap-10">
          <div className="flex flex-col gap-1">
            <h3 className="text-2xl font-medium">{book.name}</h3>
            <h5 className="text-base text-muted-foreground">{book.author}</h5>
          </div>
          <Table className="w-full">
            <TableBody>
              <TableRow>
                <TableCell>Количество страниц</TableCell>
                <TableCell className="flex justify-end">{book.pageCount}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Издательство</TableCell>
                <TableCell className="flex justify-end text-right">{book.published_by}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Опубликовано</TableCell>
                <TableCell className="flex justify-end">
                  {new Date(book.published_at).toLocaleDateString()}
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Количество заказов</TableCell>
                <TableCell className="flex justify-end">{book.orders_count}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Количество на складе</TableCell>
                <TableCell className="flex justify-end">{book.stock_count}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
          <div className="flex items-center flex-wrap gap-5">
            {book.genres.map((genre) => (
              <Badge key={genre.id} className="w-fit text-sm" variant="outline">
                {genre.name}
              </Badge>
            ))}
          </div>
          <span className="flex justify-end text-3xl font-bold">{book.price} ₽</span>
          <div className="flex gap-5 ">
            <CartButton
              disabled={
                !carts.some((cart) => cart.item.book_id === book.id) && book.stock_count === 0
              }
              inCart={carts.some((cart) => cart.item.book_id === book.id)}
            />
            <Button
              size="lg"
              variant="outline"
              className="rounded-full max-w-[300px] flex-1"
              asChild
            >
              <Link href={`/books/${book.id}/demo`} target="_blank">
                Демо
              </Link>
            </Button>
          </div>
          <div>
            <h6 className="text-lg mb-5">Аннотация</h6>
            {book.annotation.split("\n").map((line, index) => (
              <p className="mb-4 flex flex-col gap-1" key={index}>
                {line}
              </p>
            ))}
          </div>
        </div>
      </section>
    </main>
  );
}
