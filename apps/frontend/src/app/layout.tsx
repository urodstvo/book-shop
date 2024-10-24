import type { Metadata } from "next";
import "./globals.css";
import { Toaster } from "@/components/ui/sonner";

export const metadata: Metadata = {
  title: "Книжный магазин",
  description: "Книжный онлайн магазин",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ru" className="px-[10px] md:px-10" style={{ colorScheme: "light" }}>
      <body
        className={`antialiased bg-background overflow-x-hidden grid grid-rows-[auto_1fr_auto] grid-cols-1 `}
      >
        {children}
        <Toaster />
      </body>
    </html>
  );
}
