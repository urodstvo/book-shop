"use client";

import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { cn } from "@/lib/utils";
import { StarIcon } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import { startTransition } from "react";

const Rate = async ({ rating, book_id }: { rating: number; book_id: number }) => {
  const response = await fetch(API_URL + "/books/" + book_id + "/rate/" + rating, {
    method: "PUT",
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to rate book");
  }
};

export const RateButton = ({
  rating,
  isFilled,
  session,
}: {
  isFilled: boolean;
  rating: number;
  session?: string;
}) => {
  const { id } = useParams();
  const router = useRouter();

  const handleClick = async () => {
    if (session) {
      await Rate({ rating, book_id: Number(id) });
      startTransition(() => {
        router.refresh();
      });
    } else router.push("/login");
  };

  return (
    <Button variant={null} size="icon" onClick={handleClick}>
      <StarIcon strokeWidth={1} className={cn("text-yellow-600 ", isFilled && "fill-yellow-600")} />
    </Button>
  );
};
