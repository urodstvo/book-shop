import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Панель администратора",
  description: "Панель администратора книжного магазина",
};

export default function RootLayout({
  books,
  users,
  orders,
}: Readonly<{
  books: React.ReactNode;
  users: React.ReactNode;
  orders: React.ReactNode;
}>) {
  return (
    <main className="grid grid-rows-3 lg:grid-rows-2 lg:grid-cols-2 gap-5">
      <div className="row-[1]">{books}</div>
      <div className="row-[2]">{orders}</div>
      <div className="row-span-2">{users}</div>
    </main>
  );
}
