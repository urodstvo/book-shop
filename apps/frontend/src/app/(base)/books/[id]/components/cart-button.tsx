"use client";
import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { useParams, useRouter } from "next/navigation";
import { startTransition } from "react";
import { toast } from "sonner";

async function addToCart(book_id: number) {
  const response = await fetch(API_URL + "/carts", {
    method: "POST",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ book_id }),
  });

  if (!response.ok) {
    toast.error("Не удалось добавить книгу в корзину");
    throw new Error("Failed to add book to cart");
  }

  toast.success("Книга добавлена в корзину");
}

async function deleteFromCart(book_id: number) {
  const response = await fetch(API_URL + "/carts/" + book_id, {
    method: "DELETE",
    mode: "cors",
    credentials: "include",
  });

  if (!response.ok) {
    toast.error("Не удалось удалить книгу из корзины");
    throw new Error("Failed to delete book from cart");
  }

  toast.success("Книга удалена из корзины");
}

export const CartButton = ({ disabled, inCart }: { disabled: boolean; inCart: boolean }) => {
  const { id } = useParams();
  const router = useRouter();

  if (inCart)
    return (
      <Button
        size="lg"
        className="rounded-full  text-white flex-[2] max-w-[600px]"
        disabled={disabled}
        onClick={() => {
          deleteFromCart(Number(id));
          startTransition(() => router.refresh());
        }}
      >
        Удалить из корзины
      </Button>
    );

  return (
    <Button
      size="lg"
      className="rounded-full  text-white flex-[2] max-w-[600px]"
      disabled={disabled}
      onClick={() => {
        addToCart(Number(id));
        startTransition(() => router.refresh());
      }}
    >
      В корзину
    </Button>
  );
};
