"use client";

import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { ShoppingBasketIcon } from "lucide-react";
import { redirect, useRouter } from "next/navigation";
import { startTransition } from "react";
import { toast } from "sonner";

export const AddToCartButton = ({
  book_id,
  session_id,
}: {
  book_id: number;
  session_id?: string;
}) => {
  const router = useRouter();

  const addToCart = async () => {
    if (!session_id) {
      return redirect("/login");
    }

    const response = await fetch(API_URL + "/carts", {
      method: "POST",
      mode: "cors",
      credentials: "include",
      headers: {
        Cookie: `session_id=${session_id}`,
      },
      body: JSON.stringify({ book_id }),
    });

    if (!response.ok) {
      toast.error("Ошибка при добавлении книги в корзину");
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
