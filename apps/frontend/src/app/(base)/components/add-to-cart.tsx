"use client";

import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { ShoppingBasketIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { startTransition } from "react";
import { toast } from "sonner";

export const AddToCartButton = ({ book_id }: { book_id: number }) => {
  const router = useRouter();

  const addToCart = async () => {
    const response = await fetch(API_URL + "/carts", {
      method: "POST",
      mode: "cors",
      credentials: "include",
      body: JSON.stringify({ book_id }),
    });

    if (!response.ok) {
      toast.error("Ошибка при добавлении книги в корзину");
      return;
    }

    toast.success("Книга добавлена в корзину");

    startTransition(() => {
      router.refresh();
    });
  };

  return (
    <Button
      size="icon"
      variant="default"
      className="absolute top-2 right-2 [&_svg]:size-4 size-[24px] "
      onClick={() => addToCart()}
    >
      <ShoppingBasketIcon strokeWidth={1.5} color="white" />
    </Button>
  );
};
