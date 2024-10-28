import { AspectRatio } from "@/components/ui/aspect-ratio";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { ShoppingBasketIcon, StarIcon } from "lucide-react";
import { Book } from "models";
import Link from "next/link";
import Image from "next/image";
import { AddToCartButton } from "./add-to-cart";
import { cookies } from "next/headers";

export const BookCardSkeleton = () => {
  return (
    <Card className="border-transparent w-[180px] p-[10px] h-fit shadow-none bg-transparent">
      <CardContent className="p-0 flex flex-col gap-2">
        <div className="relative">
          <Link href="/books/1">
            <AspectRatio ratio={3 / 4}>
              <Skeleton className="size-full rounded-md" />
            </AspectRatio>
          </Link>
          <Button
            size="icon"
            variant="default"
            className="absolute top-2 right-2 [&_svg]:size-4 size-[24px] "
          >
            <ShoppingBasketIcon strokeWidth={1.5} />
          </Button>
        </div>
        <h4 className="scroll-m-20 text-xl font-semibold mt-2">
          <Skeleton className="w-full h-5" />
        </h4>
        <div className="flex gap-2 w-full h-5">
          <span className="h-full flex-[2]">
            <Skeleton className="size-full" />
          </span>
          <span className="h-full flex-1">
            <Skeleton className="size-full" />
          </span>
        </div>
      </CardContent>
    </Card>
  );
};

export const BookCard = (
  book: Book & {
    inCart: boolean;
  }
) => {
  return (
    <Card className="border-transparent w-[180px] p-[10px] h-fit shadow-none bg-transparent">
      <CardContent className="p-0 flex flex-col gap-2">
        <div className="relative">
          <Link href={`/books/${book.id}`}>
            <AspectRatio ratio={3 / 4}>
              <Image
                src={book.cover}
                alt={book.name}
                fill
                className="h-full w-full rounded-md object-cover"
              />
            </AspectRatio>
          </Link>
          {!book.inCart && book.stock_count > 0 && cookies().has("session_id") && (
            <AddToCartButton book_id={book.id} />
          )}
        </div>
        <h4 className="scroll-m-20 text-sm font-semibold mt-2">
          {book.author}: {book.name}
        </h4>
        <div className="flex gap-2 w-full h-5">
          <span className="h-full flex-[2] text-sm">{book.price} ₽</span>
          {book.rating && (
            <span className="h-full flex-1 flex items-center text-sm gap-1">
              <StarIcon size={16} fill="white" />
              {book.rating}
            </span>
          )}
        </div>
      </CardContent>
    </Card>
  );
};
