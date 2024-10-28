"use client"; // Error boundaries must be Client Components

import { Button } from "@/components/ui/button";
import Link from "next/link";
import { useEffect } from "react";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    console.error(error);
  }, [error]);

  return (
    <main className="flex size-full items-center justify-center">
      <h2 className="text-2xl font-bold">Что-то пошло не так!</h2>
      <div className="flex items-center gap-5">
        <Button className="rounded-full" onClick={() => reset()}>
          Попробовать снова
        </Button>
        <Button className="rounded-full" variant="secondary" asChild>
          <Link href="/">Вернуться на главную</Link>
        </Button>
      </div>
    </main>
  );
}
