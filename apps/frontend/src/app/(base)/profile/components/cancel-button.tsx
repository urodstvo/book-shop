"use client";

import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { XIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { toast } from "sonner";

async function cancelOrder(order_id: number) {
  const response = await fetch(`${API_URL}/orders/${order_id}`, {
    method: "DELETE",
    credentials: "include",
    mode: "cors",
  });

  if (!response.ok) {
    toast.error("Ошибка при отмене заказа");
    throw new Error("Failed to cancel order");
  }

  toast.success("Заказ отменен");
}

export const CancelButton = ({ order_id }: { order_id: number }) => {
  const router = useRouter();

  return (
    <Button
      size="icon"
      className="rounded-full"
      onClick={() => cancelOrder(order_id).then(() => router.refresh())}
    >
      <XIcon strokeWidth={2} />
    </Button>
  );
};
