"use client";

import { Input } from "@/components/ui/input";
import { API_URL } from "@/env";
import { useRouter } from "next/navigation";
import { startTransition, useEffect, useState } from "react";
import { toast } from "sonner";

export const QuantityInput = ({
  quantity,
  stock_count,
  session,
  book_id,
}: {
  quantity: number;
  stock_count: number;
  session?: string;
  book_id: number;
}) => {
  const router = useRouter();
  const [value, setValue] = useState(quantity);
  useEffect(() => {
    setValue(quantity);
  }, [quantity]);
  return (
    <Input
      type="number"
      value={value}
      min={1}
      max={stock_count}
      className="w-[60px] text-end [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
      onBlur={() => {
        fetch(API_URL + "/carts/" + book_id, {
          method: "PUT",
          mode: "cors",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
            Cookie: `session_id=${session}`,
          },
          body: JSON.stringify({ quantity: value }),
        })
          .then(() => {
            startTransition(() => {
              router.refresh();
            });
          })
          .catch(() => {
            toast.error("Не удалось добавить книгу в корзину");
          });
      }}
      onChange={(event) => {
        if (!session) toast.error("Авторизуйтеся, чтобы добавлять книги в корзину");
        else {
          const quantity = parseInt(event.target.value);
          if (quantity < 1 || quantity > stock_count) {
            toast.error("Количество книг должно быть от 1 до " + stock_count);
            return;
          }

          setValue(quantity);
        }
      }}
    />
  );
};
