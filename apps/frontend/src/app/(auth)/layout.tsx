import { Button } from "@/components/ui/button";
import { BookUserIcon } from "lucide-react";
import Link from "next/link";

export default function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <header className="py-5">
        <Button asChild variant={null} size="icon" className="[&_svg]:size-[24px] size-10">
          <Link href="/" className="text-base font-medium">
            <BookUserIcon strokeWidth={2} />
          </Link>
        </Button>
      </header>
      <main className="h-[calc(100vh-100px)] flex size-full items-center justify-center">
        {children}
      </main>
    </>
  );
}
