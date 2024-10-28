import { API_URL } from "@/env";
import { Metadata } from "next";
import { notFound } from "next/navigation";

export const metadata: Metadata = {
  title: "Демо | Книжный магазин",
};

export default async function BookDemoPage({ params: { id } }: { params: { id: string } }) {
  const response = await fetch(`${API_URL}/books/${id}/preview`);
  if (response.status === 404) return notFound();
  const html = await response.text();
  return <iframe srcDoc={html} className="w-full min-h-full" />;
}
