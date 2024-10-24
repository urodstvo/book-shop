import { API_URL } from "@/env";
import { notFound } from "next/navigation";

export default async function BookDemoPage({ params: { id } }: { params: { id: string } }) {
  const response = await fetch(`${API_URL}/books/${id}/preview`, { next: { revalidate: 0 } });
  if (response.status === 404) return notFound();
  return <iframe src={`${API_URL}/books/${id}/preview`} className="w-full min-h-full" />;
}
