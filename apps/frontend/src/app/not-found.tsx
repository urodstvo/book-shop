import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function NotFound() {
  return (
    <main className="flex flex-col items-center justify-center h-screen gap-5">
      <div className="flex flex-col gap-0 items-center">
        <h1 className="text-[96px] font-bold uppercase">ОШИБКА 404</h1>
        <p className="text-base text-muted-foreground">Страница не найдена</p>
      </div>
      <Button asChild size="lg" className="rounded-full">
        <Link href="/">Вернуться на главную</Link>
      </Button>
    </main>
  );
}
