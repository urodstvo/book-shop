"use client";

import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { TrashIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { startTransition, useState } from "react";
import { toast } from "sonner";

async function removeBook(book_id: number, session: string) {
  const response = await fetch(API_URL + "/carts/" + book_id, {
    method: "DELETE",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      Cookie: `session_id=${session}`,
    },
  });

  if (!response.ok) {
    toast.error("Не удалось удалить книгу из корзины");

    throw new Error("Failed to remove book");
  }
}

export const DeleteButton = ({ book_id, session }: { book_id: number; session?: string }) => {
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
          removeBook(book_id, session)
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
      <TrashIcon strokeWidth={2} />
    </Button>
  );
};
