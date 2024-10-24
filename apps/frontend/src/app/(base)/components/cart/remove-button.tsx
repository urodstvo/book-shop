"use client";
import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { MinusIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { startTransition, useState } from "react";
import { toast } from "sonner";

async function removeQuantity(book_id: number, session: string, quantity: number = 2) {
  const response = await fetch(API_URL + "/carts/" + book_id, {
    method: "PUT",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      Cookie: `session_id=${session}`,
    },
    body: JSON.stringify({ quantity: quantity - 1 }),
  });

  if (!response.ok) {
    if (response.status === 400) {
      toast.error("Невозможно уменьшить количество книги в заказе");
    }

    throw new Error("Failed to remove quantity");
  }
}

export const MinusButton = ({
  book_id,
  session,
  quantity,
}: {
  book_id: number;
  session?: string;
  quantity: number;
}) => {
  const router = useRouter();

  const [isPending, setIsPending] = useState(false);

  return (
    <Button
      size="icon"
      variant="ghost"
      disabled={isPending}
      onClick={() => {
        if (!session) toast.error("Авторизуйтеся, чтобы добавлять книги в корзину");
        else {
          setIsPending(true);
          removeQuantity(book_id, session, quantity)
            .then(() => {
              startTransition(() => {
                router.refresh();
              });
              setIsPending(false);
            })
            .catch(() => {
              setIsPending(false);
            });
        }
      }}
    >
      <MinusIcon strokeWidth={2} />
    </Button>
  );
};
