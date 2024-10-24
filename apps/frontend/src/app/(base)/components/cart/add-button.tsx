"use client";

import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { PlusIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { startTransition, useState } from "react";
import { toast } from "sonner";

async function addQuantity(book_id: number, session: string, quantity: number = 1) {
  const response = await fetch(API_URL + "/carts/" + book_id, {
    method: "PUT",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      Cookie: `session_id=${session}`,
    },
    body: JSON.stringify({ quantity: quantity + 1 }),
  });

  if (!response.ok) {
    if (response.status === 400) {
      toast.error("Невозможно увеличить количество книги в заказе");
    }

    throw new Error("Failed to add quantity");
  }
}

export const PlusButton = ({
  book_id,
  session,
  quantity,
}: {
  book_id: number;
  session?: string;
  quantity: number;
}) => {
  const [isPending, setIsPending] = useState(false);
  const router = useRouter();

  return (
    <Button
      size="icon"
      variant="ghost"
      disabled={isPending}
      onClick={() => {
        if (!session) toast.error("Авторизуйтеся, чтобы добавлять книги в корзину");
        else {
          setIsPending(true);
          addQuantity(book_id, session, quantity)
            .then(() => {
              router.refresh();
              startTransition(() => {});
              setIsPending(false);
            })
            .catch(() => {
              setIsPending(false);
            });
        }
      }}
    >
      <PlusIcon strokeWidth={2} />
    </Button>
  );
};
